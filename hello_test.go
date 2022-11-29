package main

import (
	"strings"
	"testing"
)

func TestGetGreeting(t *testing.T) {
	result := getGreeting()
	expected := ", Entire Studio!"
	if !strings.HasSuffix(result, expected) {
		t.Errorf(
			"getGreeting() returned unexpected value. Got: `%s`. Should end with: `%s`.",
			result,
			expected,
		)
	}
}

func TestGetGreetingVerb(t *testing.T) {
	result := getGreetingVerb()
	expected := []string{
		"Hello",
		"Howdy",
		"Hi",
	}

	matched := false
	for i := 0; i < len(expected); i++ {
		if result == expected[i] {
			matched = true
		}
	}

	if !matched {
		t.Errorf(
			"getGreetingVerb() returned unexpected value. Got: `%s`.",
			result,
		)
	}
}
