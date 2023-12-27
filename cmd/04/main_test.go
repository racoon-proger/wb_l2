package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	result := searchAnagrams([]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"})
	expected := map[string][]string{
		"столик": {"листок", "слиток", "столик"},
	}
	assert.Equal(t, expected, result, expected)
}
