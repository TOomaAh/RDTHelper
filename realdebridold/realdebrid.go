package realdebridold

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type RealDebridClient struct {
	BaseURL *url.URL
	headers http.Header
	client  *http.Client
}

var (
	ErrorInvalidRequest  = errors.New("invalid request")
	ErrorInvalidURL      = errors.New("invalid URL")
	ErrorCannotParsePath = errors.New("cannot parse path")
	ErrorCannotReadBody  = errors.New("Cannot read body")
	Error401             = errors.New("Unauthorized")
	Error403             = errors.New("Forbidden")
	Error404             = errors.New("Not Found")
	Error500             = errors.New("Internal Server Error")
)

func NewRealDebridClient() *RealDebridClient {

	parsed, err := url.Parse("https://api.real-debrid.com")
	if err != nil {
		panic(err)
	}

	headers := make(http.Header)
	headers.Set("Accept", "application/json")

	return &RealDebridClient{
		BaseURL: parsed,
		headers: headers,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *RealDebridClient) SetToken(token string) {
	c.headers.Set("Authorization", "Bearer "+token)
}

func (c *RealDebridClient) NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	path = "/rest/1.0" + path

	u, err := c.BaseURL.Parse(path)
	if err != nil {
		return nil, ErrorCannotParsePath
	}

	request := c.BaseURL.ResolveReference(u)
	log.Println(request)

	req, err := http.NewRequest(method, request.String(), body)
	if err != nil {
		return nil, ErrorInvalidRequest
	}
	req.Header = c.headers
	return req, nil
}

func (r *RealDebridClient) Get(req *http.Request, v interface{}) error {
	resp, err := r.client.Do(req)

	if err != nil {
		return nil
	}

	switch resp.StatusCode {
	case http.StatusNotFound:
		return Error404
	case http.StatusInternalServerError:
		return Error500
	case http.StatusUnauthorized:
		return Error401
	case http.StatusForbidden:
		return Error403
	}

	defer resp.Body.Close()

	if v != nil {
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			return ErrorCannotReadBody
		}

		return json.Unmarshal(body, v)
	}
	return nil
}
