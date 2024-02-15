package api

import (
	"encoding/json"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKeepAliveParsing(t *testing.T) {
	tests := []struct {
		name string
		req  *ChatRequest
		exp  *Duration
	}{
		{
			name: "Positive Duration",
			req: &ChatRequest{
				KeepAlive: &Duration{Duration: 42 * time.Minute},
			},
			exp: &Duration{42 * time.Minute},
		},
		{
			name: "Negative Duration",

			req: &ChatRequest{
				KeepAlive: &Duration{Duration: -1 * time.Minute},
			},
			exp: &Duration{math.MaxInt64},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ser, err := json.Marshal(test.req)
			require.NoError(t, err)

			var dec ChatRequest
			err = json.Unmarshal([]byte(ser), &dec)
			require.NoError(t, err)

			assert.Equal(t, test.exp, dec.KeepAlive)
		})
	}
}
