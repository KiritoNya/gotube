package gotube_test

import (
	"encoding/json"
	"github.com/KiritoNya/gotube"
	"testing"
)

func TestSearch(t *testing.T) {
	sr, err := gotube.Search(gotube.SearchOptions{
		Query:  "JPOP",
		Limit:  5,
		ApiKey: "AIzaSyDG4bWuL-cpgsO0nbr7gWBeWt9aXDVqgoI",
	})
	if err != nil {
		t.Fatal(err)
	}

	data, err := json.MarshalIndent(sr, " ", "\t")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))
}