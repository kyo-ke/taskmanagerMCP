package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"taskmanagerMCP/internal/taskmanager"
)

func registerTaskDelete(s *server.MCPServer, client *taskmanager.Client) {
	tool := mcp.NewTool("task_delete",
		mcp.WithDescription("Delete a task by id"),
		mcp.WithString("id", mcp.Required(), mcp.Description("Task UUID")),
	)

	s.AddTool(tool, func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := requireString(req, "id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		if err := client.DeleteTask(ctx, id); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText("deleted"), nil
	})
}
