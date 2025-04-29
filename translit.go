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

type Map map[rune]string

// ////////////////////////////////////////////////////////////////////////////////// //

type specProc func(p, c, n rune, mapping Map) (string, bool)

// ////////////////////////////////////////////////////////////////////////////////// //

var baseMap = Map{
	'А': "A", 'а': "a", 'Б': "B", 'б': "b", 'В': "V", 'в': "v", 'Г': "G", 'г': "g",
	'Д': "D", 'д': "d", 'Е': "E", 'е': "e", 'Ё': "E", 'ё': "e", 'Ж': "J", 'ж': "j",
	'З': "Z", 'з': "z", 'И': "I", 'и': "i", 'Й': "Iy", 'й': "iy", 'К': "K", 'к': "k",
	'Л': "L", 'л': "l", 'М': "M", 'м': "m", 'Н': "N", 'н': "n", 'О': "O", 'о': "o",
	'П': "P", 'п': "p", 'Р': "R", 'р': "r", 'С': "S", 'с': "s", 'Т': "T", 'т': "t",
	'У': "U", 'у': "u", 'Ф': "F", 'ф': "f", 'Х': "H", 'х': "h", 'Ц': "C",
	'ц': "c", 'Ч': "Ch", 'ч': "ch", 'Ш': "Sh", 'ш': "sh", 'Щ': "Shch", 'щ': "shch",
	'Ъ': "Ie", 'ъ': "ie", 'Ы': "Y", 'ы': "y", 'Ь': "", 'ь': "", 'Э': "E", 'э': "e",
	'Ю': "Y", 'ю': "y", 'Я': "Ya", 'я': "ya",
}

var (
	scientificMap = Map{
		'Е': "E", 'е': "e", 'Ё': "Ë", 'ё': "ë", 'Ж': "Ž", 'ж': "ž", 'Й': "J", 'й': "j",
		'Х': "Ch", 'х': "ch", 'Ц': "C", 'ц': "c", 'Ч': "Č", 'ч': "č", 'Ш': "Š", 'ш': "š",
		'Щ': "Šč", 'щ': "šč", 'Ъ': "″", 'ъ': "″", 'Ы': "Y", 'ы': "y", 'Ь': "′", 'ь': "′",
		'Э': "È", 'э': "è", 'Ю': "Ju", 'ю': "ju", 'Я': "Ja", 'я': "ja",
	}

	iso9AMap = Map{
		'Е': "E", 'е': "e", 'Ё': "Ë", 'ё': "ë", 'Ж': "Ž", 'ж': "ž", 'Й': "J", 'й': "j",
		'Х': "H", 'х': "h", 'Ц': "C", 'ц': "c", 'Ч': "Č", 'ч': "č", 'Ш': "Š", 'ш': "š",
		'Щ': "Ŝ", 'щ': "ŝ", 'Ъ': "″", 'ъ': "″", 'Ы': "Y", 'ы': "y", 'Ь': "′", 'ь': "′",
		'Э': "È", 'э': "è", 'Ю': "Û", 'ю': "û", 'Я': "Â", 'я': "â",
	}

	iso9BMap = Map{
		'Е': "E", 'е': "e", 'Ё': "Yo", 'ё': "yo", 'Ж': "Zh", 'ж': "zh", 'Й': "J", 'й': "j",
		'Х': "X", 'х': "x", 'Ч': "Ch", 'ч': "ch", 'Ш': "Sh", 'ш': "sh", 'Щ': "Shh",
		'щ': "shh", 'Ъ': "``", 'ъ': "``", 'Ы': "Y`", 'ы': "y`", 'Ь': "`", 'ь': "`",
		'Э': "E`", 'э': "e`", 'Ю': "Yu", 'ю': "yu", 'Я': "Ya", 'я': "ya",
	}

	bgnMap = Map{
		'Ж': "Zh", 'ж': "zh", 'И': "I", 'и': "i", 'Й': "Y", 'й': "y", 'Х': "Kh", 'х': "kh",
		'Ц': "Ts", 'ц': "ts", 'Ч': "Ch", 'ч': "ch", 'Ш': "Sh", 'ш': "sh", 'Щ': "Shch",
		'щ': "shch", 'Ъ': "″", 'ъ': "″", 'Ы': "Y", 'ы': "y", 'Ь': "′", 'ь': "′", 'Э': "E",
		'э': "e", 'Ю': "Yu", 'ю': "yu", 'Я': "Ya", 'я': "ya",
	}

	bsMap = Map{
		'Е': "E", 'е': "e", 'Ё': "Ë", 'ё': "ë", 'Ж': "Zh", 'ж': "zh", 'Й': "Ĭ", 'й': "ĭ",
		'Х': "Kh", 'х': "kh", 'Ц': "Ts", 'ц': "ts", 'Ч': "Ch", 'ч': "ch", 'Ш': "Sh",
		'ш': "sh", 'Щ': "Shch", 'щ': "shch", 'Ъ': "″", 'ъ': "″", 'Ы': "Ȳ", 'ы': "ȳ",
		'Ь': "′", 'ь': "′", 'Э': "É", 'э': "é", 'Ю': "Yu", 'ю': "yu", 'Я': "Ya", 'я': "ya",
	}

	alalcMap = Map{
		'Е': "E", 'е': "e", 'Ё': "Ë", 'ё': "ë", 'Ж': "Zh", 'ж': "zh", 'Й': "Ĭ", 'й': "ĭ",
		'Х': "Kh", 'х': "kh", 'Ц': "T͡s", 'ц': "t͡s", 'Ч': "Ch", 'ч': "ch", 'Ш': "Sh",
		'ш': "sh", 'Щ': "Shch", 'щ': "shch", 'Ъ': "″", 'ъ': "″", 'Ы': "Y", 'ы': "y",
		'Ь': "′", 'ь': "′", 'Э': "Ė", 'э': "ė", 'Ю': "I͡u", 'ю': "i͡u", 'Я': "I͡a", 'я': "i͡a",
	}

	icaoMap = Map{
		'Е': "E", 'е': "e", 'Ё': "E", 'ё': "e", 'Ж': "Zh", 'ж': "zh", 'Й': "I", 'й': "i",
		'Х': "Kh", 'х': "kh", 'Ц': "Ts", 'ц': "ts", 'Ч': "Ch", 'ч': "ch", 'Ш': "Sh",
		'ш': "sh", 'Щ': "Shch", 'щ': "shch", 'Ъ': "Ie", 'ъ': "ie", 'Ы': "Y", 'ы': "y",
		'Ь': "", 'ь': "", 'Э': "E", 'э': "e", 'Ю': "Iu", 'ю': "iu", 'Я': "Ia", 'я': "ia",
	}
)

var (
	iso9BExt1 = Map{'ц': "cz", 'Ц': "Cz"}
	iso9BExt2 = Map{'ц': "c", 'Ц': "C"}

	bgnExt1 = Map{'е': "e", 'Е': "E", 'ё': "ë", 'Ё': "Ë"}
	bgnExt2 = Map{'е': "ye", 'Е': "Ye", 'ё': "yë", 'Ё': "Yë"}
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Scientific encodes text with scientific mappings
func Scientific(text string) string {
	return encode(text, scientificMap, nil)
}

// ISO9A encodes text with ISO 9:1995/A ГОСТ 7.79-2000/A mappings
func ISO9A(text string) string {
	return encode(text, iso9AMap, nil)
}

// ISO9B encodes text with ISO 9:1995/B ГОСТ 7.79-2000/Б mappings
func ISO9B(text string) string {
	return encode(text, iso9BMap, iso9BSpec)
}

// BGN encodes text with BGN mappings
func BGN(text string) string {
	return encode(text, bgnMap, bgnSpec)
}

// PCGN encodes text with PCGN mappings
func PCGN(text string) string {
	return encode(text, bgnMap, bgnSpec)
}

// ALALC encodes text with ALA-LC mappings
func ALALC(text string) string {
	return encode(text, alalcMap, nil)
}

// BS encodes text with BS 2979:1958 mappings
func BS(text string) string {
	return encode(text, bsMap, nil)
}

// ICAO encodes text with ICAO mappings
func ICAO(text string) string {
	return encode(text, icaoMap, nil)
}

// Custom encodes text with custom mapping
func Custom(text string, mapping Map) string {
	return encode(text, mapping, nil)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func encode(text string, mapping Map, proc specProc) string {
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

		if !isCyrillicRune(r) {
			output.WriteRune(r)
			p = r
			continue
		}

		if proc != nil {
			n, _, _ = input.ReadRune()

			input.UnreadRune()

			rr, ok = proc(p, r, n, mapping)

			if ok {
				output.WriteString(rr)
				continue
			}
		}

		p = r

		if mapping != nil && mapping[r] != "" {
			output.WriteString(mapping[r])
			continue
		}

		if baseMap[r] != "" {
			output.WriteString(baseMap[r])
			continue
		}
	}

	return output.String()
}

func isCyrillicRune(r rune) bool {
	switch {
	case r >= 1040 && r <= 1103,
		r == 1105, r == 1025:
		return true
	}

	return false
}

func iso9BSpec(p, c, n rune, mapping Map) (string, bool) {
	switch c {
	case 'ц', 'Ц':
		// nop
	default:
		return "", false
	}

	rr := mapping[n]

	if rr == "" {
		rr = baseMap[n]
	}

	switch rr {
	case "e", "i", "y", "j", "E", "I", "Y", "J":
		return iso9BExt2[c], true
	}

	return iso9BExt1[c], true
}

func bgnSpec(p, c, n rune, mapping Map) (string, bool) {
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
		return bgnExt2[c], true
	}

	return bgnExt1[c], true
}
