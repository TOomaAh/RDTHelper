package realdebrid

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

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

type RealDebridClient struct {
	Host           *url.URL
	Token          string
	httpClient     *http.Client
	defaultHeaders *http.Header
	User           *User
}

type User struct {
	ID         int64  `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Points     int64  `json:"points"`
	Locale     string `json:"locale"`
	Avatar     string `json:"avatar"`
	Type       string `json:"type"`
	Premium    bool   `json:"premium"`
	Expiration string `json:"expiration"`
}

func Authentication(token string) *RealDebridClient {
	parsed, err := url.Parse("https://api.real-debrid.com")
	if err != nil {
		panic(err)
	}

	headers := make(http.Header)
	headers.Set("Accept", "application/json")
	headers.Set("Authorization", "Bearer "+token)

	return &RealDebridClient{
		Host:           parsed,
		Token:          token,
		defaultHeaders: &headers,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (r *RealDebridClient) NewRequest(method, path string, headers http.Header, body io.Reader) (*http.Request, error) {
	path = "/rest/1.0" + path

	u, err := r.Host.Parse(path)
	if err != nil {
		return nil, ErrorCannotParsePath
	}

	request := r.Host.ResolveReference(u)
	log.Println(request)

	req, err := http.NewRequest(method, request.String(), body)
	if err != nil {
		return nil, ErrorInvalidRequest
	}

	if headers != nil {
		// append headers to the request
		for key, values := range *r.defaultHeaders {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}
	}
	return req, nil
}

func (r *RealDebridClient) Get(req *http.Request, v interface{}) error {
	resp, err := r.httpClient.Do(req)

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
