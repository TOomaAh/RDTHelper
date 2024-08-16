package realdebrid

import (
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

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
	req, err := t.NewRequest(http.MethodDelete, "/torrents/"+id, nil, nil)

	if err != nil {
		return err
	}

	err = t.Delete(req)

	if err != nil {
		return err
	}

	return nil

}

func (t *RealDebridClient) AcceptTorrent(id string) error {
	header := http.Header{}
	header.Set("Content-Type", "application/x-www-form-urlencoded")

	body := url.Values{}
	body.Set("files", "all")

	req, err := t.NewRequest(http.MethodPost, "/torrents/selectFiles/"+id, header, strings.NewReader(body.Encode()))

	if err != nil {
		return err
	}

	err = t.Post(req, nil)

	if err != nil {
		return err
	}

	return nil
}

func (t *RealDebridClient) UploadTorrent(files []*multipart.FileHeader) error {
	header := http.Header{}
	header.Set("Content-Type", "application/x-bittorrent")

	for _, file := range files {
		content, err := file.Open()

		if err != nil {
			return err
		}

		defer content.Close()

		req, err := t.NewRequest(http.MethodPost, "/torrents/addTorrent", header, content)

		if err != nil {
			return err
		}

		err = t.Put(req, nil)

		if err != nil {
			return err
		}

	}

	return nil

}

func (t *RealDebridClient) DebridTorrent(link string) (*Link, error) {
	var l Link
	header := http.Header{}
	header.Set("Content-Type", "application/x-www-form-urlencoded")

	body := url.Values{}
	body.Set("link", link)

	req, err := t.NewRequest(http.MethodPost, "/unrestrict/link", header, strings.NewReader(body.Encode()))

	if err != nil {
		return nil, err
	}

	err = t.Post(req, &l)

	if err != nil {
		return nil, err
	}

	return &l, nil
}
