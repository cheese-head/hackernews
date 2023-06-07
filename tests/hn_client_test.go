package tests

import (
	"hackernews/api"
	"testing"
)

func TestHackernewsClient(t *testing.T) {
	client, err := api.NewClient()
	if err != nil {
		t.Fatal(err)
	}

	item, err := client.GetItem(564554)
	if err != nil {
		t.Fatal(err)
	}

	if item.ID != 564554 {
		t.Fatal(err)
	}
}
