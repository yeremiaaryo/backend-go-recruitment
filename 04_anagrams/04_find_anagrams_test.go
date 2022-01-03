package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Main(t *testing.T) {
	main()
}

func Test_findAnagrams(t *testing.T) {
	resp := findAnagrams([]string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"})
	expectation := map[string][]string{
		"aikt":  {"kita", "atik", "tika"},
		"aku":   {"aku", "kua"},
		"aik":   {"kia"},
		"aakmn": {"makan"},
	}
	assert.Equal(t, expectation, resp)
}
