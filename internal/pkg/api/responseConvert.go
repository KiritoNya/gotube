package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// ResponseConvert is a json response struct of convert response
type ResponseConvert struct {
	response
	CStatus    string `json:"c_status"`
	VideoId    string `json:"vid"`
	Title      string `json:"title"`
	Format     string `json:"ftype"`
	Quality    string `json:"quality"`
	Url        string `json:"dlink"`
	Expiration int    `json:"e_time,omitempty"`
	Bid        string `json:"b_id,omitempty"`
}

// NewResponseConvert creates a ResponseConvert instance
func NewResponseConvert(idVideo, token string) (*ResponseConvert, error) {
	var rcReturn ResponseConvert

	//Setting form data
	data := url.Values{}
	data.Add("vid", idVideo)
	data.Add("k", token)

	content, err := sendRequest(data, BaseUrl+ConvertEndpoint)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(content))

	err = json.Unmarshal(content, &rcReturn)
	if err != nil {
		return nil, err
	}

	if rcReturn.Expiration != 0 {
		for {
			var rc ResponseConvert
			data := url.Values{}
			data.Add("vid", idVideo)
			data.Add("b_id", rcReturn.Bid)

			content, err := sendRequest(data, BaseUrl+CheckTaskEndpoint)
			if err != nil {
				return nil, err
			}

			err = json.Unmarshal(content, &rc)
			if err != nil {
				return nil, err
			}

			if rc.CStatus != "FAILED" {
				rcReturn = rc
				break
			}
			time.Sleep(1 * time.Second)
		}

	}

	return &rcReturn, nil
}
