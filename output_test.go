package main

import (
	"reflect"
	"testing"
)

func TestAbbreviate(t *testing.T) {
	output := abbreviate("Nation Institute of Techology")
	if output != "NIT" {
		t.Errorf("Expected NIT but got %s\n", output)
	}
	output = abbreviate("Indian Institute of Management")
	if output != "IIM" {
		t.Errorf("Expected IIM but got %s\n", output)
	}
}

func TestPalindrome(t *testing.T) {
	input := "malayalam"
	output := isPalindrome(input)

	if !output {
		t.Errorf("Expected to be Palindrome")
	}

	input = "hello"
	output = isPalindrome(input)

	if output {
		t.Errorf("Expected not to be palindrome")
	}
}

func TestPangram(t *testing.T) {
	input := "The quick brown fox jumps over the lazy dog"
	output := panagram(input)

	if !output {
		t.Errorf("Expected to be Pangram")
	}

	input = "Hello World"
	output = panagram(input)

	if output {
		t.Errorf("Expected not to be Pangram, some error in prgrm")
	}
}

func TestFrequency(t *testing.T) {
	input := "bobby"
	expected := map[string]int{"b": 3, "o": 1, "y": 1}
	result := frequency(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v as frequency but got %v", expected, result)
	}
}

func TestFizzbizz(t *testing.T) {
	input := 15
	expected := []string{"1", "2", "fizz", "4", "bizz", "fizz", "7", "8", "fizz", "bizz", "11", "fizz", "13", "14", "fizzbizz"}
	result := fizzbizz(input)

	if reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v  but got %v", expected, result)
	}
}
