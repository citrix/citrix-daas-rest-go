package citrixclient

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// AuthResponse -
type CCAuthResponse struct {
	TokenType        string `json:"token_type"`
	Token            string `json:"access_token"`
	Expiration       string `json:"expires_in"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// AuthResponse -
type TrustAuthResponse struct {
	Token      string `json:"Token"`
	Principal  string `json:"Principal"`
	UserId     string `json:"UserId"`
	CustomerId string `json:"CustomerId"`
	ExpiresAt  string `json:"ExpiresAt"`
}

// SignIn - Get a new token for user
func (c *CitrixDaasClient) SignIn() (string, error) {
	if c.AuthConfig.ClientId == "" || c.AuthConfig.ClientSecret == "" {
		return "", fmt.Errorf("make sure customerid, clientid and clientsecret are not null")
	}

	if c.AuthToken != nil {
		tokenExpirationTime, err := time.Parse(time.RFC3339, c.AuthToken.ExpiresAt)

		if err == nil && tokenExpirationTime.After(time.Now().UTC().Add(time.Minute*time.Duration(1))) {
			return c.AuthToken.Token, nil
		}
	}

	var token string
	var expiresAt string

	if c.AuthConfig.OnPremise {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		req, _ := http.NewRequest("POST", c.AuthConfig.AuthUrl, nil)
		req.SetBasicAuth(c.AuthConfig.ClientId, c.AuthConfig.ClientSecret)
		resp, err := client.Do(req)
		if err != nil {
			return "", err
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		if resp.StatusCode > 200 {
			return "", fmt.Errorf("Could not sign into DDC, %s", string(body))
		}

		ar := TrustAuthResponse{}
		err = json.Unmarshal(body, &ar)
		if err != nil {
			return "", err
		}

		token = fmt.Sprintf("CWSAuth bearer=%s", ar.Token)
		expiresAt = ar.ExpiresAt

	} else {
		grant_type := "client_credentials"
		resp, err := http.PostForm(c.AuthConfig.AuthUrl, url.Values{"grant_type": {grant_type}, "client_id": {c.AuthConfig.ClientId}, "client_secret": {c.AuthConfig.ClientSecret}})
		if err != nil {
			return "", err
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		if resp.StatusCode > 200 {
			return "", fmt.Errorf("Could not sign into Citrix Cloud, %s", string(body))
		}

		ar := CCAuthResponse{}
		err = json.Unmarshal(body, &ar)
		if err != nil {
			return "", err
		}

		if ar.Error != "" {
			return "", fmt.Errorf("Could not sign into Citrix Cloud, %s: %s", ar.Error, ar.ErrorDescription)
		}

		token = fmt.Sprintf("CWSAuth bearer=%s", ar.Token)
		expiryInSeconds, _ := strconv.Atoi(ar.Expiration)
		expiresAtTime := time.Now().UTC().Add(time.Second * time.Duration(expiryInSeconds))
		expiresAt = expiresAtTime.Format(time.RFC3339)
	}

	authTokenModel := AuthTokenModel{}
	authTokenModel.Token = token
	authTokenModel.ExpiresAt = expiresAt
	c.AuthToken = &authTokenModel

	return token, nil
}
