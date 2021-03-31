package service

import (
	"log"
	"moose-go/api"
	"moose-go/model"
	"moose-go/util"
	"strings"
)

type AccountService struct{}

func (as AccountService) Regist(registParam *model.RegistParam) {

	password := registParam.Password
	repassword := registParam.RePassword

	if !strings.EqualFold(password, repassword) {
		panic(api.NewException(api.PasswordInconsistencyCode))
	}

	log.Print(util.CryptoSha1(registParam.Password))
	log.Print(registParam)
}
