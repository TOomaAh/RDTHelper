package realdebrid

type Link struct {
	ID       string `json:"id"`
	Filename string `json:"filename"`
	MimeType string `json:"mimeType"`
	Link     string `json:"link"`
	Host     string `json:"host"`
	Download string `json:"download"`

	Chunks   int64 `json:"chunks"`
	Crc      int64 `json:"crc"`
	FileSize int64 `json:"fileSize"`

	Streamable int `json:"streamable"`
}
