package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnpackString(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    "a6",
			expected: "aaaaaa",
		},
	}
	for _, c := range cases {
		actual := unpackString(c.input)
		assert.EqualValues(t, c.expected, actual)
	}
}
