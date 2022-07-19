package main

import (
	"aoc/cmd"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	var tests = map[string]struct {
		input    []string
		expected interface{}
	}{
		"1": {
			input: []string{
				"2x3x4",
			},
			expected: 58,
		},
		"2": {
			input: []string{
				"1x1x10",
			},
			expected: 43,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := day2{}
			in := cmd.Input{SLI: tt.input}

			res := d.Part1(in)

			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestPart2(t *testing.T) {
	var tests = map[string]struct {
		input    []string
		expected interface{}
	}{
		"1": {
			input: []string{
				"2x3x4",
			},
			expected: 34,
		},
		"2": {
			input: []string{
				"1x1x10",
			},
			expected: 14,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := day2{}
			in := cmd.Input{SLI: tt.input}

			res := d.Part2(in)

			assert.Equal(t, tt.expected, res)
		})
	}
}
