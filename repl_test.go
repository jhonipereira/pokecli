package main

import "testing"

func TestFilterWord(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello there",
			expected: []string{
				"hello", "there",
			},
		},
		{
			input: "HIYA\t PARTNER\n",
			expected: []string{
				"hiya", "partner",
			},
		},
	}

	for _, cs := range cases {
		actual := FilterWord(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("Lenghts not equal: %v VS %v", len(actual), len(cs.expected))
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Errorf("%v does not equal %v", actualWord, expectedWord)
			}
		}
	}
}
