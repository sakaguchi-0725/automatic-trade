package util

import (
	"encoding/json"
	"fmt"
)

func DecodeJSON[T any](input any) (*T, error) {
	data, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal input data: %w", err)
	}

	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return &result, nil
}
