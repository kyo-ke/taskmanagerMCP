package tools

import "encoding/json"

func jsonMarshal(v any) ([]byte, error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return nil, err
	}
	return b, nil
}
