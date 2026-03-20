package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"taskmanagerMCP/internal/taskmanager"
)

// URI scheme used by this server for task resources.
// Example: taskmanager://tasks/<uuid>
const taskURIScheme = "taskmanager://"

func registerTaskGetResource(s *server.MCPServer, client *taskmanager.Client) {
	// Resource:
	// - uri: a stable identifier for the resource.
	// - name: short display name.
	// - description/mimeType: help clients render and understand the content.
	res := mcp.NewResource(
		"taskmanager://tasks/{id}",
		"task",
		mcp.WithResourceDescription("A single task information based on id"),
		mcp.WithMIMEType("application/json"),
	)

	// Resource handler:
	// - request.Params.URI contains the concrete URI the client is requesting.
	// - return []ResourceContents (we return one TextResourceContents that contains JSON).
	s.AddResource(res, func(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		id, err := taskIDFromURI(req.Params.URI)
		if err != nil {
			return nil, err
		}
		out, err := client.GetTask(ctx, id)
		if err != nil {
			return nil, err
		}

		b, _ := json.MarshalIndent(out, "", "  ")
		return []mcp.ResourceContents{mcp.TextResourceContents{URI: req.Params.URI, Text: string(b), MIMEType: "application/json"}}, nil
	})
}

func taskIDFromURI(uri string) (string, error) {
	// expected: taskmanager://tasks/<id>
	if !strings.HasPrefix(uri, taskURIScheme) {
		return "", fmt.Errorf("unsupported uri: %s", uri)
	}
	path := strings.TrimPrefix(uri, taskURIScheme)
	path = strings.TrimPrefix(path, "/")
	if !strings.HasPrefix(path, "tasks/") {
		return "", fmt.Errorf("unsupported uri path: %s", uri)
	}
	id := strings.TrimPrefix(path, "tasks/")
	if id == "" {
		return "", fmt.Errorf("missing task id in uri")
	}
	return id, nil
}
