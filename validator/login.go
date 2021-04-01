package validator

// import (
// 	"log"
// 	"strings"

// 	"github.com/go-playground/validator/v10"
// )

// func Login(fl validator.FieldLevel) bool {

// 	// field := fl.Field()
// 	// param := fl.Param()
// 	// log.Printf("param %v field %v ", param, field)

// 	currentField, _, _, _ := fl.GetStructFieldOK2()
// 	// log.Printf("currentField %v currentKind %v nullable %v found %v", currentField, currentKind, nullable, found)

// 	loginType := currentField.FieldByName("LoginType").String()
// 	if strings.EqualFold(loginType, "password") {
// 		userName := currentField.FieldByName("UserName").String()
// 		password := currentField.FieldByName("Password").String()
// 		log.Printf("%v %v", userName, password)
// 		if userName == "" || password == "" {
// 			return false
// 		}
// 	} else {
// 		mobile := currentField.FieldByName("Mobile").String()
// 		smsCode := currentField.FieldByName("SmsCode").String()
// 		log.Printf("%v %v", mobile, smsCode)
// 		if mobile == "" || smsCode == "" {
// 			return false
// 		}
// 	}

// 	return true
// }
