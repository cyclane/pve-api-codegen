package pve

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
)

// PveApiError is an error implementation for PVE API Client errors.
type PveApiError struct {
	Status   int
	Response *http.Response
}

func (pae PveApiError) Error() string {
	return fmt.Sprintf("pve api error %d (%s)", pae.Status, http.StatusText(pae.Status))
}

// PVE API Client
type Client struct {
	httpClient    *http.Client
	userAgent     string
	apiUrl        string
	authenticator func(*http.Header)
	logger        LeveledLogger
}

// UseClient sets the underlying http.Client used for requests.
func (c *Client) UseClient(httpClient *http.Client) *Client {
	c.httpClient = httpClient
	return c
}

// UseLogger sets the logger used by the PVE API Client.
func (c *Client) UseLogger(logger LeveledLogger) *Client {
	c.logger = logger
	return c
}

// UseUserAgent sets the UserAgent header used for requests.
func (c *Client) UseUserAgent(userAgent string) *Client {
	c.userAgent = userAgent
	return c
}

// Request makes an API request through the PVE API Client.
func (c *Client) Request(method, endpoint string, data, v any) error {
	fullPath := path.Join(c.apiUrl, endpoint)
	c.logger.Debugf("request: %s %s", method, fullPath)

	var body io.Reader
	if data != nil {
		dataB, err := json.Marshal(data)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(dataB)
	}

	req, err := http.NewRequest(method, fullPath, body)
	if err != nil {
		return err
	}

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	if c.authenticator != nil {
		c.authenticator(&req.Header)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return PveApiError{
			Status:   res.StatusCode,
			Response: res,
		}
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var datakey map[string]json.RawMessage
	if err := json.Unmarshal(resBody, &datakey); err != nil {
		return err
	}

	if resData, ok := datakey["data"]; ok {
		return json.Unmarshal(resData, &v)
	}

	return json.Unmarshal(resBody, &v)
}

// NewClient creates a new PVE API Client.
func NewClient(apiUrl string) *Client {
	return &Client{
		httpClient: http.DefaultClient,
		userAgent:  "example.com/yee",
		apiUrl:     apiUrl,
		logger:     NewDefaultLogger(LevelWarn),
	}
}
