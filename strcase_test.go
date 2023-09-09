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
	"reflect"
	"testing"
)

func BenchmarkToCamelCase(b *testing.B) {
	b.Run("Foo Bar", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ToCamelCase("Foo Bar")
		}
	})
}

func BenchmarkToPascalCase(b *testing.B) {
	b.Run("Foo Bar", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ToPascalCase("Foo Bar")
		}
	})
}
func BenchmarkToDotCase(b *testing.B) {
	b.Run("Foo Bar", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ToDotCase("Foo Bar")
		}
	})
}
func BenchmarkToSnakeCase(b *testing.B) {
	b.Run("Foo Bar", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ToSnakeCase("Foo Bar")
		}
	})
}
func BenchmarkToKebabCase(b *testing.B) {
	b.Run("Foo Bar", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ToKebabCase("Foo Bar")
		}
	})
}

func TestToCamelCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "space(lower)",
			args: args{str: "field name"},
			want: "fieldName",
		},
		{
			name: "space(upper)",
			args: args{str: "FIELD NAME"},
			want: "fieldName",
		},
		{
			name: "snake_case(lower)",
			args: args{str: "field_name"},
			want: "fieldName",
		},
		{
			name: "snake_case(upper)",
			args: args{str: "FIELD_NAME"},
			want: "fieldName",
		},
		{
			name: "kebab-case(lower)",
			args: args{str: "field-name"},
			want: "fieldName",
		},
		{
			name: "kebab-case(upper)",
			args: args{str: "FIELD-NAME"},
			want: "fieldName",
		},
		{
			name: "PascalCase",
			args: args{str: "FieldName"},
			want: "fieldName",
		},
		{
			name: "camelCase",
			args: args{str: "fieldName"},
			want: "fieldName",
		},
		{
			name: "camelCase",
			args: args{str: "fieldId"},
			want: "fieldId",
		},
		{
			name: "dot",
			args: args{str: "field.Name"},
			want: "fieldName",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToCamelCase(tt.args.str); got != tt.want {
				t.Errorf("ToCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToPascalCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "space(lower)",
			args: args{str: "field name"},
			want: "FieldName",
		},
		{
			name: "space(upper)",
			args: args{str: "FIELD NAME"},
			want: "FieldName",
		},
		{
			name: "snake_case(lower)",
			args: args{str: "field_name"},
			want: "FieldName",
		},
		{
			name: "snake_case(upper)",
			args: args{str: "FIELD_NAME"},
			want: "FieldName",
		},
		{
			name: "kebab-case(lower)",
			args: args{str: "field-name"},
			want: "FieldName",
		},
		{
			name: "kebab-case(upper)",
			args: args{str: "FIELD-NAME"},
			want: "FieldName",
		},
		{
			name: "PascalCase",
			args: args{str: "FieldName"},
			want: "FieldName",
		},
		{
			name: "camelCase",
			args: args{str: "fieldName"},
			want: "FieldName",
		},
		{
			name: "dot",
			args: args{str: "field.Name"},
			want: "FieldName",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToPascalCase(tt.args.str); got != tt.want {
				t.Errorf("ToPascalCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToSnakeCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "space(lower)",
			args: args{str: "field name"},
			want: "field_name",
		},
		{
			name: "space(upper)",
			args: args{str: "FIELD NAME"},
			want: "field_name",
		},
		{
			name: "snake_case(lower)",
			args: args{str: "field_name"},
			want: "field_name",
		},
		{
			name: "snake_case(upper)",
			args: args{str: "FIELD_NAME"},
			want: "field_name",
		},
		{
			name: "kebab-case(lower)",
			args: args{str: "field-name"},
			want: "field_name",
		},
		{
			name: "kebab-case(upper)",
			args: args{str: "FIELD-NAME"},
			want: "field_name",
		},
		{
			name: "PascalCase",
			args: args{str: "FieldName"},
			want: "field_name",
		},
		{
			name: "camelCase",
			args: args{str: "fieldName"},
			want: "field_name",
		},
		{
			name: "dotCase",
			args: args{str: "field.Name"},
			want: "field_name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToSnakeCase(tt.args.str); got != tt.want {
				t.Errorf("ToSnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToKebabCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "space(lower)",
			args: args{str: "field name"},
			want: "field-name",
		},
		{
			name: "space(upper)",
			args: args{str: "FIELD NAME"},
			want: "field-name",
		},
		{
			name: "snake_case(lower)",
			args: args{str: "field_name"},
			want: "field-name",
		},
		{
			name: "snake_case(upper)",
			args: args{str: "FIELD_NAME"},
			want: "field-name",
		},
		{
			name: "kebab-case(lower)",
			args: args{str: "field-name"},
			want: "field-name",
		},
		{
			name: "kebab-case(upper)",
			args: args{str: "FIELD-NAME"},
			want: "field-name",
		},
		{
			name: "PascalCase",
			args: args{str: "FieldName"},
			want: "field-name",
		},
		{
			name: "camelCase",
			args: args{str: "fieldName"},
			want: "field-name",
		},
		{
			name: "dotCase",
			args: args{str: "field.Name"},
			want: "field-name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToKebabCase(tt.args.str); got != tt.want {
				t.Errorf("ToKebabCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseString(t *testing.T) {
	type args struct {
		str string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "multi style",
			args: args{str: "fieldName field_Name"},
			want: []string{"field", "name", "field", "name"},
		},
		{
			name: "digit",
			args: args{str: "field001"},
			want: []string{"field001"},
		},
		{
			name: "digit2",
			args: args{str: "field 001"},
			want: []string{"field", "001"},
		},
		{
			name: "space",
			args: args{str: " field name "},
			want: []string{"field", "name"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseString(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToDotCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "space(lower)",
			args: args{str: "field name"},
			want: "field.name",
		},
		{
			name: "space(upper)",
			args: args{str: "FIELD NAME"},
			want: "field.name",
		},
		{
			name: "snake_case(lower)",
			args: args{str: "field_name"},
			want: "field.name",
		},
		{
			name: "snake_case(upper)",
			args: args{str: "FIELD_NAME"},
			want: "field.name",
		},
		{
			name: "kebab-case(lower)",
			args: args{str: "field-name"},
			want: "field.name",
		},
		{
			name: "kebab-case(upper)",
			args: args{str: "FIELD-NAME"},
			want: "field.name",
		},
		{
			name: "PascalCase",
			args: args{str: "FieldName"},
			want: "field.name",
		},
		{
			name: "camelCase",
			args: args{str: "fieldName"},
			want: "field.name",
		},
		{
			name: "dotCase",
			args: args{str: "field.Name"},
			want: "field.name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToDotCase(tt.args.str); got != tt.want {
				t.Errorf("ToDotCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMergeCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "space(lower)",
			args: args{str: "field name"},
			want: "fieldname",
		},
		{
			name: "space(upper)",
			args: args{str: "FIELD NAME"},
			want: "fieldname",
		},
		{
			name: "snake_case(lower)",
			args: args{str: "field_name"},
			want: "fieldname",
		},
		{
			name: "snake_case(upper)",
			args: args{str: "FIELD_NAME"},
			want: "fieldname",
		},
		{
			name: "kebab-case(lower)",
			args: args{str: "field-name"},
			want: "fieldname",
		},
		{
			name: "kebab-case(upper)",
			args: args{str: "FIELD-NAME"},
			want: "fieldname",
		},
		{
			name: "PascalCase",
			args: args{str: "FieldName"},
			want: "fieldname",
		},
		{
			name: "camelCase",
			args: args{str: "fieldName"},
			want: "fieldname",
		},
		{
			name: "dotCase",
			args: args{str: "field.Name"},
			want: "fieldname",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToMergeCase(tt.args.str); got != tt.want {
				t.Errorf("ToMergeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestToMergeCaseAcronym(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		f    func(string) string
		want string
	}{
		{
			name: "mergecase",
			args: args{str: "field name id"},
			f:    ToMergeCaseAcronym,
			want: "fieldnameID",
		},
		{
			name: "dot.case",
			args: args{str: "field id"},
			f:    ToDotCaseAcronym,
			want: "field.ID",
		},
		{
			name: "kebab-case",
			args: args{str: "orderId"},
			f:    ToKebabCaseAcronym,
			want: "order-ID",
		},
		{
			name: "snake_case",
			args: args{str: "order_id"},
			f:    ToSnakeCaseAcronym,
			want: "order_ID",
		},
		{
			name: "camelCase",
			args: args{str: "order_id"},
			f:    ToCamelCaseAcronym,
			want: "orderID",
		},
		{
			name: "PascalCase",
			args: args{str: "order_id"},
			f:    ToPascalCaseAcronym,
			want: "OrderID",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f(tt.args.str); got != tt.want {
				t.Errorf("ToMergeCaseAcronym() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceAcronyms(t *testing.T) {
	words := []string{"order", "id"}
	want := []string{"order", "ID"}
	got := ReplaceAcronyms(words)
	if !reflect.DeepEqual(words, want) {
		t.Errorf("ReplaceAcronyms() = %v, want %v", got, want)
	}
}

func TestReplaceAcronym(t *testing.T) {
	word := "id"
	want := "ID"
	got, found := ReplaceAcronym(word)
	if got != want || !found {
		t.Errorf("ReplaceAcronym() = %v (%t), want %v (true)", got, found, want)
	}
}

func TestAddAcronym(t *testing.T) {
	AddAcronym("KKK", "kKk")
	found := false
	acrMap.Range(func(key, _ interface{}) bool {
		if key.(string) == "kKk" {
			found = true
			return false
		}
		return true
	})

	if !found {
		t.Errorf("not found added acronym in map")
	}
}

func TestIsLower(t *testing.T) {
	w := "qwerTy"
	if isLower([]rune(w)) {
		t.Errorf("word is not lower, but isLower() return true")
	}
}

// TestSetAcronyms set in global map. Call the func last
func TestSetAcronyms(t *testing.T) {
	acrs := map[string][]string{
		"MY":  {"my", "My"},
		"CAT": {"cat", "Cat"},
	}

	SetAcronyms(acrs)

	varsCnt := len(acrs) * 2

	total := 0
	acrMap.Range(func(key, _ interface{}) bool {
		total++
		return true
	})

	if varsCnt != total {
		t.Errorf("variant count not equal map variant count: %d != %d", varsCnt, total)
	}
}
