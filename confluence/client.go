package confluence

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	client    *http.Client
	baseURL   *url.URL
	publicURL *url.URL
}

type NewClientInput struct {
	site  string
	user  string
	token string
}

func NewClient(input *NewClientInput) (*Client, error) {
	publicURL := url.URL{
		Scheme: "https",
		Host:   input.site + ".atlassian.net",
	}
	baseURL := publicURL
	baseURL.User = url.UserPassword(input.user, input.token)
	return &Client{
		client: &http.Client{
			Timeout: time.Second * 10,
		},
		baseURL:   &baseURL,
		publicURL: &publicURL,
	}, nil
}

func (c *Client) Post(path string, body interface{}, result interface{}) error {
	return c.do("POST", path, body, result)
}

func (c *Client) Get(path string, result interface{}) error {
	return c.do("GET", path, nil, result)
}

func (c *Client) Put(path string, body interface{}, result interface{}) error {
	return c.do("PUT", path, body, result)
}

func (c *Client) Delete(path string) error {
	return c.do("DELETE", path, nil, nil)
}

func (c *Client) do(method string, path string, body interface{}, result interface{}) error {
	u, err := c.baseURL.Parse(path)
	if err != nil {
		return err
	}
	var bodyReader io.Reader
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}
	req, err := http.NewRequest(method, u.String(), bodyReader)
	if err != nil {
		return err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	expectedStatusCode := map[string]int{
		"POST":   200,
		"PUT":    200,
		"GET":    200,
		"DELETE": 204,
	}
	if resp.StatusCode != expectedStatusCode[method] {
		return fmt.Errorf("HTTP %s request error. Response code: %d", method, resp.StatusCode)
	}
	if result != nil {
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) URL(path string) string {
	u, err := c.publicURL.Parse(path)
	if err != nil {
		return ""
	}
	return u.String()
}
