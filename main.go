package main

import (
	"fmt"
)

// lead sounds ordered by their unicode code point
var leadSounds = []rune{
	'ㄱ', 'ㄲ', 'ㄴ', 'ㄷ', 'ㄸ',
	'ㄹ', 'ㅁ', 'ㅂ', 'ㅃ', 'ㅅ',
	'ㅆ', 'ㅇ', 'ㅈ', 'ㅉ', 'ㅊ',
	'ㅋ', 'ㅌ', 'ㅍ', 'ㅎ',
}

// vowel sounds ordered by their unicode code point
var vowelSounds = []rune{
	'ㅏ', 'ㅐ', 'ㅑ', 'ㅒ', 'ㅓ',
	'ㅔ', 'ㅕ', 'ㅖ', 'ㅗ', 'ㅘ',
	'ㅙ', 'ㅚ', 'ㅛ', 'ㅜ', 'ㅝ',
	'ㅞ', 'ㅟ', 'ㅠ', 'ㅡ', 'ㅢ',
	'ㅣ',
}

// tail sounds ordered by their unicode code point
var tailSounds = []rune{
	' ', 'ㄱ', 'ㄲ', 'ㄳ', 'ㄴ',
	'ㄵ', 'ㄶ', 'ㄷ', 'ㄹ', 'ㄺ',
	'ㄻ', 'ㄼ', 'ㄽ', 'ㄾ', 'ㄿ',
	'ㅀ', 'ㅁ', 'ㅂ', 'ㅄ', 'ㅅ',
	'ㅆ', 'ㅇ', 'ㅈ', 'ㅊ', 'ㅋ',
	'ㅌ', 'ㅍ', 'ㅎ',
}

const korean_offset rune = 0xAC00

// Char is a representation of Korean alphabet with its building parts
type Char struct {
	Lead  rune
	Vowel rune
	Tail  rune
}

// Deconstruct deconstructs a Korean alphabet into its lead, vowel, and tail characters
// and returns a Char containing containing those jvalues
func Deconstruct(c rune) Char {
	codeNum := c - korean_offset

	tailNum := codeNum % 28
	vowelNum := (codeNum / 28) % 21
	leadNum := codeNum / 28 / 21

	lead := leadSounds[leadNum]
	vowel := vowelSounds[vowelNum]
	var tail rune
	if tailNum > 0 {
		tail = tailSounds[tailNum]
	} else {
		tail = 0
	}

	return Char{
		Lead:  lead,
		Vowel: vowel,
		Tail:  tail,
	}
}

// ValidateJamo validates that a character is a Korean 'Jamo'
func ValidateJamo(c rune) bool {
	fmt.Printf("validating jamo: %0x ", c)
	return c >= 0x1100 && c <= 0x11FF
}

// ValidateSyllable validates that a character is a Korean syllable
func ValidateSyllable(c rune) bool {
	return c >= 0xAC00 && c <= 0xD7A3
}
