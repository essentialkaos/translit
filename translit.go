// Package translit convert russian text to translit
package translit

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var universal = map[rune]string{
	'а': "a", 'б': "b", 'в': "v", 'г': "g", 'д': "d", 'е': "e", 'ё': "jo", 'ж': "zh",
	'з': "z", 'и': "i", 'й': "j", 'к': "k", 'л': "l", 'м': "m", 'н': "n", 'о': "o",
	'п': "p", 'р': "r", 'с': "s", 'т': "t", 'у': "u", 'ф': "f", 'х': "h", 'ц': "c",
	'ч': "ch", 'ш': "sh", 'щ': "shh", 'ъ': "\"", 'ы': "y", 'ь': "'", 'э': "je",
	'ю': "ju", 'я': "ja",
}

var gost779b = map[rune]string{
	'а': "a", 'б': "b", 'в': "v", 'г': "g", 'д': "d", 'е': "e", 'ё': "yo", 'ж': "zh",
	'з': "z", 'и': "i", 'й': "j", 'к': "k", 'л': "l", 'м': "m", 'н': "n", 'о': "o",
	'п': "p", 'р': "r", 'с': "s", 'т': "t", 'у': "u", 'ф': "f", 'х': "x", 'ц': "c",
	'ч': "ch", 'ш': "sh", 'щ': "shh", 'ъ': "\"", 'ы': "y", 'ь': "'", 'э': "eh",
	'ю': "yu", 'я': "ya",
}

var iso9 = map[rune]string{
	'а': "a", 'б': "b", 'в': "v", 'г': "g", 'д': "d", 'е': "e", 'ё': "ë", 'ж': "ž",
	'з': "z", 'и': "i", 'й': "j", 'к': "k", 'л': "l", 'м': "m", 'н': "n", 'о': "o",
	'п': "p", 'р': "r", 'с': "s", 'т': "t", 'у': "u", 'ф': "f", 'х': "h", 'ц': "c",
	'ч': "č", 'ш': "š", 'щ': "ŝ", 'ъ': "\"", 'ы': "y", 'ь': "'", 'э': "è",
	'ю': "û", 'я': "â",
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Universal convert text with default mappings
func Universal(text string) string {
	return convert(text, universal)
}

// GOST779B convert text with GOST 7.79B mappings
func GOST779B(text string) string {
	return convert(text, gost779b)
}

// ISO9 convert text with ISO 9 mappings
func ISO9(text string) string {
	return convert(text, iso9)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func convert(text string, mapping map[rune]string) string {
	if text == "" {
		return ""
	}

	var buffer = bytes.NewBuffer(nil)

	for _, r := range text {
		if !isRussianChar(r) {
			buffer.WriteRune(r)
			continue
		}

		if isRussianCapsChar(r) {
			buffer.WriteString(strings.Title(mapping[getLowerRune(r)]))
			continue
		}

		buffer.WriteString(mapping[r])
	}

	return buffer.String()
}

func isRussianChar(r rune) bool {
	switch {
	case r >= 1040 && r <= 1103,
		r == 1105, r == 1025:
		return true
	}

	return false
}

func isRussianCapsChar(r rune) bool {
	switch {
	case r >= 1040 && r <= 1071, r == 1025:
		return true
	}

	return false
}

func getLowerRune(r rune) rune {
	if r == 1025 {
		return 1105
	}

	return r + 32
}
