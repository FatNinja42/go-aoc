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
			input:    []string{"ugknbfddgicrmopn", "aaa"},
			expected: 2,
		},
		"2": {
			input:    []string{"jchzalrnumimnmhp", "haegwjzuvuyypxyu", "dvszwmarrgswjxmb"},
			expected: 0,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := day5{}
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
			input:    []string{"qjhvhtzxzqqjkmpb", "xxyxx"},
			expected: 2,
		},
		"2": {
			input:    []string{"uurcxstgmygtbstg", "ieodomkazucvgmuy"},
			expected: 0,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d := day5{}
			in := cmd.Input{SLI: tt.input}

			res := d.Part2(in)

			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestUtilities(t *testing.T) {
	t.Run("3 vowels", func(t *testing.T) {
		var tests = map[string]struct {
			input    string
			expected bool
		}{
			"1": {
				input:    "ugknbfddgicrmopn",
				expected: true,
			},
			"2": {
				input:    "aaa",
				expected: true,
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				res := ContainsThreeVowels(tt.input)

				assert.Equal(t, tt.expected, res)
			})
		}
	})

	t.Run("duplicated letter", func(t *testing.T) {
		var tests = map[string]struct {
			input    string
			expected bool
		}{
			"1": {
				input:    "ugknbfddgicrmopn",
				expected: true,
			},
			"2": {
				input:    "aaa",
				expected: true,
			},
			"3": {
				input:    "asdfg",
				expected: false,
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				res := HasDuplicateLetter(tt.input)

				assert.Equal(t, tt.expected, res)
			})
		}
	})

	t.Run("forbidden strings", func(t *testing.T) {
		var tests = map[string]struct {
			input    string
			expected bool
		}{
			"1": {
				input:    "ugknbfddgicrmopn",
				expected: true,
			},
			"2": {
				input:    "aaa",
				expected: true,
			},
			"3": {
				input:    "asdpq",
				expected: false,
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				res := DoesNotContainStrings(tt.input)

				assert.Equal(t, tt.expected, res)
			})
		}
	})

	t.Run("contains two pairs", func(t *testing.T) {
		var tests = map[string]struct {
			input    string
			expected bool
		}{
			"1": {
				input:    "xxyxx",
				expected: true,
			},
			"2": {
				input:    "qjhvhtzxzqqjkmpb",
				expected: true,
			},
			"3": {
				input:    "aaa",
				expected: false,
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				res := ContainsTwoPairs(tt.input)

				assert.Equal(t, tt.expected, res)
			})
		}
	})

	t.Run("has duplicated letter separated", func(t *testing.T) {
		var tests = map[string]struct {
			input    string
			expected bool
		}{
			"1": {
				input:    "xxyxx",
				expected: true,
			},
			"2": {
				input:    "qjhvhtzxzqqjkmpb",
				expected: true,
			},
			"3": {
				input:    "aaa",
				expected: true,
			},
			"4": {
				input:    "uurcxstgmygtbstg",
				expected: false,
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				res := HasDuplicateLetterSeparated(tt.input)

				assert.Equal(t, tt.expected, res)
			})
		}
	})
}
