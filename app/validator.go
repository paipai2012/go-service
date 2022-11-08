package app

import (
	mv "sale-service/validator"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func BindValidator() {
	// 绑定验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("valuein", mv.ValueIn)
	}
}
