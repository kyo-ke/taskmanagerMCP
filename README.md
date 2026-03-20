# taskmanagerMCP

An MCP server for the `taskmanager` app, built with [`mark3labs/mcp-go`](https://github.com/mark3labs/mcp-go).

This server exposes tools that call the `taskmanager` HTTP API.

## Tools

- `task_create`
- `task_get`
- `task_update`
- `task_delete`
- `task_list_by_user`

## Configuration

Environment variables:

- `TASKMANAGER_HTTP_BASE_URL` (default: `http://localhost:8080`)
- `TASKMANAGER_MCP_ADDR` (default: `:8081`, used by the HTTP/SSE server only)

## Run

### Stdio mode (default)

```zsh
cd taskmanagerMCP

go run ./cmd/taskmanager-mcp
```

> Start the `taskmanager` HTTP server first.

### HTTP(SSE) mode

This starts an HTTP server (default `:8081`) and serves MCP over Server-Sent Events.

```zsh
cd taskmanagerMCP

TASKMANAGER_MCP_ADDR=:8081 go run ./cmd/taskmanager-mcp-http
```
