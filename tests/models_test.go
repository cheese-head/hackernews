package tests

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/cheese-head/hackernews/api/models"
)

func TestModels(t *testing.T) {

	item := &models.Item{}

	data, err := ioutil.ReadFile("../api/json/comment.json")
	if err != nil {
		t.Fatal("Error when opening file: ", err)
	}
	err = json.Unmarshal(data, item)
	if err != nil {
		t.Fatal(err)
	}

	switch item.Type {
	case models.TypeAsk:
		t.Log("item is of type ask")
	case models.TypeComment:
		t.Log("item is of type comment")
	default:
		t.Fatal("unknown type")
	}
}
