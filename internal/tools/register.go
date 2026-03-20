package tools

import (
	"taskmanagerMCP/internal/taskmanager"

	"github.com/mark3labs/mcp-go/server"
)

// RegisterAll registers all MCP tools exposed by this server.
func RegisterAll(s *server.MCPServer, client *taskmanager.Client) {
	registerTaskCreate(s, client)
	registerTaskUpdate(s, client)
	registerTaskDelete(s, client)
}
