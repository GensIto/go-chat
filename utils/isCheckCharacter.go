package utils

import (
	"errors"
	"unicode"
)

func IsCheckCharacter(str string) error {
	for _, r := range str {
		hiragana := unicode.In(r, unicode.Hiragana)
		katakana := unicode.In(r, unicode.Katakana)
		kanji := unicode.In(r, unicode.Han)
		if hiragana || katakana || kanji {
			return errors.New("Sorry, I can't understand Japanese😥")
		}
		if r == ' ' {
			return errors.New("Sorry, Don't use spaces please😥")
		}
	}
	return nil
}
