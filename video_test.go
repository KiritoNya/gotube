package gotube_test

import (
	"github.com/KiritoNya/gotube"
	"testing"
)

func TestNewVideoByUrl(t *testing.T) {
	v, err := gotube.NewVideoByUrl(TestInput.VideoUrl)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(v)
}

func TestVideo_GetDirectLink(t *testing.T) {
	v, err := gotube.NewVideoByUrl(TestInput.VideoUrl)
	if err != nil {
		t.Fatal(err)
	}

	options := gotube.VideoOptions{
		Format:  gotube.FormatMp3,
		Quality: gotube.QualityAudio128,
	}

	link, err := v.GetDirectLink(&options)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(link)
}
