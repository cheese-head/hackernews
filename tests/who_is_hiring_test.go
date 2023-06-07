package tests

import (
	"testing"

	"github.com/cheese-head/hackernews/api"
)

func TestWhoIsHiring(t *testing.T) {
	hackernews, err := api.NewClient()
	if err != nil {
		t.Fatal(err)
	}
	user, err := hackernews.GetUser("whoishiring")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(user.ID)
}

// func TestWhoIsHiringJobPosts(t *testing.T) {

// 	hackernews, err := api.NewClient()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	item, err := hackernews.GetItem(36152014)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	t.Log(item.Type)

// 	if models.IsStory(item.Type) {

// 		for _, id := range item.Kids {
// 			t.Log("child id: ", id)

// 			childItem, err := hackernews.GetItem(id)
// 			if err != nil {
// 				t.Log(err)
// 			}

// 			data, err := json.Marshal(&childItem)
// 			if err != nil {
// 				t.Log(err)
// 			}
// 			t.Log(string(data))
// 		}
// 	}

// }
