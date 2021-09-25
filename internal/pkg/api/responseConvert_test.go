package api_test

import (
<<<<<<< HEAD
	"github.com/KiritoNya/youtubeDownload/internal/pkg/api"
=======
	"github.com/KiritoNya/gotube/internal/pkg/api"
>>>>>>> cd15077 (Some changes)
	"testing"
)

func TestNewResponseConvert(t *testing.T) {
	rc, err := api.NewResponseConvert(TestInput.IdVideo, TestInput.Token)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(rc)
}
