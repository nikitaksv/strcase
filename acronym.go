package strcase

import (
	"sort"
	"strings"
	"sync"
)

var acrMap = &sync.Map{}
var _baseAcronyms = []string{
	"ID",
	"IDs",
	"APP",
	"IP",
	"URL",
	"UUID",
	"UUIDs",
	"HTTP",
	"HTTPS",
	"ASCII",
	"NASA",
	"LOL",
	"JSON",
	"IDE",
	"ES",
	"GUI",
	"IIFE",
	"XML",
	"SEO",
	"UX",
	"JS",
	"API",
	"UTC",
	"EOF",
	"FIFO",
	"SDK",
	"SQL",
	"SOAP",
	"ORM",
	"OOP",
	"TDD",
	"BDD",
	"SAAS",
	"PAAS",
	"IOT",
	"WYSIWYG",
	"SMACSS",
	"SOLID",
	"YAGNI",
	"CRUD",
	"CDN",
	"MVC",
}

func init() {
	for _, acronym := range _baseAcronyms {
		loadAcronym(acronym)
	}
}

func loadAcronym(acr string, variants ...string) {
	lower := false
	for _, v := range variants {
		if strings.ToLower(v) == v {
			lower = true
			break
		}
	}
	if !lower {
		variants = append(variants, strings.ToLower(acr))
	}
	for _, variant := range variants {
		acrMap.Store(variant, []rune(acr))
	}
}

func AddAcronym(acr string, variants ...string) {
	loadAcronym(acr, variants...)
}

func SetAcronyms(acrs map[string][]string) {
	values := make([]string, 0, len(acrs))
	for acr, vars := range acrs {
		values = append(values, acr)
		loadAcronym(acr, vars...)
	}
	sort.Strings(values)
	acrMap.Range(func(key, value interface{}) bool {
		if !bins(values, string(value.([]rune))) {
			acrMap.Delete(key)
		}
		return true
	})
}

func bins(arr []string, target string) bool {
	index := sort.SearchStrings(arr, target)
	return index < len(arr) && arr[index] == target
}
