package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeUtil(t *testing.T) {
	t.Run("unix to jst", func(t *testing.T) {
		tests := map[string]struct {
			unix     int64
			isMilli  bool
			expected string
		}{
			"Unix seconds - 2025-02-02 15:00:00 JST": {
				unix:     1738476000, // 2025-02-02 15:00:00 JST (Unix seconds)
				isMilli:  false,
				expected: "2025-02-02T15:00:00+09:00",
			},
			"Unix milliseconds - 2025-02-02 15:00:00 JST": {
				unix:     1738476000000, // 2025-02-02 15:00:00 JST (Unix milliseconds)
				isMilli:  true,
				expected: "2025-02-02T15:00:00+09:00",
			},
			"Unix seconds - Future date (2040-01-01 00:00:00 JST)": {
				unix:     2208956400, // 2040-01-01 00:00:00 JST (Unix seconds)
				isMilli:  false,
				expected: "2040-01-01T00:00:00+09:00",
			},
			"Unix milliseconds - Future date (2040-01-01 00:00:00 JST)": {
				unix:     2208956400000, // 2040-01-01 00:00:00 JST (Unix milliseconds)
				isMilli:  true,
				expected: "2040-01-01T00:00:00+09:00",
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				got := UnixToJST(tt.unix, tt.isMilli)
				gotFormatted := got.Format(time.RFC3339)

				assert.Equal(t, tt.expected, gotFormatted)
			})
		}
	})
}
