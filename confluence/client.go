package confluence

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Client provides a connection to the Confluence API
type Client struct {
	client    *http.Client
	baseURL   *url.URL
	publicURL *url.URL
}

// NewClientInput provides information to connect to the Confluence API
type NewClientInput struct {
	site  string
	user  string
	token string
}

// ErrorResponse describes why a request failed
type ErrorResponse struct {
	StatusCode int `json:"statusCode,omitempty"`
	Data       struct {
		Authorized bool     `json:"authorized,omitempty"`
		Valid      bool     `json:"valid,omitempty"`
		Errors     []string `json:"errors,omitempty"`
		Successful bool     `json:"successful,omitempty"`
	} `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

// NewClient returns an authenticated client ready to use
func NewClient(input *NewClientInput) *Client {
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
	}
}

// Get uses the client to send a GET request
func (c *Client) Get(path string, result interface{}) error {
	body := new(bytes.Buffer)
	return c.do("GET", path, "", body, result)
}

// Delete uses the client to send a DELETE request
func (c *Client) Delete(path string) error {
	body := new(bytes.Buffer)
	return c.do("DELETE", path, "", body, nil)
}

// Post uses the client to send a POST request
func (c *Client) Post(path string, body interface{}, result interface{}) error {
	b, err := jsonBytesBuffer(body)
	if err != nil {
		return err
	}
	return c.do("POST", path, "application/json", b, result)
}

// Put uses the client to send a PUT request
func (c *Client) Put(path string, body interface{}, result interface{}) error {
	b, err := jsonBytesBuffer(body)
	if err != nil {
		return err
	}
	return c.do("PUT", path, "application/json", b, result)
}

// PostForm uses the client to send a multi-part form POST request
func (c *Client) PostForm(path, filename, body string, result interface{}) error {
	b, ct, err := formBytesBuffer(filename, body)
	if err != nil {
		return err
	}
	return c.do("POST", path, ct, b, result)
}

// PutForm uses the client to send a multi-part form PUT request
func (c *Client) PutForm(path, filename, body string, result interface{}) error {
	b, ct, err := formBytesBuffer(filename, body)
	if err != nil {
		return err
	}
	return c.do("PUT", path, ct, b, result)
}

func jsonBytesBuffer(body interface{}) (*bytes.Buffer, error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(bodyBytes), nil
}

// formBytesBuffer returns the body as a multi-part form and the content type
func formBytesBuffer(filename, body string) (*bytes.Buffer, string, error) {
	bodyBytes := &bytes.Buffer{}
	writer := multipart.NewWriter(bodyBytes)
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return nil, "", err
	}
	_, err = io.WriteString(part, body)
	if err != nil {
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return bodyBytes, writer.FormDataContentType(), nil
}

// do uses the client to send a specified request
func (c *Client) do(method, path, contentType string, body *bytes.Buffer, result interface{}) error {
	u, err := c.baseURL.Parse(path)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return err
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	req.Header.Add("X-Atlassian-Token", "nocheck")
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
		var responseBody string
		var errResponse ErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&errResponse)
		if err != nil {
			responseBody = "Could not decode error"
		} else {
			responseBody = errResponse.String()
		}
		s := body.String()
		return fmt.Errorf("%s\n\n%s %s\n%s\n\n%s",
			resp.Status, method, path, s, responseBody)
	}
	if result != nil {
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *ErrorResponse) String() string {
	d := e.Data
	var errorsString string
	if len(d.Errors) > 0 {
		errorsString = fmt.Sprintf("\n  * %s", strings.Join(d.Errors, "\n  * "))
	}
	return fmt.Sprintf("%s\nAuthorized: %t\nValid: %t\nSuccessful: %t%s",
		e.Message, d.Authorized, d.Valid, d.Successful, errorsString)
}

// URL returns the public URL for a given path
func (c *Client) URL(path string) string {
	u, err := c.publicURL.Parse(path)
	if err != nil {
		return ""
	}
	return u.String()
}
