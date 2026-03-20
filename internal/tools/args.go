package tools

import (
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func argumentsMap(req mcp.CallToolRequest) (map[string]any, error) {
	if req.Params.Arguments == nil {
		return map[string]any{}, nil
	}
	args, ok := req.Params.Arguments.(map[string]any)
	if ok {
		return args, nil
	}
	return nil, fmt.Errorf("arguments must be an object")
}

func requireString(req mcp.CallToolRequest, key string) (string, error) {
	args, err := argumentsMap(req)
	if err != nil {
		return "", err
	}
	v, ok := args[key]
	if !ok || v == nil {
		return "", fmt.Errorf("missing argument: %s", key)
	}
	s, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("argument %s must be string", key)
	}
	if s == "" {
		return "", fmt.Errorf("argument %s must be non-empty", key)
	}
	return s, nil
}

func requireNumber(req mcp.CallToolRequest, key string) (float64, error) {
	args, err := argumentsMap(req)
	if err != nil {
		return 0, err
	}
	v, ok := args[key]
	if !ok || v == nil {
		return 0, fmt.Errorf("missing argument: %s", key)
	}
	// JSON numbers are decoded as float64
	if f, ok := v.(float64); ok {
		return f, nil
	}
	// try int-ish
	if i, ok := v.(int); ok {
		return float64(i), nil
	}
	return 0, fmt.Errorf("argument %s must be number", key)
}
