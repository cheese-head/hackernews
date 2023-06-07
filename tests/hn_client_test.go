package tests

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/cheese-head/hackernews/api"
)

func TestHackernewsClient(t *testing.T) {
	client, err := api.NewClient()
	if err != nil {
		t.Fatal(err)
	}

	item, err := client.GetItem(context.Background(), 564554)
	if err != nil {
		t.Fatal(err)
	}

	if item.ID != 564554 {
		t.Fatal(err)
	}
}

func TestHackerNewsClientWithOptions(t *testing.T) {

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	client, err := api.NewClient(
		api.WithHttpClient(httpClient),
	)
	if err != nil {
		t.Fatal(err)
	}

	item, err := client.GetItem(context.Background(), 564554)
	if err != nil {
		t.Fatal(err)
	}

	if item.ID != 564554 {
		t.Fatal(err)
	}
}
