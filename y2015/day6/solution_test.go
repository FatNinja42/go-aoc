package main

import (
	"testing"

	"aoc/cmd"

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
		x1, y1, x2, y2 int
		m              map[point]int
	}

	t.Run("getCoords", func(t *testing.T) {
		var tests = map[string]struct {
			input      string
			expectedP1 point
			expectedP2 point
		}{
			"turn on": {
				input:      "turn on 1,2 through 98,99",
				expectedP1: point{1, 2},
				expectedP2: point{98, 99},
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				p1, p2 := getCoords(tt.input, "turn on ")
				assert.Equal(t, tt.expectedP1, p1)
				assert.Equal(t, tt.expectedP2, p2)
			})
		}
	})

	t.Run("turn on", func(t *testing.T) {
		var tests = map[string]struct {
			input    in
			expected interface{}
		}{
			"1": {
				input: in{
					0, 0, 0, 5, map[point]int{},
				},
				expected: map[point]int{
					point{0, 0}: 1,
					point{0, 1}: 1,
					point{0, 2}: 1,
					point{0, 3}: 1,
					point{0, 4}: 1,
					point{0, 5}: 1,
				},
			},
			"2": {
				input: in{
					0, 0, 2, 2, map[point]int{},
				},
				expected: map[point]int{
					point{0, 0}: 1,
					point{0, 1}: 1,
					point{0, 2}: 1,
					point{1, 0}: 1,
					point{1, 1}: 1,
					point{1, 2}: 1,
					point{2, 0}: 1,
					point{2, 1}: 1,
					point{2, 2}: 1,
				},
			},
		}
		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				turnOn(tt.input.m, point{tt.input.x1, tt.input.y1}, point{tt.input.x2, tt.input.y2})

				assert.Equal(t, tt.expected, tt.input.m)
			})
		}

	})

}
