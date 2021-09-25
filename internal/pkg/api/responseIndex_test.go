package api_test

import (
	"fmt"
<<<<<<< HEAD
	"github.com/KiritoNya/youtubeDownload/internal/pkg/api"
=======
	"github.com/KiritoNya/gotube/internal/pkg/api"
>>>>>>> cd15077 (Some changes)
	"testing"
)

func TestNewResponseIndex(t *testing.T) {
	ri, err := api.NewResponseIndex(TestInput.UrlVideo)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ri)
}

func TestResponseIndex_ExtractToken(t *testing.T) {
	ri, err := api.NewResponseIndex(TestInput.UrlVideo)
	if err != nil {
		t.Fatal(err)
	}

	token := ri.ExtractToken(api.Options{Format: "mp3", Quality: "128kbps"})

	if token != OutputTest.ExtractToken {
		t.Fatal(fmt.Sprintf("\nExpected: '%s'\nObtained: '%s'", OutputTest.ExtractToken, token))
	}
}

func TestResponseIndex_VideoOptions(t *testing.T) {
	ri, err := api.NewResponseIndex(TestInput.UrlVideo)
	if err != nil {
		t.Fatal(err)
	}
	opts := ri.VideoOptions()

	t.Log(opts)
}
