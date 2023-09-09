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

// Package strcase converts strings or runes to various cases:
//
//	| Function                        | Output                   |
//	|---------------------------------|--------------------------|
//	| ToMergeCase(s)                  | fieldname                |
//	| ToMergeCaseRunes(s)             | fieldname                |
//	| ToSnakeCase(s)                  | field_name               |
//	| ToSnakeCaseRunes(rs)            | field_name               |
//	| ToCamelCase(s)                  | fieldName                |
//	| ToCamelCaseRunes(rs)            | fieldName                |
//	| ToKebabCase(s)                  | field-name               |
//	| ToKebabCaseRunes(rs)            | field-name               |
//	| ToPascalCase(s)                 | FieldName                |
//	| ToPascalCaseRunes(rs)           | FieldName                |
//	| ToDotCase(s)                    | field.name               |
//	| ToDotCaseRunes(rs)              | field.name               |
//	| ParseString(s)                  | []string{"field","name"} |
//	| ParseRunes(rs)                  | [][]rune{"field","name"} |
package strcase
