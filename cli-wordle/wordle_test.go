package main

import (
	"testing"
)

func TestNotCorrectSquare(t *testing.T) {
	sq1 := GetNotCorrectSquare('a', "apple")
	sq2 := GetNotCorrectSquare('f', "apple")
	if sq1 != HAS {
		t.Errorf("Expected %q got %q", HAS, sq1)
	}
	if sq2 != WRONG {
		t.Errorf("Expected %q got %q", WRONG, sq2)
	}
}

func TestWordMatch(t *testing.T) {
	t1 := CreateSquareOutput("apple", "apple")
	t2 := CreateSquareOutput("apple", "timed")
	t3 := CreateSquareOutput("apple", "house")

	p1 := CreateSquareOutput("steak", "house")
	o1 := WRONG + WRONG + WRONG + HAS + HAS
	p2 := CreateSquareOutput("steak", "strip")
	o2 := CORRECT + CORRECT + WRONG + WRONG + WRONG
	p3 := CreateSquareOutput("steak", "steal")
	o3 := CORRECT + CORRECT + CORRECT + CORRECT + WRONG
	p4 := CreateSquareOutput("steak", "steam")
	o4 := CORRECT + CORRECT + CORRECT + CORRECT + WRONG
	p5 := CreateSquareOutput("steak", "stead")
	o5 := CORRECT + CORRECT + CORRECT + CORRECT + WRONG
	p6 := CreateSquareOutput("steak", "steak")
	o6 := CORRECT + CORRECT + CORRECT + CORRECT + CORRECT

	if t1 != CORRECT+CORRECT+CORRECT+CORRECT+CORRECT {
		t.Errorf("Expected %q got %q", CORRECT+CORRECT+CORRECT+CORRECT+CORRECT, t1)
	}
	if t2 != WRONG+WRONG+WRONG+HAS+WRONG {
		t.Errorf("Expected %q got %q", WRONG+WRONG+WRONG+HAS+WRONG, t2)
	}

	if t3 != WRONG+WRONG+WRONG+WRONG+CORRECT {
		t.Errorf("Expected %q got %q", WRONG+WRONG+WRONG+WRONG+CORRECT, t3)
	}

	if p1 != o1 {
		t.Errorf("Expected %q got %q", o1, p1)
	}
	if p2 != o2 {
		t.Errorf("Expected %q got %q", o2, p2)
	}
	if p3 != o3 {
		t.Errorf("Expected %q got %q", o3, p3)
	}
	if p4 != o4 {
		t.Errorf("Expected %q got %q", o4, p4)
	}
	if p5 != o5 {
		t.Errorf("Expected %q got %q", o5, p5)
	}
	if p6 != o6 {
		t.Errorf("Expected %q got %q", o6, p6)
	}

	steakCase := CreateSquareOutput("steak", "steaa")
	steakOutput := CORRECT + CORRECT + CORRECT + CORRECT + WRONG
	if steakCase != steakOutput {
		t.Errorf("Expected %q got %q", steakOutput, steakCase)
	}
	pshawCase := CreateSquareOutput("phsaw", "ahead")
	pshawOutput := WRONG + CORRECT + WRONG + CORRECT + WRONG
	if pshawCase != pshawOutput {
		t.Errorf("Expected %q got %q", pshawOutput, pshawCase)
	}

}

func TestDictionaryBuilding(t *testing.T) {
	d := CreateDictionary()
	if !d["write"] {
		t.Error("Expected write to be in words")
	}
	if d["ffff"] {
		t.Error("Expected fffff to be not be in words")
	}
}

func TestReplaceRuneIndex(t *testing.T) {
	t1 := ReplaceAtIndex("nice", 'd', 0)
	if t1 != "dice" {
		t.Errorf("Expected dice instead got %q", t1)
	}
	t2 := ReplaceAtIndex("wowee", 'l', 2)
	t3 := ReplaceAtIndex(t2, 'a', 1)
	if t2 != "wolee" {
		t.Errorf("Expected wolee instead got %q", t2)
	}
	if t3 != "walee" {
		t.Errorf("Expected walee instead got %q", t3)
	}
}
