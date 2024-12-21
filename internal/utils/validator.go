package utils

import "regexp"

// ValidateExpression проверяет, что строка выражения содержит только цифры, операции и допустимые символы.
func ValidateExpression(expression string) bool {
	validExpression := regexp.MustCompile(`^[0-9+\-*/(). ]+$`)
	return validExpression.MatchString(expression)
}
