package gotube

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"log"
)

// SearchOptions is a struct with the search option information
type SearchOptions struct {
	Query  string
	Limit  int64
	ApiKey string
}

type SearchResult struct {
	Channels  []*Channel
	Videos    []*Video
	Playlists []*Playlist
}

// Search è una funzione che data una query e un limit effettua una ricerca e torna una slice di video info
func Search(opts SearchOptions) (results *SearchResult, err error) {
	var sr SearchResult

	// Search on youtube
	items, err := searchYoutube(opts)
	if err != nil {
		panic(err)
	}

	// Iterate through each item and add it to the videos list.
	for _, item := range items {
		switch item.Id.Kind {
		case "youtube#video":
			v, err := searchParseVideo(item)
			if err != nil {
				return nil, err
			}
			sr.Videos = append(sr.Videos, v)
		case "youtube#channel":
			c, err := searchParseChannel(item)
			if err != nil {
				return nil, err
			}
			sr.Channels = append(sr.Channels, c)
		case "youtube#playlist":
			p, err := searchParsePlaylist(item)
			if err != nil {
				return nil, err
			}
			sr.Playlists = append(sr.Playlists, p)
		}
	}

	return &sr, nil
}

// searchYoutube is a utility function that search on youtube and return a slice of youtube.SearchResult
func searchYoutube(opts SearchOptions) ([]*youtube.SearchResult, error) {
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(opts.ApiKey))
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	// Make the API call to YouTube.
	call := service.Search.List([]string{"id", "snippet"}).
		Q(opts.Query).
		MaxResults(opts.Limit)
	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	return response.Items, nil
}

// searchParseVideo is a utility function that parse video search results
func searchParseVideo(item *youtube.SearchResult) (*Video, error) {
	var v Video
	var c Channel

	v.Id = item.Id.VideoId
	v.Title = item.Snippet.Title
	v.Description = item.Snippet.Description
	v.PublishedAt = item.Snippet.PublishedAt
	v.Thumbnails = Thumbnails{*item.Snippet.Thumbnails}

	// Parse channel
	c.Id = item.Snippet.ChannelId
	c.Title = item.Snippet.ChannelTitle
	v.Channel = &c

	return &v, nil
}

// searchParseChannel is a utility function that parse channel search results
func searchParseChannel(item *youtube.SearchResult) (*Channel, error) {
	var c Channel

	c.Id = item.Snippet.ChannelId
	c.Title = item.Snippet.ChannelTitle
	c.Description = item.Snippet.Description
	c.Thumbnails = Thumbnails{*item.Snippet.Thumbnails}

	data, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	fmt.Println("Channel:", string(data))

	return &c, nil
}

// searchParsePlaylist is a utility function that parse playlist search results
func searchParsePlaylist(item *youtube.SearchResult) (*Playlist, error) {
	var p Playlist
	var c Channel

	p.Id = item.Snippet.ChannelId
	p.Title = item.Snippet.ChannelTitle
	p.Description = item.Snippet.Description
	p.Thumbnails = Thumbnails{*item.Snippet.Thumbnails}

	// Parse channel
	c.Id = item.Snippet.ChannelId
	c.Title = item.Snippet.ChannelTitle
	p.Channel = &c

	return &p, nil
}
