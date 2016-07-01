package main

import (
	"gopkg.in/d4l3k/messagediff.v1"
	"testing"
)

var tests = []struct {
	version   string
	platforms []Platform
}{
	{"go1.0", Platforms_1_0},
	{"go1.1", Platforms_1_1},
	{"go1.3", Platforms_1_3},
	{"go1.4", Platforms_1_4},
	{"foo", Platforms_1_4},
}

func TestSupportedPlatforms(t *testing.T) {
	var ps []Platform
	for _, test := range tests {
		ps = SupportedPlatforms(test.version)
		diff, equal := messagediff.PrettyDiff(test.platforms, ps)
		if !equal {
			t.Errorf("Support platforms for version %s were not properly combined:\n%s", test.version, diff)
		}
	}
}
