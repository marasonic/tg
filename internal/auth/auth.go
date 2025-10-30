package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/spf13/viper"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func GetToken() (string, error) {
	keycloakURL := viper.GetString("keycloak.url")
	realm := viper.GetString("keycloak.realm")
	clientID := viper.GetString("keycloak.client_id")
	clientSecret := viper.GetString("keycloak.client_secret")

	if clientID == "" || clientSecret == "" {
		return "", fmt.Errorf("keycloak client_id and client_secret must be configured")
	}

	tokenURL := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token", keycloakURL, realm)

	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create token request: %w", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to get token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get token, status code: %d", resp.StatusCode)
	}

	var tokenResponse TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return "", fmt.Errorf("failed to decode token response: %w", err)
	}

	return tokenResponse.AccessToken, nil
}
