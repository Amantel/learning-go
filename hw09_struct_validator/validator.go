package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrLen    = errors.New("length mismatch")
	ErrMin    = errors.New("value below minimum")
	ErrMax    = errors.New("value above maximum")
	ErrIn     = errors.New("value not in allowed set")
	ErrRegexp = errors.New("value does not match regexp")
)

var (
	ErrNotStruct            = errors.New("input is not a struct")
	ErrInvalidRule          = errors.New("invalid validation rule")
	ErrInvalidRegexp        = errors.New("invalid regexp in tag")
	ErrUnsupportedFieldKind = errors.New("unsupported field kind")
)

type ValidationError struct {
	Field    string
	Err      error
	Expected any
	Got      any
}

func (e ValidationError) Error() string {
	switch {
	case errors.Is(e.Err, ErrLen):
		return fmt.Sprintf("%s: length must be %v (got %v)", e.Field, e.Expected, e.Got)
	case errors.Is(e.Err, ErrMin):
		return fmt.Sprintf("%s: must be >= %v (got %v)", e.Field, e.Expected, e.Got)
	case errors.Is(e.Err, ErrMax):
		return fmt.Sprintf("%s: must be <= %v (got %v)", e.Field, e.Expected, e.Got)
	case errors.Is(e.Err, ErrIn):
		return fmt.Sprintf("%s: must be in %v (got %v)", e.Field, e.Expected, e.Got)
	case errors.Is(e.Err, ErrRegexp):
		return fmt.Sprintf("%s: does not match %v (got %v)", e.Field, e.Expected, e.Got)
	default:
		return fmt.Sprintf("%s: %v", e.Field, e.Err)
	}
}

func (e ValidationError) Unwrap() error {
	return e.Err
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var sb strings.Builder

	for _, err := range v {
		sb.WriteString(err.Error())
		sb.WriteString("\n")
	}

	return sb.String()
}

func Validate(input interface{}) error {
	v := reflect.ValueOf(input)
	t := reflect.TypeOf(input)

	if t == nil {
		return ErrNotStruct
	}

	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return ErrNotStruct
	}

	var errs ValidationErrors

	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := t.Field(i)

		tag := fieldType.Tag.Get("validate")
		if tag == "" {
			continue
		}

		if !fieldType.IsExported() {
			continue
		}

		rules := strings.Split(tag, "|")

		switch fieldType.Type.Kind() {
		case reflect.String, reflect.Int:
			vErrs, pErr := applyRules(fieldType.Name, fieldValue, rules)
			if pErr != nil {
				return fmt.Errorf("field %q: %w", fieldType.Name, pErr)
			}
			errs = append(errs, vErrs...)

		case reflect.Slice:
			for j := 0; j < fieldValue.Len(); j++ {
				itemName := fmt.Sprintf("%s[%d]", fieldType.Name, j)
				vErrs, pErr := applyRules(itemName, fieldValue.Index(j), rules)
				if pErr != nil {
					return fmt.Errorf("field %q: %w", itemName, pErr)
				}
				errs = append(errs, vErrs...)
			}

		default:
			return fmt.Errorf("field %q: %w", fieldType.Name, ErrUnsupportedFieldKind)
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func applyRules(field string, value reflect.Value, rules []string) (ValidationErrors, error) {
	var errs ValidationErrors

	for _, rule := range rules {
		var (
			ve  *ValidationError
			err error
		)

		switch value.Kind() {
		case reflect.String:
			ve, err = validateString(value.String(), rule)
		case reflect.Int:
			ve, err = validateInt(int(value.Int()), rule)
		default:
			return nil, ErrUnsupportedFieldKind
		}

		if err != nil {
			return nil, err
		}

		if ve != nil {
			ve.Field = field
			errs = append(errs, *ve)
		}
	}

	return errs, nil
}

func validateString(str string, rule string) (*ValidationError, error) {
	name, arg, ok := strings.Cut(rule, ":")
	if !ok {
		return nil, fmt.Errorf("%w: %q", ErrInvalidRule, rule)
	}

	switch name {
	case "len":
		n, err := strconv.Atoi(arg)
		if err != nil {
			return nil, fmt.Errorf("%w: len arg %q not int", ErrInvalidRule, arg)
		}

		if len(str) != n {
			return &ValidationError{Err: ErrLen, Expected: n, Got: len(str)}, nil
		}

	case "regexp":
		re, err := regexp.Compile(arg)
		if err != nil {
			return nil, fmt.Errorf("%w: %q: %v", ErrInvalidRegexp, arg, err)
		}

		if !re.MatchString(str) {
			return &ValidationError{Err: ErrRegexp, Expected: arg, Got: str}, nil
		}

	case "in":
		values := strings.Split(arg, ",")
		for _, v := range values {
			if str == v {
				return nil, nil
			}
		}

		return &ValidationError{Err: ErrIn, Expected: values, Got: str}, nil

	default:
		return nil, fmt.Errorf("%w: unknown %q", ErrInvalidRule, name)
	}

	return nil, nil
}

func validateInt(num int, rule string) (*ValidationError, error) {
	name, arg, ok := strings.Cut(rule, ":")
	if !ok {
		return nil, fmt.Errorf("%w: %q", ErrInvalidRule, rule)
	}

	switch name {
	case "min":
		n, err := strconv.Atoi(arg)
		if err != nil {
			return nil, fmt.Errorf("%w: min arg %q not int", ErrInvalidRule, arg)
		}

		if num < n {
			return &ValidationError{Err: ErrMin, Expected: n, Got: num}, nil
		}

	case "max":
		n, err := strconv.Atoi(arg)
		if err != nil {
			return nil, fmt.Errorf("%w: max arg %q not int", ErrInvalidRule, arg)
		}

		if num > n {
			return &ValidationError{Err: ErrMax, Expected: n, Got: num}, nil
		}

	case "in":
		values := strings.Split(arg, ",")
		ints := make([]int, 0, len(values))

		for _, v := range values {
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, fmt.Errorf("%w: in arg %q not int", ErrInvalidRule, v)
			}
			ints = append(ints, n)
		}

		for _, n := range ints {
			if num == n {
				return nil, nil
			}
		}

		return &ValidationError{Err: ErrIn, Expected: ints, Got: num}, nil

	default:
		return nil, fmt.Errorf("%w: unknown %q", ErrInvalidRule, name)
	}

	return nil, nil
}
