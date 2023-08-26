package utils

import "testing"

func TestFormatCmd(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  help map ",
			expected: []string{"help", "map"},
		},
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "Help MAP",
			expected: []string{"help", "map"},
		},
	}

	for _, testCase := range cases {
		actual := FormatCmd(testCase.input)

		if lenAct, lenExp := len(actual), len(testCase.expected); lenAct != lenExp {
			t.Errorf("lengths don't match: actual:%d expected:%d", lenAct, lenExp)
			continue
		}

		for index, word := range actual {
			if word != testCase.expected[index] {
				t.Errorf("FormatCmd(%s) == %v, expected %v", testCase.input, actual, testCase.expected)
				break
			}
		}
	}
}
