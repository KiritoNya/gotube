package api

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// ResponseConvert is a json response struct of convert response
type ResponseConvert struct {
	response
	CStatus string `json:"c_status"`
	VideoId string `json:"vid"`
	Title   string `json:"title"`
	Format  string `json:"ftype"`
	Quality string `json:"quality"`
	Url     string `json:"dlink"`
}

// NewResponseConvert creates a ResponseConvert instance
func NewResponseConvert(idVideo, token string) (ResponseConvert, error) {
	var rc ResponseConvert

	//Setting form data
	data := url.Values{}
	data.Add("vid", idVideo)
	data.Add("k", token)

	content, err := sendRequest(data, BaseUrl+ConvertEndpoint)
	if err != nil {
		return ResponseConvert{}, err
	}

	fmt.Println(string(content))

	err = json.Unmarshal(content, &rc)
	if err != nil {
		return ResponseConvert{}, err
	}

	return rc, nil
}
