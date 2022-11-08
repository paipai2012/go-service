package validator

import (
	"sale-service/util"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValueIn(fl validator.FieldLevel) bool {

	param := fl.Param()
	if param == "" {
		return true
	}

	field := fl.Field()
	values := strings.Split(param, " ")

	// log.Println(field.Interface())
	// log.Println(field.Kind(), currentField.Kind())
	// log.Printf("currentKind %v nullable %v found %v", currentKind, nullable, found)

	return util.In(field.String(), values)
}
