package api

const BaseUrl string = "https://yt1s.com/api"
const IndexEndpoint string = "/ajaxSearch/index"
const ConvertEndpoint string = "/ajaxConvert/convert"
const IndexEndpointParam string = "home"

// response is a type that contains info of response
type response struct {
	Status  string `json:"status"`
	Message string `json:"mess"`
}
