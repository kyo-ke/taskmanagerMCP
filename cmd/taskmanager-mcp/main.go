package main

import (
	"fmt"
	"os"

	"github.com/mark3labs/mcp-go/server"

	"taskmanagerMCP/internal/prompt"
	"taskmanagerMCP/internal/resource"
	"taskmanagerMCP/internal/taskmanager"
	"taskmanagerMCP/internal/tools"
)

func main() {
	baseURL := envOr("TASKMANAGER_HTTP_BASE_URL", "http://localhost:8080")
	client := taskmanager.NewClient(baseURL)

	s := server.NewMCPServer(
		"taskmanager-mcp",
		"0.1.0",
		server.WithPromptCapabilities(false),
		server.WithToolCapabilities(false),
		server.WithRecovery(),
	)

	tools.RegisterAll(s, client)
	resource.RegisterAll(s, client)
	prompt.RegisterAll(s, client)

	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
