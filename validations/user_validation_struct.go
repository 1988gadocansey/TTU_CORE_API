package validations

import (
	"TTU_CORE_ERP_API/data/dto"
	"TTU_CORE_ERP_API/data/response"
	"github.com/go-playground/validator"
	"reflect"
	_ "reflect"
)

var UserStruct dto.User

func ValidateUserStruct() response.ErrorStruct {
	var return_errors response.ErrorStruct

	if !reflect.DeepEqual(dto.User{}, UserStruct) {
		validate := validator.New()
		err := validate.Struct(UserStruct)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				error_data := response.ErrorData{
					Field: err.StructField(),
					Tag:   err.Tag(),
				}
				return_errors = append(return_errors, error_data)
			}
		}
	}

	return return_errors
}
