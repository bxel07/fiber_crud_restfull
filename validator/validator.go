package validator

import (
	"errors"
	"reflect"
	"strings"

	"github.com/microcosm-cc/bluemonday"

)
var sanitizer *bluemonday.Policy

func init() {
	sanitizer = bluemonday.UGCPolicy()
}

type FieldRule struct {
	Min      float64
	Max      float64
	Required bool
	Sanitize bool
}

type Validator struct {
	Rules map[string]FieldRule
}

func NewValidator() *Validator {
	return &Validator{
		Rules: make(map[string]FieldRule),
	}
}

func (v *Validator) AddRule(field string, rule FieldRule) {
	v.Rules[field] = rule
}

func (v *Validator) Validate(data interface{}) error {
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return errors.New("input must be a struct")
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		fieldValue := val.Field(i)
		fieldName := field.Tag.Get("json")
		if fieldName == "" {
			fieldName = strings.ToLower(field.Name)
		}

		if rule, ok := v.Rules[fieldName]; ok {
			if err := validateAndSanitizeField(fieldName, fieldValue, rule); err != nil {
				return err
			}
		}
	}

	return nil
}

func validateAndSanitizeField(name string, value reflect.Value, rule FieldRule) error {
	switch value.Kind() {
	case reflect.String:
		str := value.String()
		if rule.Required && strings.TrimSpace(str) == "" {
			return errors.New(name + " is required")
		}
		if rule.Max > 0 && float64(len(str)) > rule.Max {
			return errors.New(name + " is too long")
		}
		if rule.Sanitize {
			sanitized := sanitizer.Sanitize(str)
			value.SetString(sanitized)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v := float64(value.Int())
		if rule.Min != 0 && v < rule.Min {
			return errors.New(name + " is too small")
		}
		if rule.Max != 0 && v > rule.Max {
			return errors.New(name + " is too large")
		}
	case reflect.Float32, reflect.Float64:
		v := value.Float()
		if rule.Min != 0 && v < rule.Min {
			return errors.New(name + " is too small")
		}
		if rule.Max != 0 && v > rule.Max {
			return errors.New(name + " is too large")
		}
	}
	return nil
}