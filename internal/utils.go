package internal

import "github.com/go-playground/validator/v10"

func GetMessageFromFieldError(mapJson *map[string]string, objError validator.FieldError) {
	tag := objError.Tag()
	field := objError.StructField()
	param := objError.Param()
	switch tag {
	case "startswith":
		(*mapJson)[field] = "O campo " + field + " deve iniciar com " + param
	case "required":
		(*mapJson)[field] = "O campo " + field + " é obrigatorio"
	case "min":
		(*mapJson)[field] = "O campo " + field + " deve ter no minimo " + param + " letras."
	case "max":
		(*mapJson)[field] = "O campo " + field + " deve ter no máximo " + param + " letras."
	}
}
