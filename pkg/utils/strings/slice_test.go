package strings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleSliceToCSV() {
	s := []string{"foo", "bar", "baz"}
	csv := SliceToCSV(s)
	fmt.Println(csv)
	// Output: foo,bar,baz
}

func TestSliceToCSVStrings(t *testing.T) {
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

			actual := SliceToCSV(test.s)
			assert.Equal(t, test.expected, actual)
		})
	}
}
