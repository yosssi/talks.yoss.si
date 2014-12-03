// +build ignore
package main

import "testing"

// START OMIT

var testCases = []struct {
	in  string
	out string
}{
	{"a", "A"},
	{"b", "B"},
	{"c", "C"},
	{"d", "D"},
	{"e", "E"},
}

func TestToUpper(t *testing.T) {
	for _, tc := range testCases {
		if result := ToUpper(tc.in); result != tc.out {
			t.Errorf("ToUpper(%q) => %q, want %q", tc.in, result, tc.out)
		}
	}
}

// END OMIT
