package api

import (
	"encoding/json"
<<<<<<< HEAD
	"fmt"
	"net/url"
=======
	"net/url"
	"time"
>>>>>>> cd15077 (Some changes)
)

// ResponseConvert is a json response struct of convert response
type ResponseConvert struct {
	response
<<<<<<< HEAD
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
=======
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
>>>>>>> cd15077 (Some changes)

	//Setting form data
	data := url.Values{}
	data.Add("vid", idVideo)
	data.Add("k", token)

	content, err := sendRequest(data, BaseUrl+ConvertEndpoint)
	if err != nil {
<<<<<<< HEAD
		return ResponseConvert{}, err
	}

	fmt.Println(string(content))

	err = json.Unmarshal(content, &rc)
	if err != nil {
		return ResponseConvert{}, err
	}

	return rc, nil
=======
		return nil, err
	}

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
>>>>>>> cd15077 (Some changes)
}
