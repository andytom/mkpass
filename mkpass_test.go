package main

import (
	"math/rand"
	"testing"
)

func init() {
	// Fix the random seed for testing
	rand.Seed(1)
}

// Test readDict filters length correctly
func TestReadDictLength(t *testing.T) {
	path := "test_fixtures/length"
	minLen := 3

	ret, _ := readDict(path, minLen)

	if len(ret) == 0 {
		t.Fatalf("readDict retunred an empty array '%s'", ret)
	}

	for _, str := range ret {
		if len(str) < minLen {
			t.Fatalf("readDict Returned %s which contains '%s' that is less than %s", ret, str, minLen)
		}
	}
}

// Test readDict filters words that contain non alphanumeric characters
func TestReadDictFilter(t *testing.T) {
	path := "test_fixtures/filter"
	minLen := 3

	ret, _ := readDict(path, minLen)

	if len(ret) != 0 {
		t.Fatalf("readDict Returned %s which contains strings that fails the filter", ret)
	}
}

// Test genPassword works
func TestGenPassword(t *testing.T) {
	testDict := []string{"aa", "bb", "cc"}
	numWords := 3

	ret, _ := genPassword(testDict, numWords)

	if len(ret) != numWords {
		t.Fatalf("genPassword Returned %s which is too few words", ret)
	}
}

// Test genPassword error handling
func TestGenPasswordErrorHandlingBadDict(t *testing.T) {
	emptyDict := []string{}
	numWords := 3

	_, err := genPassword(emptyDict, numWords)

	if err == nil {
		t.Fatalf("genPassword didn't return an error with an empty dict")
	}
}

func TestGenPasswordErrorHandlingZeroLen(t *testing.T) {
	testDict := []string{"aa", "bb", "cc"}
	numWords := 0

	_, err := genPassword(testDict, numWords)

	if err == nil {
		t.Fatalf("genPassword didn't return an error with 0 words")
	}
}

func TestGenPasswordErrorHandlingNegativeLen(t *testing.T) {
	testDict := []string{"aa", "bb", "cc"}
	numWords := -1

	_, err := genPassword(testDict, numWords)

	if err == nil {
		t.Fatalf("genPassword didn't return an error negative words")
	}
}
