package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name		string
		input 		string
		expected	[]string
	}{
		{
			name:		"empty string",
			input: 		"",
			expected: 	[]string{},
		},
		{
			name: 		"all small letters",
			input:		"hello world",
			expected:	[]string{"hello", "world"},
		},
		{
			name: 		"all capital letters",
			input:		"HELLO WORLD",
			expected:	[]string{"hello", "world"},
		},
		{
			name: 		"camel cs",
			input:		"hELLO wORLD",
			expected:	[]string{"hello", "world"},
		},
		{
			name: 		"title cs",
			input:		"Hello World",
			expected:	[]string{"hello", "world"},
		},
		{
			name: 		"symbols included",
			input:		"Hello World!",
			expected:	[]string{"hello", "world!"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf(
				"length of the output do not match: Actual '%v' vs Expected '%v'", len(actual), len(c.expected))
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf(
				"result do not match: Actual '%v' vs Expected '%v'", actual[i], c.expected[i])
			}
		}
	}

}

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
