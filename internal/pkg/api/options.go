package api

// Format is a video format
type Format string

const (
	// FormatMp4 is mp4 video format
	FormatMp4 Format = "mp4"

	// FormatM4a is m4a video format
	FormatM4a Format = "m4a"

	// FormatMp3 is mp3 video format
	FormatMp3 Format = "mp3"

	// Format3gp is 3gp video format
	Format3gp Format = "3gp"
)

type Options struct {
	Format  Format
	Quality string
}
