package errors

import "errors"

var (
	ErrDivisionByZero  = errors.New("ошибка - деление на ноль")
	ErrInvalidOperator = errors.New("неверный ввод (проверьте ввод операторов)")
	ErrExtraOperands   = errors.New("лишние операнды")
	ErrInvalidInput    = errors.New("введено некорректное значение")
	ErrMissingOperator = errors.New("отсутсвуют операторы")
)
