// Package translit convert russian text to translit
package translit

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type specProc func(p, c, n rune, extMap map[string]string) (string, bool)

// ////////////////////////////////////////////////////////////////////////////////// //

var baseRuEn = map[string]string{
	"а": "a", "А": "A", "Б": "B", "б": "b", "В": "V", "в": "v", "Г": "G", "г": "g",
	"Д": "D", "д": "d", "З": "Z", "з": "z", "И": "I", "и": "i", "К": "K", "к": "k",
	"Л": "L", "л": "l", "М": "M", "м": "m", "Н": "N", "н": "n", "О": "O", "о": "o",
	"П": "P", "п": "p", "Р": "R", "р": "r", "С": "S", "с": "s", "Т": "T", "т": "t",
	"У": "U", "у": "u", "Ф": "F", "ф": "f",
}

var scientificRuEn = map[string]string{
	"Е": "E", "е": "e", "Ё": "Ë", "ё": "ë", "Ж": "Ž", "ж": "ž", "Й": "J", "й": "j",
	"Х": "Ch", "х": "ch", "Ц": "C", "ц": "c", "Ч": "Č", "ч": "č", "Ш": "Š", "ш": "š",
	"Щ": "Šč", "щ": "šč", "Ъ": "″", "ъ": "″", "Ы": "Y", "ы": "y", "Ь": "′", "ь": "′",
	"Э": "È", "э": "è", "Ю": "Ju", "ю": "ju", "Я": "Ja", "я": "ja",
}

var iso9ARuEn = map[string]string{
	"Е": "E", "е": "e", "Ё": "Ë", "ё": "ë", "Ж": "Ž", "ж": "ž", "Й": "J", "й": "j",
	"Х": "H", "х": "h", "Ц": "C", "ц": "c", "Ч": "Č", "ч": "č", "Ш": "Š", "ш": "š",
	"Щ": "Ŝ", "щ": "ŝ", "Ъ": "″", "ъ": "″", "Ы": "Y", "ы": "y", "Ь": "′", "ь": "′",
	"Э": "È", "э": "è", "Ю": "Û", "ю": "û", "Я": "Â", "я": "â",
}

var iso9BRuEn = map[string]string{
	"Е": "E", "е": "e", "Ё": "Yo", "ё": "yo", "Ж": "Zh", "ж": "zh", "Й": "J", "й": "j",
	"Х": "X", "х": "x", "Ч": "Ch", "ч": "ch", "Ш": "Sh", "ш": "sh", "Щ": "Shh",
	"щ": "shh", "Ъ": "``", "ъ": "``", "Ы": "Y`", "ы": "y`", "Ь": "`", "ь": "`",
	"Э": "E`", "э": "e`", "Ю": "Yu", "ю": "yu", "Я": "Ya", "я": "ya",
}

var bgnRuEn = map[string]string{
	"Ж": "Zh", "ж": "zh", "И": "I", "и": "i", "Й": "Y", "й": "y", "Х": "Kh", "х": "kh",
	"Ц": "Ts", "ц": "ts", "Ч": "Ch", "ч": "ch", "Ш": "Sh", "ш": "sh", "Щ": "Shch",
	"щ": "shch", "Ъ": "″", "ъ": "″", "Ы": "Y", "ы": "y", "Ь": "′", "ь": "′", "Э": "E",
	"э": "e", "Ю": "Yu", "ю": "yu", "Я": "Ya", "я": "ya",
}

var bsRuEn = map[string]string{
	"Е": "E", "е": "e", "Ё": "Ë", "ё": "ë", "Ж": "Zh", "ж": "zh", "Й": "Ĭ", "й": "ĭ",
	"Х": "Kh", "х": "kh", "Ц": "Ts", "ц": "ts", "Ч": "Ch", "ч": "ch", "Ш": "Sh",
	"ш": "sh", "Щ": "Shch", "щ": "shch", "Ъ": "″", "ъ": "″", "Ы": "Ȳ", "ы": "ȳ",
	"Ь": "′", "ь": "′", "Э": "É", "э": "é", "Ю": "Yu", "ю": "yu", "Я": "Ya", "я": "ya",
}

var alalcRuEn = map[string]string{
	"Е": "E", "е": "e", "Ё": "Ë", "ё": "ë", "Ж": "Zh", "ж": "zh", "Й": "Ĭ", "й": "ĭ",
	"Х": "Kh", "х": "kh", "Ц": "T͡s", "ц": "t͡s", "Ч": "Ch", "ч": "ch", "Ш": "Sh",
	"ш": "sh", "Щ": "Shch", "щ": "shch", "Ъ": "″", "ъ": "″", "Ы": "Y", "ы": "y",
	"Ь": "′", "ь": "′", "Э": "Ė", "э": "ė", "Ю": "I͡u", "ю": "i͡u", "Я": "I͡a", "я": "i͡a",
}

var icaoRuEn = map[string]string{
	"Е": "E", "е": "e", "Ё": "E", "ё": "e", "Ж": "Zh", "ж": "zh", "Й": "I", "й": "i",
	"Х": "Kh", "х": "kh", "Ц": "Ts", "ц": "ts", "Ч": "Ch", "ч": "ch", "Ш": "Sh",
	"ш": "sh", "Щ": "Shch", "щ": "shch", "Ъ": "Ie", "ъ": "ie", "Ы": "Y", "ы": "y",
	"Ь": "", "ь": "", "Э": "E", "э": "e", "Ю": "Iu", "ю": "iu", "Я": "Ia", "я": "ia",
}

var iso9BSpecial = map[string]string{
	"ц": "cz", "Ц": "Cz", "ц+": "с", "Ц+": "С",
}

var bgnSpecial = map[string]string{
	"е": "e", "Е": "E", "ё": "ë", "Ё": "Ë",
	"е+": "ye", "Е+": "Ye", "ё+": "yë", "Ё+": "Yë",
}

// ////////////////////////////////////////////////////////////////////////////////// //

// EncodeToScientific encodes text with scientific mappings
//
// Deprecated: Use Scientific method instead
func EncodeToScientific(text string) string {
	return Scientific(text)
}

// EncodeToISO9A encodes text with ISO 9:1995/A ГОСТ 7.79-2000/A mappings
//
// Deprecated: Use ISO9A method instead
func EncodeToISO9A(text string) string {
	return ISO9A(text)
}

// EncodeToISO9B encodes text with ISO 9:1995/B ГОСТ 7.79-2000/Б mappings
//
// Deprecated: Use ISO9B method instead
func EncodeToISO9B(text string) string {
	return ISO9B(text)
}

// EncodeToBGN encodes text with BGN mappings
//
// Deprecated: Use BGN method instead
func EncodeToBGN(text string) string {
	return BGN(text)
}

// EncodeToPCGN encodes text with PCGN mappings
//
// Deprecated: Use PCGN method instead
func EncodeToPCGN(text string) string {
	return PCGN(text)
}

// EncodeToALALC encodes text with ALA-LC mappings
//
// Deprecated: Use ALALC method instead
func EncodeToALALC(text string) string {
	return ALALC(text)
}

// EncodeToBS encodes text with BS 2979:1958 mappings
//
// Deprecated: Use BS method instead
func EncodeToBS(text string) string {
	return BS(text)
}

// EncodeToBS encodes text with ICAO mappings
//
// Deprecated: Use ICAO method instead
func EncodeToICAO(text string) string {
	return ICAO(text)
}

// Scientific encodes text with scientific mappings
func Scientific(text string) string {
	return encode(text, scientificRuEn, nil)
}

// ISO9A encodes text with ISO 9:1995/A ГОСТ 7.79-2000/A mappings
func ISO9A(text string) string {
	return encode(text, iso9ARuEn, nil)
}

// ISO9B encodes text with ISO 9:1995/B ГОСТ 7.79-2000/Б mappings
func ISO9B(text string) string {
	return encode(text, iso9BRuEn, iso9BSpec)
}

// BGN encodes text with BGN mappings
func BGN(text string) string {
	return encode(text, bgnRuEn, bgnSpec)
}

// PCGN encodes text with PCGN mappings
func PCGN(text string) string {
	return encode(text, bgnRuEn, bgnSpec)
}

// ALALC encodes text with ALA-LC mappings
func ALALC(text string) string {
	return encode(text, alalcRuEn, nil)
}

// BS encodes text with BS 2979:1958 mappings
func BS(text string) string {
	return encode(text, bsRuEn, nil)
}

// ICAO encodes text with ICAO mappings
func ICAO(text string) string {
	return encode(text, icaoRuEn, nil)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func encode(text string, extMap map[string]string, proc specProc) string {
	if text == "" {
		return ""
	}

	var input = bytes.NewBufferString(text)
	var output = bytes.NewBuffer(nil)

	// Previous, next letter for special processor
	var p, n rune
	var rr string
	var ok bool

	for {
		r, _, err := input.ReadRune()

		if err != nil {
			break
		}

		if !isRussianChar(r) {
			output.WriteRune(r)
			p = r
			continue
		}

		if proc != nil {
			n, _, _ = input.ReadRune()

			input.UnreadRune()

			rr, ok = proc(p, r, n, extMap)

			if ok {
				output.WriteString(rr)
				continue
			}
		}

		p = r

		rr, ok = baseRuEn[string(r)]

		if ok {
			output.WriteString(rr)
			continue
		}

		rr, ok = extMap[string(r)]

		if ok {
			output.WriteString(rr)
		}
	}

	return output.String()
}

func iso9BSpec(p, c, n rune, extMap map[string]string) (string, bool) {
	if c != 'ц' && c != 'Ц' {
		return "", false
	}

	rr, ok := baseRuEn[string(n)]

	if !ok {
		rr = extMap[string(n)]
	}

	switch rr {
	case "e", "i", "y", "j", "E", "I", "Y", "J":
		return iso9BSpecial[string(c)+"+"], true
	}

	return iso9BSpecial[string(c)], true
}

func bgnSpec(p, c, n rune, extMap map[string]string) (string, bool) {
	switch c {
	case 'е', 'Е', 'ё', 'Ё':
		// nop
	default:
		return "", false
	}

	switch p {
	case 0, ' ',
		'а', 'у', 'о', 'ы', 'и', 'э', 'я', 'ю',
		'А', 'У', 'О', 'Ы', 'И', 'Э', 'Я', 'Ю':
		return bgnSpecial[string(c)+"+"], true
	}

	return bgnSpecial[string(c)], true
}

func isRussianChar(r rune) bool {
	switch {
	case r >= 1040 && r <= 1103,
		r == 1105, r == 1025:
		return true
	}

	return false
}
