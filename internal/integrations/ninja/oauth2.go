// Package ninja
package ninja

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
}

func NewClient(ctx context.Context) (*Client, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	clientID := os.Getenv("NINJA_CLIENT_ID")
	clientSecret := os.Getenv("NINJA_CLIENT_SECRET")
	tokenURL := os.Getenv("NINJA_TOKEN_URL")
	baseURL := os.Getenv("NINJA_INSTANCE_URL")

	if clientID == "" || clientSecret == "" || tokenURL == "" || baseURL == "" {
		return nil, fmt.Errorf("missing required NINJA_* environment variables")
	}

	cfg := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenURL,
		Scopes:       []string{"monitoring"},
	}

	tokenFetchClient := &http.Client{Timeout: 10 * time.Second}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, tokenFetchClient)

	return &Client{
		httpClient: cfg.Client(ctx),
		baseURL:    baseURL,
	}, nil
}
