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

// ToMergeCase MergeCase ex. mergecase
func ToMergeCase(str string) string {
	return string(ToMergeCaseRunes([]rune(str)))
}

// ToMergeCaseAcronym Replace acronym in string. Ex. mergecase
func ToMergeCaseAcronym(str string) string {
	return string(ToMergeCaseAcronymRunes([]rune(str)))
}

// ToMergeCaseRunes MergeCase ex. mergecaseID
func ToMergeCaseRunes(runes []rune) []rune {
	return concatRuneWords(ParseRunes(runes), nil)
}

// ToMergeCaseAcronymRunes Replace acronym in slice of runes. Ex. mergecaseID
func ToMergeCaseAcronymRunes(runes []rune) []rune {
	return concatRuneWords(ReplaceAcronymsRunes(ParseRunes(runes)), nil)
}

// ToDotCase DotCase ex. dot.case
func ToDotCase(str string) string {
	return string(ToDotCaseRunes([]rune(str)))
}

// ToDotCaseAcronym Replace acronym in string. Ex. dot.case.ID
func ToDotCaseAcronym(str string) string {
	return string(ToDotCaseAcronymRunes([]rune(str)))
}

// ToDotCaseRunes DotCase ex. dot.case
func ToDotCaseRunes(runes []rune) []rune {
	return concatRuneWords(ParseRunes(runes), []rune{SeparatorDot})
}

// ToDotCaseAcronymRunes Replace acronym in slice of runes. Ex. dot.case.ID
func ToDotCaseAcronymRunes(runes []rune) []rune {
	return concatRuneWords(ReplaceAcronymsRunes(ParseRunes(runes)), []rune{SeparatorDot})
}

// ToKebabCase KebabCase ex. kebab-case
func ToKebabCase(str string) string {
	return string(ToKebabCaseRunes([]rune(str)))
}

// ToKebabCaseAcronym Replace acronym in string. Ex. kebab-case-ID
func ToKebabCaseAcronym(str string) string {
	return string(ToKebabCaseAcronymRunes([]rune(str)))
}

// ToKebabCaseRunes KebabCase ex. kebab-case
func ToKebabCaseRunes(runes []rune) []rune {
	return concatRuneWords(ParseRunes(runes), []rune{SeparatorDash})
}

// ToKebabCaseAcronymRunes Replace acronym in slice of runes. Ex. kebab-case-ID
func ToKebabCaseAcronymRunes(runes []rune) []rune {
	return concatRuneWords(ReplaceAcronymsRunes(ParseRunes(runes)), []rune{SeparatorDash})
}

// ToSnakeCase SnakeCase ex. snake_case
func ToSnakeCase(str string) string {
	return string(ToSnakeCaseRunes([]rune(str)))
}

// ToSnakeCaseAcronym Replace acronym in string. Ex. snake_case_ID
func ToSnakeCaseAcronym(str string) string {
	return string(ToSnakeCaseAcronymRunes([]rune(str)))
}

// ToSnakeCaseRunes SnakeCase ex. snake_case
func ToSnakeCaseRunes(runes []rune) []rune {
	return concatRuneWords(ParseRunes(runes), []rune{SeparatorUnderscore})
}

// ToSnakeCaseAcronymRunes Replace acronym in slice of runes. Ex. snake_case_ID
func ToSnakeCaseAcronymRunes(runes []rune) []rune {
	return concatRuneWords(ReplaceAcronymsRunes(ParseRunes(runes)), []rune{SeparatorUnderscore})
}

// ToCamelCase CamelCase ex. camelCase
func ToCamelCase(str string) string {
	return string(ToCamelCaseRunes([]rune(str)))
}

// ToCamelCaseAcronym Replace acronym in string. Ex. camelCaseID
func ToCamelCaseAcronym(str string) string {
	return string(ToCamelCaseAcronymRunes([]rune(str)))
}

// ToCamelCaseRunes CamelCase ex. camelCase
func ToCamelCaseRunes(runes []rune) []rune {
	return camelCase(runes, false, false)
}

// ToCamelCaseAcronymRunes Replace acronym in slice of runes. Ex. camelCaseID
func ToCamelCaseAcronymRunes(runes []rune) []rune {
	return camelCase(runes, false, true)
}

// ToPascalCase PascalCase ex. PascalCase
func ToPascalCase(str string) string {
	return string(ToPascalCaseRunes([]rune(str)))
}

// ToPascalCaseAcronym Replace acronym in string. Ex. PascalCaseID
func ToPascalCaseAcronym(str string) string {
	return string(ToPascalCaseAcronymRunes([]rune(str)))
}

// ToPascalCaseRunes PascalCase ex. PascalCase
func ToPascalCaseRunes(runes []rune) []rune {
	return camelCase(runes, true, false)
}

// ToPascalCaseAcronymRunes Replace acronym in slice of runes. Ex. PascalCaseID
func ToPascalCaseAcronymRunes(runes []rune) []rune {
	return camelCase(runes, true, true)
}

// ParseString Splits the input line into words.
// Delimiters: "(unicode space)","_", "-",".","A-Z (Upper Letter second word)"
func ParseString(str string) []string {
	var words []string
	for _, rs := range ParseRunes([]rune(str)) {
		words = append(words, string(rs))
	}
	return words
}

// ParseRunes Splits the input line into words
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

func camelCase(runes []rune, upper bool, replaceAcronym bool) []rune {
	camelCase := make([]rune, 0, len(runes))
	for i, rs := range ParseRunes(runes) {
		var nRs []rune
		var foundReplace bool
		if replaceAcronym {
			nRs, foundReplace = ReplaceAcronymRunes(rs)
		}

		if foundReplace {
			rs = nRs
		} else if i == 0 {
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
	for _, rw := range runeWords {
		if len(word) > 0 {
			word = append(word, sep...)
		}
		word = append(word, rw...)
	}

	return word
}

func ReplaceAcronyms(words []string) []string {
	for i, w := range words {
		nW, _ := ReplaceAcronymRunes([]rune(w))
		words[i] = string(nW)
	}
	return words
}

func ReplaceAcronym(word string) (string, bool) {
	rW, found := ReplaceAcronymRunes([]rune(word))
	return string(rW), found
}

func ReplaceAcronymsRunes(runeWords [][]rune) [][]rune {
	for i, rw := range runeWords {
		runeWords[i], _ = ReplaceAcronymRunes(rw)
	}
	return runeWords
}

func ReplaceAcronymRunes(runeWord []rune) ([]rune, bool) {
	return replaceAcronymRunes(runeWord, false)
}

func replaceAcronymRunes(runeWord []rune, exit bool) ([]rune, bool) {
	if acr, ok := acrMap.Load(string(runeWord)); ok {
		if acr, ok := acr.([]rune); ok {
			return acr, true
		}
	} else if !exit {
		return replaceAcronymRunes(toLowerRunes(runeWord), true)
	}
	return runeWord, false
}

func isLower(runeWord []rune) bool {
	for _, r := range runeWord {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return true
}
