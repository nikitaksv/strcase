/*
 * Copyright (c) 2021 Nikita Krasnikov
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package strcase

import (
	"unicode"
)

const (
	SeparatorUnderscore = '_'
	SeparatorDash       = '-'
	SeparatorDot        = '.'
)

// DotCase ex. dot.case
func ToDotCase(str string) string {
	return string(ToDotCaseRunes([]rune(str)))
}

// DotCase ex. dot.case
func ToDotCaseRunes(runes []rune) []rune {
	return concatRuneWords(ParseRunes(runes), []rune{SeparatorDot})
}

// KebabCase ex. kebab-case
func ToKebabCase(str string) string {
	return string(ToKebabCaseRunes([]rune(str)))
}

// KebabCase ex. kebab-case
func ToKebabCaseRunes(runes []rune) []rune {
	return concatRuneWords(ParseRunes(runes), []rune{SeparatorDash})
}

// SnakeCase ex. snake_case
func ToSnakeCase(str string) string {
	return string(ToSnakeCaseRunes([]rune(str)))
}

// SnakeCase ex. snake_case
func ToSnakeCaseRunes(runes []rune) []rune {
	return concatRuneWords(ParseRunes(runes), []rune{SeparatorUnderscore})
}

// CamelCase ex. camelCase
func ToCamelCase(str string) string {
	return string(ToCamelCaseRunes([]rune(str)))
}

// CamelCase ex. camelCase
func ToCamelCaseRunes(runes []rune) []rune {
	return camelCase(runes, false)
}

// PascalCase ex. PascalCase
func ToPascalCase(str string) string {
	return string(ToPascalCaseRunes([]rune(str)))
}

// PascalCase ex. PascalCase
func ToPascalCaseRunes(runes []rune) []rune {
	return camelCase(runes, true)
}

// Splits the input line into words.
// Delimiters: "(unicode space)","_", "-",".","A-Z (Upper Letter second word)"
func ParseString(str string) []string {
	var words []string
	for _, rs := range ParseRunes([]rune(str)) {
		words = append(words, string(rs))
	}
	return words
}

// Splits the input line into words
func ParseRunes(rs []rune) [][]rune {
	var words [][]rune

	var word []rune
	for _, r := range rs {
		if isDelimiter(r) {
			if unicode.IsLetter(r) {
				if len(word) > 0 && !isAllUpper(word) {
					words = append(words, toLowerRunes(word))
					word = []rune{}
				}
				word = append(word, r)
			} else if len(word) > 0 {
				words = append(words, toLowerRunes(word))
				word = []rune{}
			}
		} else {
			word = append(word, r)
		}
	}
	if len(word) > 0 {
		words = append(words, toLowerRunes(word))
	}

	return words
}

func camelCase(runes []rune, upper bool) []rune {
	var camelCase []rune
	for i, rs := range ParseRunes(runes) {
		if i == 0 {
			if upper {
				rs[0] = unicode.ToUpper(rs[0])
			} else {
				rs[0] = unicode.ToLower(rs[0])
			}
		} else if i > 0 {
			rs[0] = unicode.ToUpper(rs[0])
		}

		camelCase = append(camelCase, rs...)
	}

	return camelCase
}

func isDelimiter(r rune) bool {
	return r == SeparatorDot ||
		r == SeparatorDash ||
		r == SeparatorUnderscore ||
		unicode.IsSpace(r) ||
		unicode.IsUpper(r)
}

// All runes isUppercase
func isAllUpper(rs []rune) bool {
	flag := false
	for _, r := range rs {
		if unicode.IsUpper(r) {
			flag = true
		} else {
			return false
		}
	}

	return flag
}

func toLowerRunes(runes []rune) []rune {
	for i, r := range runes {
		runes[i] = unicode.ToLower(r)
	}
	return runes
}

func concatRuneWords(runeWords [][]rune, sep []rune) []rune {
	var word []rune
	for _, rs := range runeWords {
		if len(word) > 0 {
			word = append(word, sep...)
		}
		word = append(word, rs...)
	}

	return word
}
