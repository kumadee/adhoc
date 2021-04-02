package multipersist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNaiveMultiPersitence(t *testing.T) {
	type Expected struct {
		value uint64
		err   error
	}
	var cases = []struct {
		input    uint64
		expected Expected
	}{
		{
			input:    9,
			expected: Expected{0, nil},
		},
		{
			input:    10,
			expected: Expected{1, nil},
		},
		{
			input:    2718,
			expected: Expected{2, nil},
		},
		{
			input:    77,
			expected: Expected{4, nil},
		},
	}
	for _, tc := range cases {
		description := fmt.Sprintf("F(%d)", tc.input)
		v, err := NaiveMultiPersitence(tc.input)
		assert.Equal(t, tc.expected.value, v, description)
		assert.ErrorIs(t, tc.expected.err, err, description)
	}
}

func TestMinMaxHighestPersistNumbers(t *testing.T) {
	type Expected struct {
		min uint64
		max uint64
		per uint64
	}
	var cases = []struct {
		start    uint64
		end      uint64
		function Persistence
		expected Expected
	}{
		{
			start:    0,
			end:      100,
			function: NaiveMultiPersitence,
			expected: Expected{77, 77, 4},
		},
	}
	for _, tc := range cases {
		min, max, per := MinMaxHighestPersistNumbers(tc.start, tc.end, tc.function)
		assert.Equal(t, tc.expected.min, min, fmt.Sprintf("min F(%d, %d) = %d", tc.start, tc.end, min))
		assert.Equal(t, tc.expected.max, max, fmt.Sprintf("max F(%d, %d) = %d", tc.start, tc.end, max))
		assert.Equal(t, tc.expected.per, per, fmt.Sprintf("per F(%d, %d) = %d", tc.start, tc.end, per))
	}
}

func BenchmarkNaiveMultiPersitence(b *testing.B) {
	for i := uint64(0); i <= uint64(b.N); i++ {
		NaiveMultiPersitence(i)
	}
}
