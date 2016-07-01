package main

import (
	"reflect"
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
		if !reflect.DeepEqual(ps, test.platforms) {
			t.Fatalf("Supported platforms for version %s were not properly combined:\nExpected: %v\nActual: %v", test.version, test.platforms, ps)
		}
	}
}
