package strconv

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleFormatSliceToCSV() {
	s := []string{"foo", "bar", "baz"}
	csv := FormatSliceToCSV(s)
	fmt.Println(csv)
	// Output: foo,bar,baz
}

func TestFormatSliceToCSV(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		s        []string
		expected string
	}{
		{
			name:     "empty",
			s:        []string{},
			expected: "",
		},
		{
			name:     "single",
			s:        []string{"foo"},
			expected: "foo",
		},
		{
			name:     "multiple",
			s:        []string{"foo", "bar", "baz"},
			expected: "foo,bar,baz",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actual := FormatSliceToCSV(test.s)
			assert.Equal(t, test.expected, actual)
		})
	}
}
