// Copyright © 2025. Citrix Systems, Inc.

package citrixclient_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	citrixclient "github.com/citrix/citrix-daas-rest-go/client"
	"github.com/citrix/citrix-daas-rest-go/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ========== Test Helpers ==========

func setupMockAuthServer(t *testing.T, handler http.HandlerFunc) *httptest.Server {
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close) // instead of making callers defer server.Close()
	return server
}

func setupClientWithAuth(authUrl string, onPremises, isGov bool) *citrixclient.CitrixDaasClient {
	client, _ := test.NewTestDaaSClient()
	client.AuthConfig = &citrixclient.AuthenticationConfiguration{
		ClientId:     "test-client",
		ClientSecret: "test-secret",
		AuthUrl:      authUrl,
		OnPremises:   onPremises,
		IsGov:        isGov,
	}
	return client
}

func setupWemClient(authUrl, username, password string) *citrixclient.CitrixDaasClient {
	client, _ := test.NewTestDaaSClient()
	client.WemOnPremAuthConfig = &citrixclient.WemOnPremAuthentication{
		AuthUrl:       authUrl,
		AdminUserName: username,
		AdminPassword: password,
	}
	return client
}

// ========== SignIn Tests ==========

func TestSignIn_MissingClientId(t *testing.T) {
	client, _ := test.NewTestDaaSClient()
	client.AuthConfig = &citrixclient.AuthenticationConfiguration{
		ClientId:     "",
		ClientSecret: "test-secret",
		AuthUrl:      "https://auth.example.com/token",
	}

	token, resp, err := client.SignIn()

	require.Error(t, err)
	assert.Equal(t, "make sure customerid, clientid and clientsecret are not null", err.Error())
	assert.Empty(t, token)
	assert.Nil(t, resp)
}

func TestSignIn_MissingClientSecret(t *testing.T) {
	client, _ := test.NewTestDaaSClient()
	client.AuthConfig = &citrixclient.AuthenticationConfiguration{
		ClientId:     "test-client",
		ClientSecret: "",
		AuthUrl:      "https://auth.example.com/token",
	}

	token, resp, err := client.SignIn()

	require.Error(t, err)
	assert.Equal(t, "make sure customerid, clientid and clientsecret are not null", err.Error())
	assert.Empty(t, token)
	assert.Nil(t, resp)
}

func TestSignIn_ReuseValidToken(t *testing.T) {
	client, _ := test.NewTestDaaSClient()
	client.AuthConfig = &citrixclient.AuthenticationConfiguration{
		ClientId:     "test-client",
		ClientSecret: "test-secret",
		AuthUrl:      "https://auth.example.com/token",
	}
	// Token expires in 2 minutes, which is beyond the 1-minute refresh threshold
	futureTime := time.Now().UTC().Add(time.Minute * 2).Format(time.RFC3339)
	client.AuthToken = &citrixclient.AuthTokenModel{
		Token:     "CWSAuth bearer=existing-valid-token",
		ExpiresAt: futureTime,
	}

	token, resp, err := client.SignIn()

	assert.NoError(t, err)
	assert.Equal(t, "CWSAuth bearer=existing-valid-token", token)
	assert.Nil(t, resp)
}

func TestSignIn_ExpiredToken_ShouldRefresh(t *testing.T) {
	server := setupMockAuthServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(citrixclient.CCAuthResponse{
			TokenType:  "Bearer",
			Token:      "new-refreshed-token",
			Expiration: "3600",
		})
	}))

	client := setupClientWithAuth(server.URL, false, false)
	// Set token that expired 5 minutes ago to trigger refresh
	pastTime := time.Now().UTC().Add(time.Minute * -5).Format(time.RFC3339)
	client.AuthToken = &citrixclient.AuthTokenModel{
		Token:     "CWSAuth bearer=expired-token",
		ExpiresAt: pastTime,
	}

	token, resp, err := client.SignIn()

	assert.NoError(t, err)
	assert.Equal(t, "CWSAuth bearer=new-refreshed-token", token)
	assert.Nil(t, resp)
	assert.NotNil(t, client.AuthToken)
	assert.Equal(t, "CWSAuth bearer=new-refreshed-token", client.AuthToken.Token)
}

func TestSignIn_OnPremises_Success(t *testing.T) {
	expiryTime := time.Now().UTC().Add(time.Hour).Format(time.RFC3339)
	server := setupMockAuthServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		assert.True(t, ok)
		assert.Equal(t, "test-client", username)
		assert.Equal(t, "test-secret", password)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(citrixclient.TrustAuthResponse{
			Token:      "onprem-trust-token",
			Principal:  "test-user",
			UserId:     "user-123",
			CustomerId: "customer-456",
			ExpiresAt:  expiryTime,
		})
	}))

	client := setupClientWithAuth(server.URL, true, false)

	token, resp, err := client.SignIn()

	assert.NoError(t, err)
	assert.Equal(t, "CWSAuth bearer=onprem-trust-token", token)
	assert.Nil(t, resp)
	assert.NotNil(t, client.AuthToken)
	assert.Equal(t, expiryTime, client.AuthToken.ExpiresAt)
}

func TestSignIn_OnPremises_ErrorResponse(t *testing.T) {
	server := setupMockAuthServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid credentials"))
	}))

	client := setupClientWithAuth(server.URL, true, false)

	token, resp, err := client.SignIn()

	require.Error(t, err)
	assert.Contains(t, err.Error(), "could not sign into DDC")
	assert.Empty(t, token)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestSignIn_CloudOAuth_Success(t *testing.T) {
	server := setupMockAuthServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		err := r.ParseForm()
		assert.NoError(t, err)
		assert.Equal(t, "client_credentials", r.FormValue("grant_type"))
		assert.Equal(t, "test-client", r.FormValue("client_id"))
		assert.Equal(t, "test-secret", r.FormValue("client_secret"))

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(citrixclient.CCAuthResponse{
			TokenType:  "Bearer",
			Token:      "cloud-oauth-token",
			Expiration: "7200",
		})
	}))

	client := setupClientWithAuth(server.URL, false, false)

	token, resp, err := client.SignIn()

	assert.NoError(t, err)
	assert.Equal(t, "CWSAuth bearer=cloud-oauth-token", token)
	assert.Nil(t, resp)
	assert.NotNil(t, client.AuthToken)

	expiresAt, err := time.Parse(time.RFC3339, client.AuthToken.ExpiresAt)
	assert.NoError(t, err)
	expectedExpiry := time.Now().UTC().Add(time.Second * 7200)
	// Allow 5 second tolerance for time calculation differences
	assert.WithinDuration(t, expectedExpiry, expiresAt, time.Second*5)
}

func TestSignIn_CloudGov_Success(t *testing.T) {
	server := setupMockAuthServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		assert.Contains(t, string(body), "test-client")
		assert.Contains(t, string(body), "test-secret")

		w.WriteHeader(http.StatusOK)
		// Gov Cloud auth is coming from trust service so it is in a different format
		json.NewEncoder(w).Encode(citrixclient.CCTrustAuthResponse{
			Token:      "gov-trust-token",
			Principal:  "gov-user",
			UserId:     "gov-user-123",
			CustomerId: "gov-customer-456",
			ExpiresAt:  3600,
		})
	}))

	client := setupClientWithAuth(server.URL, false, true)

	token, resp, err := client.SignIn()

	assert.NoError(t, err)
	assert.Equal(t, "CWSAuth bearer=gov-trust-token", token)
	assert.Nil(t, resp)
	assert.NotNil(t, client.AuthToken)

	expiresAt, err := time.Parse(time.RFC3339, client.AuthToken.ExpiresAt)
	assert.NoError(t, err)
	expectedExpiry := time.Now().UTC().Add(time.Second * 3600)
	assert.WithinDuration(t, expectedExpiry, expiresAt, time.Second*5)
}

func TestPerformCCAuth_ErrorInResponse(t *testing.T) {
	server := setupMockAuthServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(citrixclient.CCAuthResponse{
			Error:            "invalid_client",
			ErrorDescription: "Client authentication failed",
		})
	}))

	client := setupClientWithAuth(server.URL, false, false)

	token, resp, err := client.SignIn()

	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid_client")
	assert.Contains(t, err.Error(), "Client authentication failed")
	assert.Empty(t, token)
	assert.NotNil(t, resp)
}

func TestPerformCCAuth_HttpError(t *testing.T) {
	server := setupMockAuthServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
	}))

	client := setupClientWithAuth(server.URL, false, false)

	token, resp, err := client.SignIn()

	require.Error(t, err)
	assert.Contains(t, err.Error(), "could not sign into Citrix Cloud")
	assert.Empty(t, token)
	assert.NotNil(t, resp)
}

// ========== SignInWemOnPrem Tests ==========

func TestSignInWemOnPrem_Success(t *testing.T) {
	server := setupMockAuthServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		assert.True(t, ok)
		assert.Equal(t, "wem-admin", username)
		assert.Equal(t, "wem-password", password)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(citrixclient.WemOnPremAuthResponse{
			Sid:       "session-sid-123",
			UserName:  "wem-admin",
			SessionId: "wem-session-abc-xyz",
		})
	}))

	client := setupWemClient(server.URL, "wem-admin", "wem-password")

	sessionId, resp, err := client.SignInWemOnPrem()

	assert.NoError(t, err)
	assert.Equal(t, "session wem-session-abc-xyz", sessionId)
	assert.Nil(t, resp)
}

func TestSignInWemOnPrem_AuthenticationFailed(t *testing.T) {
	server := setupMockAuthServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid username or password"))
	}))

	client := setupWemClient(server.URL, "wrong-user", "wrong-password")

	sessionId, resp, err := client.SignInWemOnPrem()

	require.Error(t, err)
	assert.Contains(t, err.Error(), "could not sign into wem onPremises host")
	assert.Empty(t, sessionId)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestSignInWemOnPrem_InvalidJson(t *testing.T) {
	server := setupMockAuthServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("not-valid-json"))
	}))

	client := setupWemClient(server.URL, "wem-admin", "wem-password")

	sessionId, resp, err := client.SignInWemOnPrem()

	require.Error(t, err)
	assert.Empty(t, sessionId)
	assert.Nil(t, resp)
}

// ========== Edge Cases and Error Scenarios ==========

func TestSignIn_InvalidExpiryFormat_ShouldRefresh(t *testing.T) {
	server := setupMockAuthServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(citrixclient.CCAuthResponse{
			TokenType:  "Bearer",
			Token:      "new-token-after-invalid",
			Expiration: "3600",
		})
	}))

	client := setupClientWithAuth(server.URL, false, false)
	// Invalid expiry format should cause token refresh due to parse failure
	client.AuthToken = &citrixclient.AuthTokenModel{
		Token:     "CWSAuth bearer=token-with-invalid-expiry",
		ExpiresAt: "not-a-valid-date",
	}

	token, resp, err := client.SignIn()

	assert.NoError(t, err)
	assert.Equal(t, "CWSAuth bearer=new-token-after-invalid", token)
	assert.Nil(t, resp)
}

func TestSignIn_TokenExpiresInLessThanOneMinute_ShouldRefresh(t *testing.T) {
	server := setupMockAuthServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(citrixclient.CCAuthResponse{
			TokenType:  "Bearer",
			Token:      "refreshed-token",
			Expiration: "3600",
		})
	}))

	client := setupClientWithAuth(server.URL, false, false)
	// Token expires in 30 seconds, which is within the 1-minute refresh threshold
	nearFutureTime := time.Now().UTC().Add(time.Second * 30).Format(time.RFC3339)
	client.AuthToken = &citrixclient.AuthTokenModel{
		Token:     "CWSAuth bearer=almost-expired-token",
		ExpiresAt: nearFutureTime,
	}

	token, resp, err := client.SignIn()

	assert.NoError(t, err)
	assert.Equal(t, "CWSAuth bearer=refreshed-token", token)
	assert.Nil(t, resp)
}

func TestSignIn_OnPremises_InvalidJson(t *testing.T) {
	server := setupMockAuthServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("not-valid-json"))
	}))

	client := setupClientWithAuth(server.URL, true, false)

	token, resp, err := client.SignIn()

	require.Error(t, err)
	assert.Empty(t, token)
	assert.Nil(t, resp)
}

func TestSignIn_CloudGov_InvalidJson(t *testing.T) {
	server := setupMockAuthServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("not-valid-json"))
	}))

	client := setupClientWithAuth(server.URL, false, true)

	token, resp, err := client.SignIn()

	require.Error(t, err)
	assert.Empty(t, token)
	assert.Nil(t, resp)
}

func TestSignIn_CloudGov_HttpError(t *testing.T) {
	server := setupMockAuthServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}))

	client := setupClientWithAuth(server.URL, false, true)

	// Use test context with zero-delay retry config for 500 error
	token, resp, err := client.SignInWithContext(test.DaaSTestContext())

	require.Error(t, err)
	assert.Contains(t, err.Error(), "could not sign into DDC")
	assert.Empty(t, token)
	assert.NotNil(t, resp)
}
