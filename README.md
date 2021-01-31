# strcase

[![Godoc Reference](https://godoc.org/github.com/nikitaksv/strcase?status.svg)](http://godoc.org/github.com/nikitaksv/strcase)

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

}
```

## Func table

| Function                          | Output                     |
|-----------------------------------|----------------------------|
| `ToSnakeCase(s)`                  | `field_name`               |
| `ToSnakeCaseRunes(rs)`            | `field_name`               |
| `ToCamelCase(s)`                  | `fieldName`                |
| `ToCamelCaseRunes(rs)`            | `fieldName`                |
| `ToKebabCase(s)`                  | `field-name`               |
| `ToKebabCaseRunes(rs)`            | `field-name`               |
| `ToPascalCase(s)`                 | `FieldName`                |
| `ToPascalCaseRunes(rs)`           | `FieldName`                |
| `ToDotCase(s)`                    | `field.name`               |
| `ToDotCaseRunes(rs)`              | `field.name`               |
| `ParseString(s)`                  | `[]string{"field","name"}` |
| `ParseRunes(rs)`                  | `[][]rune{"field","name"}` |
