package seki

import "fmt"

type Validator interface {
	// Valid checks the object and returns any
	// problems. If len(problems) == 0 then
	// the object is valid.
	Valid() (problems map[string]string)
}

func StringMinLength(minLength int, value any) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("is not a string")
	}
	if len(str) < minLength {
		return fmt.Errorf("must be more than or equal to %d characters for value %v", minLength, value)
	}
	return nil
}

func StringMaxLength(maxLength int, value any) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("is not a string for value %v", value)
	}
	if len(str) > maxLength {
		return fmt.Errorf("must be less than or equal to %d characters for value %v", maxLength, value)
	}
	return nil
}

func IsString(value any) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("is not a string for value %v", value)
	}

	if str == "" || len(str) < 1 {
		return fmt.Errorf("can't be blank for value %v", value)
	}
	return nil
}

func IsInteger(value any) error {
	integer, ok := value.(int)
	if !ok {
		return fmt.Errorf("is not a integer for value %v", value)
	}

	if integer < 1 {
		return fmt.Errorf("can't be negative for value %v", value)
	}
	return nil
}

func IsEnum(value any) error {
	// TODO
	integer, ok := value.(int)
	if !ok {
		return fmt.Errorf("is not a enum for value %v", value)
	}

	if integer < 1 {
		return fmt.Errorf("is not a enum for value %v", value)
	}
	return nil
}

func IsFloat(value any) error {
	// TODO
	integer, ok := value.(int)
	if !ok {
		return fmt.Errorf("is not a enum for value %v", value)
	}

	if integer < 1 {
		return fmt.Errorf("is not a enum for value %v", value)
	}
	return nil
}

func IntInRange(min int, max int, value any) error {
	// TODO
	integer, ok := value.(int)
	if !ok {
		return fmt.Errorf("is not a enum")
	}

	if integer < 1 {
		return fmt.Errorf("is not a enum for value %v", value)
	}
	return nil
}

func IsEmail(value any) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("is not a string for value %v", value)
	}
	// simplified email validation for example purposes
	if str == "" || str[:1] == "@" || str[len(str)-1:] == "@" {
		return fmt.Errorf("is not a valid email address for value %v", value)
	}
	return nil
}
