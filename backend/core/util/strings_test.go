package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringUtil(t *testing.T) {
	t.Run("StringToInt64", func(t *testing.T) {
		tests := map[string]struct {
			input    string
			expected int64
			hasError bool
		}{
			"Valid integer":       {"12345", 12345, false},
			"Zero":                {"0", 0, false},
			"Negative integer":    {"-9876", -9876, false},
			"Large integer":       {"9223372036854775807", 9223372036854775807, false},   // int64最大値
			"Small integer":       {"-9223372036854775808", -9223372036854775808, false}, // int64最小値
			"Empty string":        {"", 0, true},
			"Non-numeric string":  {"hello", 0, true},
			"Float number string": {"3.14", 0, true},
			"Overflowing integer": {"9223372036854775808", 0, true},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				result, err := StringToInt64(tt.input)
				if tt.hasError {
					assert.Error(t, err, "Expected error for input: %s", tt.input)
				} else {
					assert.NoError(t, err, "Unexpected error for input: %s", tt.input)
					assert.Equal(t, tt.expected, result, "Mismatch for input: %s", tt.input)
				}
			})
		}
	})

	t.Run("StringToFloat64", func(t *testing.T) {
		tests := map[string]struct {
			input    string
			expected float64
			hasError bool
		}{
			"Valid float":        {"123.45", 123.45, false},
			"Valid integer":      {"123", 123.0, false},
			"Zero":               {"0", 0.0, false},
			"Negative float":     {"-987.65", -987.65, false},
			"Large float":        {"1.7976931348623157e+308", 1.7976931348623157e+308, false},   // float64最大値
			"Small float":        {"-1.7976931348623157e+308", -1.7976931348623157e+308, false}, // float64最小値
			"Empty string":       {"", 0.0, true},
			"Non-numeric string": {"world", 0.0, true},
			"Multiple dots":      {"3.14.159", 0.0, true},
			"Overflowing float":  {"1e309", 0.0, true},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				result, err := StringToFloat64(tt.input)
				if tt.hasError {
					assert.Error(t, err, "Expected error for input: %s", tt.input)
				} else {
					assert.NoError(t, err, "Unexpected error for input: %s", tt.input)
					assert.Equal(t, tt.expected, result, "Mismatch for input: %s", tt.input)
				}
			})
		}
	})
}
