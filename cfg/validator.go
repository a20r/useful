package cfg

import (
	"reflect"

	"github.com/a20r/falta"
)

// ErrCfgMissingRequiredFields is returned when a required field is not set.
var ErrCfgMissingRequiredFields = falta.Newf("cfg: missing required fields %v")

// ErrCfgSelfValidationFailed is returned when a self validation fails.
var ErrCfgSelfValidationFailed = falta.Newf("cfg: self validation failed")

// Validate checks if all required fields are set
func Validate(config any) error {
	structVal := reflect.ValueOf(config)
	structType := reflect.TypeOf(config)

	missingFields := make([]string, 0, structVal.NumField())

	for i := 0; i < structVal.NumField(); i++ {
		fieldVal := structVal.Field(i)
		fieldType := structType.Field(i)
		tag, ok := fieldType.Tag.Lookup("cfg")

		if !ok {
			continue
		}

		if tag == "required" && fieldVal.IsZero() {
			missingFields = append(missingFields, fieldType.Name)
		}
	}

	if len(missingFields) > 0 {
		return ErrCfgMissingRequiredFields.New(missingFields)
	}

	if validator, ok := config.(validator); ok {
		if err := validator.Validate(); err != nil {
			return ErrCfgSelfValidationFailed.New().Wrap(err)
		}
	}

	return nil
}

type validator interface {
	Validate() error
}
