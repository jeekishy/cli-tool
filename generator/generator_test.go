package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateEvenNumbers(t *testing.T) {
	limit := 5
	expectEvenNumber := 2
	count := 0

	for n := range GenerateEvenNumbers(limit) {
		assert.Equal(t, expectEvenNumber, n, "unexpected even number")
		expectEvenNumber += 2
		count++ // keep track of how many times we are iterating
	}

	assert.Equal(t, limit, count)
}

func TestGenerateEvenNumbersWithZeroLimit(t *testing.T) {
	limit := 0
	expectEvenNumber := 2
	count := 0

	for n := range GenerateEvenNumbers(0) {
		assert.Equal(t, expectEvenNumber, n, "unexpected even number")
		expectEvenNumber += 2
		count++ // keep track of how many times we are iterating
	}

	assert.Equal(t, limit, count)
}
