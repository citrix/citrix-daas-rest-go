// Copyright © 2024. Citrix Systems, Inc.

package citrixclient

import (
	"bytes"
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

// AuthResponse -
type CCTrustAuthResponse struct {
	Token      string `json:"Token"`
	Principal  string `json:"Principal"`
	UserId     string `json:"UserId"`
	CustomerId string `json:"CustomerId"`
	ExpiresAt  int    `json:"ExpiresAt"`
}

// SignIn - Get a new token for user
func (c *CitrixDaasClient) SignIn() (string, *http.Response, error) {
	if c.AuthConfig.ClientId == "" || c.AuthConfig.ClientSecret == "" {
		return "", nil, fmt.Errorf("make sure customerid, clientid and clientsecret are not null")
	}

	if c.AuthToken != nil {
		tokenExpirationTime, err := time.Parse(time.RFC3339, c.AuthToken.ExpiresAt)

		if err == nil && tokenExpirationTime.After(time.Now().UTC().Add(time.Minute*time.Duration(1))) {
			return c.AuthToken.Token, nil, nil
		}
	}

	var token string
	var expiresAt string

	if c.AuthConfig.OnPremises {
		client := c.ApiClient.GetConfig().HTTPClient
		req, err := http.NewRequest(http.MethodPost, c.AuthConfig.AuthUrl, nil)
		if err != nil {
			return "", nil, err
		}
		req.SetBasicAuth(c.AuthConfig.ClientId, c.AuthConfig.ClientSecret)
		resp, err := client.Do(req)
		if err != nil {
			return "", nil, err
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", nil, err
		}

		if resp.StatusCode > 200 {
			return "", resp, fmt.Errorf("could not sign into DDC, %s", string(body))
		}

		ar := TrustAuthResponse{}
		err = json.Unmarshal(body, &ar)
		if err != nil {
			return "", nil, err
		}

		token = fmt.Sprintf("CWSAuth bearer=%s", ar.Token)
		expiresAt = ar.ExpiresAt
	} else if !c.AuthConfig.IsGov {
		// If Auth is not for Citrix Cloud Gov, use oauth2
		grant_type := "client_credentials"
		operation := func() (CCAuthResponse, *http.Response, error) {
			return performCCAuth(c.AuthConfig.AuthUrl, url.Values{"grant_type": {grant_type}, "client_id": {c.AuthConfig.ClientId}, "client_secret": {c.AuthConfig.ClientSecret}})
		}

		ccAuthResponse, resp, err := RetryOperationWithExponentialBackOff(operation, 10, 3)

		if err != nil {
			return "", resp, err
		}

		token = fmt.Sprintf("CWSAuth bearer=%s", ccAuthResponse.Token)
		expiryInSeconds, _ := strconv.Atoi(ccAuthResponse.Expiration)
		expiresAtTime := time.Now().UTC().Add(time.Second * time.Duration(expiryInSeconds))
		expiresAt = expiresAtTime.Format(time.RFC3339)
	} else {
		// If Auth is for Citrix Cloud Gov, use trust service
		client := c.ApiClient.GetConfig().HTTPClient

		operation := func() (CCTrustAuthResponse, *http.Response, error) {
			return performCCTrustAuth(client, c.AuthConfig.AuthUrl, []byte(fmt.Sprintf(`{"ClientId":"%s", "ClientSecret":"%s"}`, c.AuthConfig.ClientId, c.AuthConfig.ClientSecret)))
		}

		ccTrustAuthResponse, resp, err := RetryOperationWithExponentialBackOff(operation, 10, 3)

		if err != nil {
			return "", resp, err
		}

		token = fmt.Sprintf("CWSAuth bearer=%s", ccTrustAuthResponse.Token)
		expiresAtTime := time.Now().UTC().Add(time.Second * time.Duration(ccTrustAuthResponse.ExpiresAt))
		expiresAt = expiresAtTime.Format(time.RFC3339)
	}

	authTokenModel := AuthTokenModel{}
	authTokenModel.Token = token
	authTokenModel.ExpiresAt = expiresAt
	c.AuthToken = &authTokenModel

	return token, nil, nil
}

func performCCAuth(authUrl string, urlValues url.Values) (CCAuthResponse, *http.Response, error) {
	ccAuthResponse := CCAuthResponse{}
	httpResp, err := http.PostForm(authUrl, urlValues)
	if err != nil {
		return ccAuthResponse, httpResp, err
	}
	defer httpResp.Body.Close()
	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return ccAuthResponse, httpResp, err
	}

	if httpResp.StatusCode > 200 {
		return ccAuthResponse, httpResp, fmt.Errorf("could not sign into Citrix Cloud, %s", string(body))
	}

	err = json.Unmarshal(body, &ccAuthResponse)
	if err != nil {
		return ccAuthResponse, httpResp, err
	}

	if ccAuthResponse.Error != "" {
		return ccAuthResponse, httpResp, fmt.Errorf("could not sign into Citrix Cloud, %s: %s", ccAuthResponse.Error, ccAuthResponse.ErrorDescription)
	}

	return ccAuthResponse, httpResp, err
}

func performCCTrustAuth(client *http.Client, authUrl string, jsonStr []byte) (CCTrustAuthResponse, *http.Response, error) {
	ccTrustResponse := CCTrustAuthResponse{}

	req, err := http.NewRequest(http.MethodPost, authUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		return ccTrustResponse, nil, err
	}
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return ccTrustResponse, nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ccTrustResponse, nil, err
	}

	if resp.StatusCode > 200 {
		return ccTrustResponse, resp, fmt.Errorf("could not sign into DDC, %s", string(body))
	}

	err = json.Unmarshal(body, &ccTrustResponse)
	if err != nil {
		return ccTrustResponse, nil, err
	}

	return ccTrustResponse, resp, err
}
