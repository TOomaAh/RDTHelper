package realdebrid

type File struct {
	ID       int64  `json:"id"`
	Path     string `json:"path"`
	Bytes    int64  `json:"bytes"`
	Selected bool   `json:"selected"`
}
