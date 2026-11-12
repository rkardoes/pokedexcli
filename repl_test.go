package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {

	tests := map[string]struct {
		input    string
		expected []string
	}{
		"padded":      {input: " hello world ", expected: []string{"hello", "world"}},
		"capitals":    {input: "HeLLo wOrld", expected: []string{"hello", "world"}},
		"padded&caps": {input: " Hello World   ", expected: []string{"hello", "world"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := cleanInput(tc.input)
			if len(got) != len(tc.expected) {
				t.Fatalf("expected len: %d, got len: %d", len(tc.expected), len(got))
			}
			if !reflect.DeepEqual(tc.expected, got) {
				t.Fatalf("expected: %v, got: %v", tc.expected, got)
			}
		})
	}
}
