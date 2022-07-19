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
			input:    ">",
			expected: 2,
		},
		"2": {
			input:    "^>v<",
			expected: 4,
		},
		"3": {
			input:    "^v^v^v^v^v",
			expected: 2,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := day3{}
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
			input:    ">v",
			expected: 3,
		},
		"2": {
			input:    "^>v<",
			expected: 3,
		},
		"3": {
			input:    "^v^v^v^v^v",
			expected: 11,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := day3{}
			in := cmd.Input{SI: tt.input}

			res := d.Part2(in)

			assert.Equal(t, tt.expected, res)
		})
	}
}
