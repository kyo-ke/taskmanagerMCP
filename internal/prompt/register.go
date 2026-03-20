package prompt

import (
	"github.com/mark3labs/mcp-go/server"

	"taskmanagerMCP/internal/taskmanager"
)

// RegisterAll registers all MCP prompts exposed by this server.
func RegisterAll(s *server.MCPServer, client *taskmanager.Client) {
	registerCreateTaskPrompt(s, client)
}
