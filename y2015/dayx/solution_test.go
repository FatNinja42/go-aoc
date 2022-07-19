package main

import (
	"aoc/cmd"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	var tests = map[string]struct {
		input    string
		expected interface{}
	}{
		"1": {
			input:    "",
			expected: 0,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := dayx{}
			in := cmd.Input{SI: tt.input}

			res := d.Part1(in)

			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestPart2(t *testing.T) {
	var tests = map[string]struct {
		input    string
		expected interface{}
	}{
		"1": {
			input:    "",
			expected: 0,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := dayx{}
			in := cmd.Input{SI: tt.input}

			res := d.Part2(in)

			assert.Equal(t, tt.expected, res)
		})
	}
}
