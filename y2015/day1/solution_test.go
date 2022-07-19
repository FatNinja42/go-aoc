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
			input:    "(())",
			expected: 0,
		},
		"2": {
			input:    "()()",
			expected: 0,
		},
		"3": {
			input:    "(((",
			expected: 3,
		},
		"4": {
			input:    "(()(()(",
			expected: 3,
		},
		"5": {
			input:    "())",
			expected: -1,
		},
		"6": {
			input:    "))(",
			expected: -1,
		},
		"7": {
			input:    ")))",
			expected: -3,
		},
		"8": {
			input:    ")())())",
			expected: -3,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := day1{}
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
			input:    ")",
			expected: 1,
		},
		"2": {
			input:    "()())",
			expected: 5,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := day1{}
			in := cmd.Input{SI: tt.input}

			res := d.Part1(in)

			assert.Equal(t, tt.expected, res)
		})
	}
}
