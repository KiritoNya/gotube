package api

// Modality is a video modality
type Modality struct {
	Size    string `json:"size"`
	Format  string `json:"f"`
	Quality string `json:"q"`
	Key     string `json:"k"`
}
