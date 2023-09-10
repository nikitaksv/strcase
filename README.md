# strcase

[![Godoc Reference](https://godoc.org/github.com/nikitaksv/strcase?status.svg)](http://godoc.org/github.com/nikitaksv/strcase)
[![Coverage Status](https://coveralls.io/repos/github/nikitaksv/strcase/badge.svg?branch=main)](https://coveralls.io/github/nikitaksv/strcase?branch=main)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fnikitaksv%2Fstrcase.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fnikitaksv%2Fstrcase?ref=badge_shield)

Package strcase converts strings or runes to various cases

## Install

```sh
go get -u github.com/nikitaksv/strcase
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/nikitaksv/strcase"
)

func main() {
	// Ex. snake_case
	snake := strcase.ToSnakeCase("oneWord two-word THREE_WORD FourWord")
	fmt.Println(snake) // out: one_word_two_word_three_word_four_word

	// Ex. kebab-case
	kebab := strcase.ToKebabCase("oneWord two-word THREE_WORD FourWord")
	fmt.Println(kebab) // out: one-word-two-word-three-word-four-word

	// Ex. camelCase
	camel := strcase.ToCamelCase("oneWord two-word THREE_WORD FourWord")
	fmt.Println(camel) // out: oneWordTwoWordThreeWordFourWord

	// Ex. PascalCase
	pascal := strcase.ToPascalCase("oneWord two-word THREE_WORD FourWord")
	fmt.Println(pascal) // out: OneWordTwoWordThreeWordFourWord

	// Ex. dot.case
	dot := strcase.ToDotCase("oneWord two-word THREE_WORD FourWord")
	fmt.Println(dot) // out: one.word.two.word.three.word.four.word

	// Ex. mergecase
	merge := strcase.ToDotCase("oneWord two-word THREE_WORD FourWord")
	fmt.Println(merge) // out: onewordtwowordthreewordfourword

	// Add GUID acronym
	strcase.AddAcronym("GUID", "Guid", "GuId")
	// Convert text to camelCase and replace Guid to acronym GUID
	camelAcronym := strcase.ToCamelCaseAcronym("my Order Guid")
	fmt.Println(camelAcronym) // out: myOrderGUID
}
```

## Func table

| Function                          | Output                     |
|-----------------------------------|----------------------------|
| `ToSnakeCase(string)`             | `field_name`               |
| `ToSnakeCaseAcronym(string)`      | `field_ID`                 |
| `ToSnakeCaseRunes(runes)`         | `field_name`               |
| `ToCamelCase(string)`             | `fieldName`                |
| `ToCamelCaseAcronym(string)`      | `fieldID`                  |
| `ToCamelCaseRunes(runes)`         | `fieldName`                |
| `ToKebabCase(string)`             | `field-name`               |
| `ToKebabCaseAcronym(string)`      | `field-name-ID`            |
| `ToKebabCaseRunes(runes)`         | `field-name`               |
| `ToPascalCase(string)`            | `FieldName`                |
| `ToPascalCaseAcronym(string)`     | `FieldNameID`              |
| `ToPascalCaseRunes(runes)`        | `FieldName`                |
| `ToDotCase(string)`               | `field.name`               |
| `ToDotCaseAcronym(string)`        | `field.name.ID`            |
| `ToDotCaseRunes(runes)`           | `field.name`               |
| `ToMergeCase(string)`             | `fieldname`                |
| `ToMergeCaseAcronym(string)`      | `fieldnameID`              |
| `ToMergeCaseRunes(runes)`         | `fieldname`                |
| `ParseString(string)`             | `[]string{"field","name"}` |
| `ParseRunes(runes)`               | `[][]rune{"field","name"}` |
| `AddAcronym(string)`              | void                       |
| `SetAcronym(map[string][]string)` | void                       |
| `ReplaceAcronym(string)`          | `ID`                       |

## License

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fnikitaksv%2Fstrcase.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fnikitaksv%2Fstrcase?ref=badge_large)