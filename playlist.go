package gotube

// Playlist is a struct with playlist information
type Playlist struct {
	Id string
	Title string
	Description string
	Channel *Channel
	Thumbnails Thumbnails
}