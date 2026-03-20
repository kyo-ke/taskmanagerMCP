package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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

	addr := envOr("TASKMANAGER_MCP_ADDR", ":8081")

	// mcp-go provides an SSE server that can be mounted on a regular net/http server.
	// Default endpoints (mcp-go v0.45.0):
	// - GET  /sse     : Server-Sent Events stream
	// - POST /message : client -> server JSON-RPC messages
	sse := server.NewSSEServer(s)

	h := http.NewServeMux()
	h.Handle(sse.CompleteSsePath(), sse.SSEHandler())
	h.Handle(sse.CompleteMessagePath(), sse.MessageHandler())
	// Optional: expose root for debugging / human-friendly 404s.
	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("not found\nexpected endpoints:\n- GET  " + sse.CompleteSsePath() + "\n- POST " + sse.CompleteMessagePath() + "\n"))
	})

	httpSrv := &http.Server{
		Addr:              addr,
		Handler:           h,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("taskmanager-mcp (http+sse) listening on %s", addr)
	log.Printf("mcp endpoints: GET %s, POST %s", sse.CompleteSsePath(), sse.CompleteMessagePath())
	if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
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
