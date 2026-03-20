package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"taskmanagerMCP/internal/taskmanager"
)

func registerTasksByUserResource(s *server.MCPServer, client *taskmanager.Client) {
	res := mcp.NewResource(
		"taskmanager://users/{userId}/tasks",
		"tasks_by_user",
		mcp.WithResourceDescription("Tasks assigned to a user (JSON array)"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(res, func(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		userID, err := userIDFromURI(req.Params.URI)
		if err != nil {
			return nil, err
		}
		out, err := client.ListTasksByUser(ctx, userID)
		if err != nil {
			return nil, err
		}
		b, _ := json.MarshalIndent(out, "", "  ")
		return []mcp.ResourceContents{mcp.TextResourceContents{URI: req.Params.URI, Text: string(b), MIMEType: "application/json"}}, nil
	})
}

func userIDFromURI(uri string) (string, error) {
	if !strings.HasPrefix(uri, taskURIScheme) {
		return "", fmt.Errorf("unsupported uri: %s", uri)
	}
	path := strings.TrimPrefix(uri, taskURIScheme)
	path = strings.TrimPrefix(path, "/")
	if !strings.HasPrefix(path, "users/") {
		return "", fmt.Errorf("unsupported uri path: %s", uri)
	}
	path = strings.TrimPrefix(path, "users/")
	parts := strings.SplitN(path, "/", 2)
	if len(parts) != 2 {
		return "", fmt.Errorf("unsupported uri path: %s", uri)
	}
	userID := parts[0]
	if userID == "" {
		return "", fmt.Errorf("missing userId in uri")
	}
	// decode, in case the client URL-encoded it.
	if decoded, err := url.PathUnescape(userID); err == nil {
		userID = decoded
	}
	if parts[1] != "tasks" {
		return "", fmt.Errorf("unsupported uri path: %s", uri)
	}
	return userID, nil
}
