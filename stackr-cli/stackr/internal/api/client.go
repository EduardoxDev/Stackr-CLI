package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const baseURL = "https://api.stackr.lat/v1"

type Client struct {
	token      string
	httpClient *http.Client
}

func New(token string) *Client {
	return &Client{token: token, httpClient: &http.Client{}}
}

func (c *Client) do(method, endpoint string, body interface{}) ([]byte, error) {
	var reader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewReader(data)
	}
	req, err := http.NewRequest(method, baseURL+endpoint, reader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("User-Agent", "stackr-cli/1.2.0")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		var e map[string]interface{}
		if json.Unmarshal(data, &e) == nil {
			for _, k := range []string{"message", "error", "msg"} {
				if msg, ok := e[k].(string); ok && msg != "" {
					return nil, fmt.Errorf("%s", msg)
				}
			}
		}
		return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
	}
	return data, nil
}

func (c *Client) Get(ep string) ([]byte, error)                  { return c.do("GET", ep, nil) }
func (c *Client) Post(ep string, b interface{}) ([]byte, error)  { return c.do("POST", ep, b) }
func (c *Client) Patch(ep string, b interface{}) ([]byte, error) { return c.do("PATCH", ep, b) }
func (c *Client) Put(ep string, b interface{}) ([]byte, error)   { return c.do("PUT", ep, b) }
func (c *Client) Delete(ep string) ([]byte, error)               { return c.do("DELETE", ep, nil) }

func (c *Client) UploadZip(endpoint, zipPath string) ([]byte, error) {
	f, err := os.Open(zipPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	fw, err := w.CreateFormFile("file", filepath.Base(zipPath))
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(fw, f); err != nil {
		return nil, err
	}
	w.Close()
	req, err := http.NewRequest("POST", baseURL+endpoint, &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Set("User-Agent", "stackr-cli/1.2.0")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		var e map[string]interface{}
		if json.Unmarshal(data, &e) == nil {
			if msg, ok := e["message"].(string); ok {
				return nil, fmt.Errorf("%s", msg)
			}
		}
		return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
	}
	return data, nil
}

func (c *Client) StreamGet(endpoint string) (io.ReadCloser, error) {
	req, err := http.NewRequest("POST", baseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("User-Agent", "stackr-cli/1.2.0")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		resp.Body.Close()
		return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
	}
	return resp.Body, nil
}

func Decode(data []byte, v interface{}) error { return json.Unmarshal(data, v) }
