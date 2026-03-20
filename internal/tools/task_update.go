package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"taskmanagerMCP/internal/taskmanager"
)

func registerTaskUpdate(s *server.MCPServer, client *taskmanager.Client) {
	tool := mcp.NewTool("task_update",
		mcp.WithDescription("Update a task by id"),
		mcp.WithString("id", mcp.Required(), mcp.Description("Task UUID")),
		mcp.WithNumber("priority", mcp.Required(), mcp.Description("Priority")),
		mcp.WithString("assignedUserId", mcp.Required(), mcp.Description("Assigned user id")),
		mcp.WithString("status", mcp.Required(), mcp.Description("todo|in_progress|done")),
		mcp.WithString("description", mcp.Required(), mcp.Description("Task description")),
	)

	s.AddTool(tool, func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := requireString(req, "id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		priority, err := requireNumber(req, "priority")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		assignedUserID, err := requireString(req, "assignedUserId")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		status, err := requireString(req, "status")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		description, err := requireString(req, "description")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		out, err := client.UpdateTask(ctx, id, taskmanager.Task{
			Priority:       int(priority),
			AssignedUserID: assignedUserID,
			Status:         status,
			Description:    description,
		})
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		b, _ := jsonMarshal(out)
		return mcp.NewToolResultText(string(b)), nil
	})
}
