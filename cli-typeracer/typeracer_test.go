package main

import (
	"testing"
)

func TestTargetWords(t *testing.T) {
	targetWords := CreateTargetWords()

	if targetWords[0] != "Many" {
		t.Errorf("Expected %q got %q", "Many", targetWords[0])
	}
	if targetWords[8] != "composition" {
		t.Errorf("Expected %q got %q", "composition", targetWords[0])
	}
}

func TestIsEndOfString(t *testing.T) {
	if !isEndOfString("Many", 3) {
		t.Error("Should be end of string for Many at index 3")
	}
	if isEndOfString("Many", 2) {
		t.Error("2 should not be end of string for Many")
	}
	if isEndOfString("dancer", 6) {
		t.Error("6 Should be end of string for dancer")
	}
}
