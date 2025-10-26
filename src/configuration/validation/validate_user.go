package validation

import (
	"encoding/json"
	"errors"

	"github.com/Swampertian/CRUD_GOLANG/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		uni := ut.New(en, en)
		transl, _ = uni.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserErro(
	validation_err error,
) *rest_err.RestErr {

	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Campo com tipo inválido")
	} else if errors.As(validation_err, &jsonValidationError) {

		errorsCauses := []rest_err.Causes{}
		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}
		return rest_err.NewBadRequestValidationError("Alguns campos estão incorretos", errorsCauses)
	} else {
		return rest_err.NewBadRequestError("Erro de validação ao converter campos")
	}
}
