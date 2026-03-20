package prompt

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"taskmanagerMCP/internal/taskmanager"
)

// registerCreateTaskPrompt registers a prompt that guides an LLM to create a task
// using the task_create tool with sensible defaults.
//
// Defaults:
// - priority: 3
// - assignedUserId: anonymous
// - status: todo
func registerCreateTaskPrompt(s *server.MCPServer, _ *taskmanager.Client) {
	p := mcp.NewPrompt("task_create_default",
		mcp.WithPromptDescription("Create a new task (defaults: priority=3, user=anonymous, status=todo)."),
		mcp.WithArgument("description", mcp.ArgumentDescription("Task description"), mcp.RequiredArgument()),
		mcp.WithArgument("priority", mcp.ArgumentDescription("Priority (default: 3)")),
		mcp.WithArgument("assignedUserId", mcp.ArgumentDescription("Assigned user id (default: anonymous)")),
		mcp.WithArgument("status", mcp.ArgumentDescription("todo|in_progress|done (default: todo)")),
	)

	s.AddPrompt(p, func(ctx context.Context, req mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
		_ = ctx

		args := req.Params.Arguments
		if args == nil {
			args = map[string]string{}
		}

		description := args["description"]
		priority := args["priority"]
		assignedUserId := args["assignedUserId"]
		status := args["status"]

		if priority == "" {
			priority = "3"
		}
		if assignedUserId == "" {
			assignedUserId = "anonymous"
		}
		if status == "" {
			status = "todo"
		}

		// Prompt makes it easy for the model to call `task_create` with defaults.
		return &mcp.GetPromptResult{
			Description: "Create a task using task_create with defaults.",
			Messages: []mcp.PromptMessage{
				{
					Role: mcp.RoleUser,
					Content: mcp.TextContent{
						Type: "text",
						Text: "Create a task in taskmanager. Use the task_create tool with these fields.\n\n" +
							"priority: " + priority + "\n" +
							"assignedUserId: " + assignedUserId + "\n" +
							"status: " + status + "\n" +
							"description: " + description + "\n",
					},
				},
			},
		}, nil
	})
}
