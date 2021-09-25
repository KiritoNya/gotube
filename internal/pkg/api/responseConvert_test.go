package api_test

import (
	"github.com/KiritoNya/youtubeDownload/internal/pkg/api"
	"testing"
)

func TestNewResponseConvert(t *testing.T) {
	rc, err := api.NewResponseConvert(TestInput.IdVideo, TestInput.Token)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(rc)
}
