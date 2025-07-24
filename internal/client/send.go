package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *client) send(ctx context.Context, path string, method string, body io.Reader, out any) error {
	req, err := http.NewRequestWithContext(ctx, method, c.rootURL+path, body)
	if err != nil {
		return err
	}
	req.Header.Add("X-API-KEY", c.apiKey)
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("code: %s, response: %s", resp.Status, string(result))
	}
	return json.Unmarshal(result, out)
}
