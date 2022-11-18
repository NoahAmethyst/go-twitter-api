package twitter

import (
	"context"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
)

type Client struct {
	*http.Client
}

// Use Oath2
func NewClient(twitterConsumeKey, twitterConsumerSecret string) *Client {
	// oauth2 configures a client that uses app credentials to keep a fresh token
	config := &clientcredentials.Config{
		ClientID:     twitterConsumeKey,
		ClientSecret: twitterConsumerSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	// http.Client will automatically authorize Requests
	TwitterCli := &Client{config.Client(context.Background())}
	return TwitterCli
}
