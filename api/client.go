package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cheese-head/hackernews/api/models"
)

const (
	itemURL = "https://hacker-news.firebaseio.com/v0/item"
	userURL = "https://hacker-news.firebaseio.com/v0/user"
)

// Client is the Hackernews client for interacting with the Hackernews API.
type Client struct {
	httpClient *http.Client
}

// Options are used to configure the Hackernews client.
type Options func(*Client)

// NewClient returns a new Hackernews client with the given options.
func NewClient(options ...Options) (*Client, error) {
	client := &Client{
		httpClient: http.DefaultClient,
	}
	for _, option := range options {
		option(client)
	}
	return client, nil
}

// WithHttpClient sets the http client for the Hackernews client.
// By default, the http.DefaultClient is used.
func WithHttpClient(httpClient *http.Client) func(*Client) {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// GetItem returns the item with the given id.
func (c *Client) GetItem(id int) (*models.Item, error) {
	url := fmt.Sprintf("%s/%v.json", itemURL, id)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	item := &models.Item{}
	err = json.NewDecoder(resp.Body).Decode(item)
	return item, err
}

// GetUser returns the user with the given username.
func (c *Client) GetUser(username string) (*models.User, error) {
	url := fmt.Sprintf("%s/%s.json", userURL, username)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	v := &models.User{}
	err = json.NewDecoder(resp.Body).Decode(v)
	return v, err
}
