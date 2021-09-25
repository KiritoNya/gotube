package gotube

import (
	"github.com/KiritoNya/gotube/internal/pkg/api"
)

// Video is a type that contains video information
type Video struct {
	Id      string
	Title   string
	Url     string
	Options []*VideoOptions
	links   map[string]map[string]api.Modality
}

// VideoOptions are option of video
type VideoOptions struct {
	Format  Format
	Quality Quality
}

// Format of video
type Format string

var (
	FormatMp4 Format = "mp4"
	FormatM4a Format = "m4a"
	FormatMp3 Format = "mp3"
	Format3gp Format = "3gp"
)

// Quality of video
type Quality string

const (
	QualityVideo144   Quality = "144p"
	QualityVideo240   Quality = "240p"
	QualityVideo480   Quality = "480p"
	QualityVideo630   Quality = "630p"
	QualityVideo720   Quality = "720p"
	QualityVideo1080p Quality = "1080p"
	QualityAudio128   Quality = "128kbps"
)

// NewVideoById is a function that creates a new Video instance by id
func NewVideoById(id string) (*Video, error) {
	return newVideo(VideoBaseUrl + id)
}

// NewVideoByUrl is a function that creates a new Video instance by url
func NewVideoByUrl(url string) (*Video, error) {
	return newVideo(url)
}

// GetDirectLink is a Video method that returns the direct link of the video
func (v *Video) GetDirectLink(opts *VideoOptions) (string, error) {
	token := v.extractToken(
		api.Options{
			Format:  api.Format(opts.Format),
			Quality: string(opts.Quality),
		},
	)

	rc, err := api.NewResponseConvert(v.Id, token)
	if err != nil {
		return "", err
	}

	return rc.Url, nil
}

// extractToken is a utility method of Video for extract the option token
func (v *Video) extractToken(opts api.Options) string {
	var ri api.ResponseIndex
	ri.Links = v.links

	return ri.ExtractToken(opts)
}

// newVideo is a utility function for create a new video
func newVideo(url string) (*Video, error) {
	var v Video

	// Create responseIndex object
	ri, err := api.NewResponseIndex(url)
	if err != nil {
		return nil, err
	}

	opts := ri.VideoOptions()
	var vopts []*VideoOptions
	for _, opt := range opts {
		var v VideoOptions

		v.Format = Format(opt.Format)
		v.Quality = Quality(opt.Quality)
		vopts = append(vopts, &v)
	}

	v.Url = url
	v.Id = ri.VideoId
	v.Title = ri.Title
	v.Options = vopts
	v.links = ri.Links
	return &v, nil
}
