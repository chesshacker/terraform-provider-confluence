package main

import (
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
	instance string
	user     string
	token    string
}

func NewClient(input *NewClientInput) (*Client, error) {

	// TODO: Is there a better way to build the *URL.url object? I'm worried that
	// input.user contains an @ symbol... but it seems to work as-is.

	u := "https://" + input.user + ":" + input.token + "@" + input.instance + ".atlassian.net"
	baseURL, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	u = "https://" + input.instance + ".atlassian.net"
	publicURL, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	return &Client{
		client: &http.Client{
			Timeout: time.Second * 10,
		},
		baseURL:   baseURL,
		publicURL: publicURL,
	}, nil
}

func (c *Client) Post(path string, body io.Reader) (*http.Response, error) {
	return c.do("POST", path, body)
}

func (c *Client) Get(path string) (*http.Response, error) {
	return c.do("GET", path, nil)
}

func (c *Client) Put(path string, body io.Reader) (*http.Response, error) {
	return c.do("PUT", path, body)
}

func (c *Client) Delete(path string) (*http.Response, error) {
	return c.do("DELETE", path, nil)
}

func (c *Client) do(method string, path string, body io.Reader) (*http.Response, error) {
	u, err := c.baseURL.Parse(path)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return c.client.Do(req)
}

func (c *Client) URL(path string) string {
	u, err := c.publicURL.Parse(path)
	if err != nil {
		return ""
	}
	return u.String()
}
