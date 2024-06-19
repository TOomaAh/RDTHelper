package realdebrid

import "net/http"

type Torrent struct {
	ID       string   `json:"id"`
	Filename string   `json:"filename"`
	Hash     string   `json:"hash"`
	Bytes    int64    `json:"bytes"`
	Host     string   `json:"host"`
	Split    int64    `json:"split"`
	Progress float64  `json:"progress"`
	Status   string   `json:"status"`
	Added    string   `json:"added"`
	Files    []File   `json:"files"`
	Links    []string `json:"links"`
	Ended    string   `json:"ended"`
	Speed    int64    `json:"speed"`
	Seeders  int64    `json:"seeders"`
}

func (t *RealDebridClient) GetTorrents() (*[]Torrent, error) {
	var torrents *[]Torrent
	req, err := t.NewRequest(http.MethodGet, "/torrents", nil, nil)

	if err != nil {
		return nil, err
	}

	err = t.Get(req, &torrents)

	if err != nil {
		return nil, err
	}

	return torrents, nil
}

func (t *RealDebridClient) GetTorrentInfo(id string) *Torrent {
	return &Torrent{}
}

func (t *RealDebridClient) DeleteTorrent(id string) error {
	return nil
}

func (t *RealDebridClient) AcceptTorrent(id string) error {
	return nil
}

func (t *RealDebridClient) UploadTorrent() error {
	return nil
}

func (t *RealDebridClient) DebridTorrent() error {
	return nil
}
