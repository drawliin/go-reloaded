package main

import (
	"testing"

	helper "project1/helpers"
)

func TestProcessText(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "hex basic",
			in:   "1E (hex) files were added",
			out:  "30 files were added",
		},
		{
			name: "bin basic",
			in:   "It has been 10 (bin) years",
			out:  "It has been 2 years",
		},
		{
			name: "up single",
			in:   "go (up) now",
			out:  "GO now",
		},
		{
			name: "low multi",
			in:   "HELLO WORLD (low, 2)",
			out:  "hello world",
		},
		{
			name: "cap multi",
			in:   "this is amazing (cap, 3)",
			out:  "This Is Amazing",
		},
		{
			name: "punctuation spacing",
			in:   "hello ,world !",
			out:  "hello, world!",
		},
		{
			name: "punctuation group",
			in:   "wait ... here !!",
			out:  "wait... here!!",
		},
		{
			name: "quotes single word",
			in:   "I am ' awesome ' here",
			out:  "I am 'awesome' here",
		},
		{
			name: "quotes multiword",
			in:   "He said: ' I am good ' today",
			out:  "He said: 'I am good' today",
		},
		{
			name: "a to an vowel",
			in:   "a amazing story",
			out:  "an amazing story",
		},
		{
			name: "a to an with h",
			in:   "a hour",
			out:  "an hour",
		},
		{
			name: "mixed heavy",
			in:   "It was 1E (hex) files , added in 10 (bin) seconds (up, 2) !",
			out:  "It was 30 files, added IN 2 SECONDS!",
		},
		{
			name: "opened (",
			in:   "hey up) dfssdsd (up)",
			out:  "hey up) DFSSDSD",
		},
		{
			name: "tag in newline",
			in:   "hello\n(up)",
			out:  "hello",
		},
		{
			name: "find word to apply tag",
			in:   "hh kk mm, .. 55 (up)",
			out:  "hh kk MM, .. 55",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := helper.ParseString(tt.in)
			if got != tt.out {
				t.Errorf("\nInput:    %q\nExpected: %q\nGot:      %q", tt.in, tt.out, got)
			}
		})
	}
}
