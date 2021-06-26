package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Gom hết test case vào một function thay vì tạo từng func unit test riêng biệt
func TestRound(t *testing.T) {
	scenarios := []struct {
		desc     string
		in       float32
		expected int
	}{
		{
			desc:     "round down",
			in:       1.1,
			expected: 1,
		},
		{
			desc:     "round up",
			in:       3.7,
			expected: 4,
		},
		{
			desc:     "unchanged",
			in:       6.0,
			expected: 6,
		},
	}

	for _, item := range scenarios {
		in := float32(item.in)

		result := Round(in)
		assert.Equal(t, item.expected, result)
	}
}
