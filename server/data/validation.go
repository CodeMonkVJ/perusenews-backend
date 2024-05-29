package data

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator"
)

type ValidationError struct {
    validator.FieldError
}

func (v ValidationError) Error() string {
    return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
        v.Namespace(),
        v.Field(),
        v.Tag(),
    )
}

type ValidationErrors []ValidationError

// Errors converts the slice into a string slice
func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

// Validation contains
type Validation struct {
	validate *validator.Validate
}

func NewValidation() *Validation {
    validate := validator.New()
    validate.RegisterValidation("script", validateScript)

    return &Validation{validate}
}

func (v *Validation) Validate(i interface{}) ValidationErrors {
    errs := v.validate.Struct(i)

    if errs == nil {
        return []ValidationError{}
    }
    
    valErrs := errs.(validator.ValidationErrors)

    if len(valErrs) == 0 {
        return []ValidationError{}
    }

    var returnErrs []ValidationError
    for _, err := range valErrs {
        ve := ValidationError{err.(validator.FieldError)}
        returnErrs = append(returnErrs, ve)
    }

    return returnErrs
}

func validateScript(fl validator.FieldLevel) bool {
    re := regexp.MustCompile(`^https:\/\/utfs\.io\/f\/[a-z0-9-.]+$`)
    matches := re.FindAllString(fl.Field().String(), -1)

    return len(matches) == 1 
}
