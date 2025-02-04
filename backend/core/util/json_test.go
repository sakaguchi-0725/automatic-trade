package util_test

import (
	"automatic-trade/backend/core/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SampleStruct struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func TestDecodeJSON(t *testing.T) {
	tests := map[string]struct {
		input       any
		expected    *SampleStruct
		expectedErr bool
	}{
		"valid struct input": {
			input: map[string]any{
				"id":   1,
				"name": "Alice",
			},
			expected:    &SampleStruct{ID: 1, Name: "Alice"},
			expectedErr: false,
		},
		"valid map input": {
			input: map[string]any{
				"id":   2,
				"name": "Bob",
			},
			expected:    &SampleStruct{ID: 2, Name: "Bob"},
			expectedErr: false,
		},
		"invalid json string": {
			input:       `{"id": 3, "name": "Charlie"`,
			expected:    nil,
			expectedErr: true,
		},
		"invalid data type": {
			input:       42,
			expected:    nil,
			expectedErr: true,
		},
		"empty map": {
			input:       map[string]any{},
			expected:    &SampleStruct{},
			expectedErr: false,
		},
		"nil input": {
			input:       nil,
			expected:    &SampleStruct{},
			expectedErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := util.DecodeJSON[SampleStruct](tt.input)

			if tt.expectedErr {
				assert.Error(t, err, "expected error for test case: %s", name)
			} else {
				assert.NoError(t, err, "unexpected error for test case: %s", name)
				assert.Equal(t, tt.expected, result, "mismatch for test case: %s", name)
			}
		})
	}
}
