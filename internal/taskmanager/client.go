package taskmanager

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	baseURL string
	hc      *http.Client
}

func NewClient(baseURL string) *Client {
	baseURL = strings.TrimRight(baseURL, "/")
	return &Client{
		baseURL: baseURL,
		hc: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

type Task struct {
	ID             string `json:"id,omitempty"`
	Priority       int    `json:"priority"`
	AssignedUserID string `json:"assignedUserId"`
	Status         string `json:"status"`
	Description    string `json:"description"`
	CreatedAt      string `json:"createdAt,omitempty"`
	UpdatedAt      string `json:"updatedAt,omitempty"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func (c *Client) CreateTask(ctx context.Context, in Task) (Task, error) {
	var out Task
	if err := c.doJSON(ctx, http.MethodPost, "/tasks", in, &out); err != nil {
		return Task{}, err
	}
	return out, nil
}

func (c *Client) GetTask(ctx context.Context, id string) (Task, error) {
	var out Task
	if err := c.doJSON(ctx, http.MethodGet, "/tasks/"+url.PathEscape(id), nil, &out); err != nil {
		return Task{}, err
	}
	return out, nil
}

func (c *Client) UpdateTask(ctx context.Context, id string, in Task) (Task, error) {
	var out Task
	if err := c.doJSON(ctx, http.MethodPut, "/tasks/"+url.PathEscape(id), in, &out); err != nil {
		return Task{}, err
	}
	return out, nil
}

func (c *Client) DeleteTask(ctx context.Context, id string) error {
	return c.doJSON(ctx, http.MethodDelete, "/tasks/"+url.PathEscape(id), nil, nil)
}

func (c *Client) ListTasksByUser(ctx context.Context, userID string) ([]Task, error) {
	var out []Task
	if err := c.doJSON(ctx, http.MethodGet, "/users/"+url.PathEscape(userID)+"/tasks", nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) doJSON(ctx context.Context, method, path string, in any, out any) error {
	var body io.Reader
	if in != nil {
		b, err := json.Marshal(in)
		if err != nil {
			return err
		}
		body = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, body)
	if err != nil {
		return err
	}
	if in != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// No content
	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var er errorResponse
		_ = json.Unmarshal(b, &er)
		if er.Error != "" {
			return fmt.Errorf("taskmanager http %d: %s", resp.StatusCode, er.Error)
		}
		return fmt.Errorf("taskmanager http %d: %s", resp.StatusCode, strings.TrimSpace(string(b)))
	}

	if out == nil {
		return nil
	}
	return json.Unmarshal(b, out)
}
