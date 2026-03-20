package resource

import (
	"github.com/mark3labs/mcp-go/server"

	"taskmanagerMCP/internal/taskmanager"
)

// RegisterAll registers all MCP resources exposed by this server.
func RegisterAll(s *server.MCPServer, client *taskmanager.Client) {
	registerTaskGetResource(s, client)
	registerTasksByUserResource(s, client)
}
