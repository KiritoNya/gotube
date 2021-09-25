package gotube_test

import "github.com/KiritoNya/gotube"

var TestInput = struct {
	VideoUrl string
}{
	"https://www.youtube.com/watch?v=GintFI8C2qw",
}

var TestOutput = struct {
	Video *gotube.Video
}{
	Video: &gotube.Video{
		Id:      "GintFI8C2qw",
		Title:   "Miss Kobayashi's Dragon Maid - Opening (HD)",
		Url:     TestInput.VideoUrl,
		Options: nil,
	},
}
