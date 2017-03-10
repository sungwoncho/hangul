package main

import (
	"testing"
)

func TestDeconstruct(t *testing.T) {
	char := Deconstruct('감')

	if char.Lead != 'ㄱ' {
		t.Errorf("Expected lead sound: ㄱ, got %c", char.Lead)
	}
	if char.Vowel != 'ㅏ' {
		t.Errorf("Expected vowel sound: ㅏ, got %c", char.Vowel)
	}
	if char.Tail != 'ㅁ' {
		t.Errorf("Expected tail sound: ㅁ, got %c", char.Tail)
	}
}

func TestDeconstruct_WithoutTail(t *testing.T) {
	char := Deconstruct('가')

	if char.Lead != 'ㄱ' {
		t.Errorf("Expected lead sound: ㄱ, got %c", char.Lead)
	}
	if char.Vowel != 'ㅏ' {
		t.Errorf("Expected vowel sound: ㅏ, got %c", char.Vowel)
	}
	if char.Tail != 0 {
		t.Errorf("Expected tail sound: 0, got %c", char.Tail)
	}
}

func TestValidateJamo(t *testing.T) {
	if !ValidateJamo('ㄱ') {
		t.Error("Expected 'ㄱ' to be a valid Jamo")
	}
	if ValidateJamo('a') {
		t.Error("Expected 'a' not to be a valid Jamo")
	}
}
