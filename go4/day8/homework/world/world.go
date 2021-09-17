package world

import "unicode"

// IsPalindrome 判断是否回文
func IsPalindrome(s string) bool {
	var letters []rune
	for _, l := range s {
		// 判断l是不是已和letter
		if unicode.IsLetter(l) {
			letters = append(letters, unicode.ToLower(l))
		}
	}
	for i := 0; i < len(letters); i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true

}
