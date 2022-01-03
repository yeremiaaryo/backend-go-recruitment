package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Main(t *testing.T) {
	main()
}

func Test_findFirstStringInBracket(t *testing.T) {
	t.Run("When no open bracket found", func(t *testing.T) {
		result := findFirstStringInBracket("Hello world)")
		expected := ""
		assert.Equal(t, expected, result)
	})

	t.Run("When no close bracket found", func(t *testing.T) {
		result := findFirstStringInBracket("Hello (world")
		expected := ""
		assert.Equal(t, expected, result)
	})

	t.Run("When no bracket found", func(t *testing.T) {
		result := findFirstStringInBracket("Hello world")
		expected := ""
		assert.Equal(t, expected, result)
	})

	t.Run("When both bracket found", func(t *testing.T) {
		result := findFirstStringInBracket("Hello (world)")
		expected := "world"
		assert.Equal(t, expected, result)
	})
}
