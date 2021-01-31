// Package strcase converts strings or runes to various cases:
//   | Function                        | Output                   |
//   |---------------------------------|--------------------------|
//   | ToSnakeCase(s)                  | field_name               |
//   | ToSnakeCaseRunes(rs)            | field_name               |
//   | ToCamelCase(s)                  | fieldName                |
//   | ToCamelCaseRunes(rs)            | fieldName                |
//   | ToKebabCase(s)                  | field-name               |
//   | ToKebabCaseRunes(rs)            | field-name               |
//   | ToPascalCase(s)                 | FieldName                |
//   | ToPascalCaseRunes(rs)           | FieldName                |
//   | ToDotCase(s)                    | field.name               |
//   | ToDotCaseRunes(rs)              | field.name               |
//   | ParseString(s)                  | []string{"field","name"} |
//   | ParseRunes(rs)                  | [][]rune{"field","name"} |
package strcase
