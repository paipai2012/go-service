package validator

import (
	"sort"
	"strings"

	"github.com/go-playground/validator/v10"
)

func in(target string, arr []string) bool {
	sort.Strings(arr)
	index := sort.SearchStrings(arr, target)
	if index < len(arr) && arr[index] == target {
		return true
	}
	return false
}

func ValueIn(fl validator.FieldLevel) bool {

	param := fl.Param()
	if param == "" {
		return true
	}

	field := fl.Field()
	valeus := strings.Split(param, " ")

	// log.Println(field.Interface())
	// log.Println(field.Kind(), currentField.Kind())
	// log.Printf("currentKind %v nullable %v found %v", currentKind, nullable, found)

	return in(field.String(), valeus)
}
