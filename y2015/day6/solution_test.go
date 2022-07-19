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
			d := day6{}
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
			d := day6{}
			in := cmd.Input{SI: tt.input}

			res := d.Part2(in)

			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestUtilities(t *testing.T) {
	type in struct {
		x1, x2, y1, y2 int
		m              map[point]bool
	}

	t.Run("turn on", func(t *testing.T) {
		var tests = map[string]struct {
			input    in
			expected interface{}
		}{
			"1": {
				input: in{
					0, 0, 0, 5, map[point]bool{},
				},
				expected: map[point]bool{
					point{0, 0}: true,
					point{0, 1}: true,
					point{0, 2}: true,
					point{0, 3}: true,
					point{0, 4}: true,
					point{0, 5}: true,
				},
			},
			"2": {
				input: in{
					0, 0, 2, 2, map[point]bool{},
				},
				expected: map[point]bool{
					point{0, 0}: true,
					point{0, 1}: true,
					point{0, 2}: true,
					point{1, 0}: true,
					point{1, 1}: true,
					point{1, 2}: true,
					point{2, 0}: true,
					point{2, 1}: true,
					point{2, 2}: true,
				},
			},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				TurnOn(point{tt.input.x1, tt.input.y1}, point{tt.input.x2, tt.input.y2}, tt.input.m)

				assert.Equal(t, tt.expected, tt.input.m)
			})
		}
	})

}
