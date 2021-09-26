package gotube

import "google.golang.org/api/youtube/v3"

type Thumbnails struct {
	youtube.ThumbnailDetails
}

//type Thumbnails map[string]Thumbnail