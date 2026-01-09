package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/saadahmedbd/Treestore/Config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// GoogleUserInfo represents the user data from Google
type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

// GetGoogleOAuthConfig returns the Google OAuth2 configuration
func GetGoogleOAuthConfig() *oauth2.Config {
	cnf := Config.GetConfig()
	return &oauth2.Config{
		ClientID:     cnf.GoogleClientID,
		ClientSecret: cnf.GoogleClientSecret,
		RedirectURL:  cnf.GoogleRedirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

// GetGoogleUserInfo fetches user information from Google
func GetGoogleUserInfo(accessToken string) (*GoogleUserInfo, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get user info: %s", string(body))
	}

	var userInfo GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to decode user info: %w", err)
	}

	return &userInfo, nil
}

// GenerateStateToken generates a random state token for OAuth
func GenerateStateToken() (string, error) {
	return GenerateRefreshToken() // Reuse existing random token generator
}
