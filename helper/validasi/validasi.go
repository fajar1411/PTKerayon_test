package validasi

import (
	"errors"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func ValidationErrorHandle(data any) error {
	var msg string
	validate = validator.New()
	err := validate.Struct(data)

	if err != nil {
		log.Println(err)
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
		}
		if strings.Contains(err.Error(), "required") {
			msg = "field  wajib diisi"
		} else if strings.Contains(err.Error(), "min") {
			msg = "minimal karakter 5"
		}
		return errors.New(msg)
	}
	return nil
}

// 	castedObject, ok := err.(validator.ValidationErrors)
// 	if ok {
// 		for _, v := range castedObject {
// 			switch v.Tag() {
// 			case "required":
// 				message = fmt.Sprintf("%s Ada Tabel yang Kosong, Harap Diisi !!!", v.Field())
// 			case "min":
// 				message = fmt.Sprintf("%s Input Minimal %s character !!!", v.Field(), v.Param())
// 			case "max":
// 				message = fmt.Sprintf("%s Input Maksimal %s character !!!", v.Field(), v.Param())
// 			case "lte":
// 				message = fmt.Sprintf("%s Input tidak boleh di bawah %s !!!", v.Field(), v.Param())
// 			case "gte":
// 				message = fmt.Sprintf("%s Input tidak boleh di atas %s !!!", v.Field(), v.Param())
// 			case "numeric":
// 				message = fmt.Sprintf("%s Input Harus Numeric !!!", v.Field())
// 			case "url":
// 				message = fmt.Sprintf("%s Input Harus Url !!!", v.Field())
// 			case "email":
// 				message = fmt.Sprintf("%s Input Harus Email !!!", v.Field())
// 			case "password":
// 				message = fmt.Sprintf("%s Input value must be filled", v.Field())

// 			}
// 		}
// 	}

// return message
