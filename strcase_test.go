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
	want := []string{"field", "name"}
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
			name: "space(lower)",
			args: args{str: "field name"},
			want: want,
		},
		{
			name: "space(upper)",
			args: args{str: "FIELD NAME"},
			want: want,
		},
		{
			name: "snake_case(lower)",
			args: args{str: "field_name"},
			want: want,
		},
		{
			name: "snake_case(upper)",
			args: args{str: "FIELD_NAME"},
			want: want,
		},
		{
			name: "kebab-case(lower)",
			args: args{str: "field-name"},
			want: want,
		},
		{
			name: "kebab-case(upper)",
			args: args{str: "FIELD-NAME"},
			want: want,
		},
		{
			name: "PascalCase",
			args: args{str: "FieldName"},
			want: want,
		},
		{
			name: "camelCase",
			args: args{str: "fieldName"},
			want: want,
		},
		{
			name: "dotCase",
			args: args{str: "field.Name"},
			want: want,
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
