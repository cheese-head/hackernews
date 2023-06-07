# Hackernews Go Client
This is a Go client library that provides convenient access to the hackernews firebase API.

## Installation

```bash
go get github.com/cheese-head/hackernews
```

## Usage

Some examples of how to use the library are shown below.

```go
package tests

import (
	"github.com/cheese-head/hackernews/api"
	"net/http"
	"testing"
	"time"
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

	item, err := client.GetItem(564554)
	if err != nil {
		t.Fatal(err)
	}

	if item.ID != 564554 {
		t.Fatal(err)
	}
}

```
## Contributing

Contributions are welcome! If you find any issues or want to add new features to the library, please submit a pull request.

## License

This library is licensed under the MIT License. Feel free to use, modify, and distribute it according to your needs.


