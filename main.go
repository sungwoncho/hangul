package main

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

const KOREAN_OFFSET rune = 0xAC00

// Char is a representation of Korean alphabet with its building parts
type Char struct {
	Lead  rune
	Vowel rune
	Tail  rune
}

// Deconstruct deconstructs a Korean alphabet into its lead, vowel, and tail characters
// and returns a Char containing containing those jvalues
func Deconstruct(c rune) Char {
	codeNum := c - KOREAN_OFFSET

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
