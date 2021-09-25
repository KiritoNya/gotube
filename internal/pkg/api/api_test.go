package api_test

var TestInput = struct {
	IdVideo  string
	Token    string
	UrlVideo string
}{
	IdVideo:  "GintFI8C2qw",
	Token:    "0+e4UkTXLLnzE4KHW5ZonRM2PRF1HQ7L5l7v3sVU68OOlzifQPX/+2+SOnP54w==",
	UrlVideo: "https://www.youtube.com/watch?v=GintFI8C2qw",
}

var OutputTest = struct {
	ExtractToken string
}{
	"0+e4UkTXLLnzE4KHW5ZonRM2PRF1HQ7L5l7v3sVU68OOlzifQPX/+2+SOnP54w==",
}
