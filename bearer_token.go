package main

/*

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type TokenResponse struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   string `json:"expires_in"`
	AccessToken string `json:"access_token"`
}

func (cfg *config) getBearerToken(url string) (string, time.Time, error) {
	url := fmt.Sprintf("https://api.cloud.com/cctrustoauth2/%s/tokens/clients", cfg.CCID)

	data := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", cfg.ClientID, cfg.ClientSecret)
	body := bytes.NewBufferString(data)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("error making http request: %v", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("error making token request: %v", err)
	}
	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("error reading token response body: %v", err)
	}

	var tokenResp TokenResponse
	if err := json.Unmarshal(respBody, &tokenResp); err != nil {
		return "", time.Time{}, fmt.Errorf("error unmarshalling token response: %v", err)
	}

	expiryInt, err := strconv.Atoi(tokenResp.ExpiresIn)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("error converting expires_in to int: %v", err)
	}

	expireTime := time.Now().Add(time.Duration(expiryInt) * time.Second)

	return tokenResp.AccessToken, expireTime, nil
}
*/
