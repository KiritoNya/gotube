package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ResponseIndex is a type of index api response
type ResponseIndex struct {
	response
	Process string                         `json:"p"`
	VideoId string                         `json:"vid"`
	Title   string                         `json:"title"`
	Timing  int                            `json:"t"`
	Channel string                         `json:"a"`
	Links   map[string]map[string]Modality `json:"links"`
}

// NewResponseIndex creates a new ResponseIndex instance
func NewResponseIndex(videoLink string) (*ResponseIndex, error) {
	var ri ResponseIndex

	// Create settings form data
	data := url.Values{}
	data.Add("q", videoLink)
	data.Add("vt", IndexEndpointParam)

	// Send request
	content, err := sendRequest(data, BaseUrl+IndexEndpoint)
	if err != nil {
		return nil, err
	}

	// Parsing json response
	err = json.Unmarshal(content, &ri)
	if err != nil {
		return nil, err
	}

	return &ri, nil
}

// ExtractToken is a ResponseIndex method that extract token from ResponseIndex specifying the video options
func (ri *ResponseIndex) ExtractToken(opt Options) string {
	for key, value := range ri.Links {
		if key == string(opt.Format) {
			for _, value2 := range value {
				if value2.Quality == opt.Quality {
					return value2.Key
				}
			}
		}
	}
	return ""
}

func (ri *ResponseIndex) VideoOptions() (opts []*Options) {
	for format, value := range ri.Links {
		var vo Options
		for _, value2 := range value {
			vo.Format = Format(format)
			vo.Quality = value2.Quality
			opts = append(opts, &vo)
			opts = append(opts, &vo)
		}
	}

	return opts
}

// sendRequest is a utility function to send request to the api endpoint
func sendRequest(data url.Values, link string) ([]byte, error) {
	//Create request
	req, err := http.NewRequest("POST", link, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	//Set headers of HTTP request
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	req.Header.Add("Connection", "close")

	// Create client
	client := &http.Client{}

	//Do request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}
