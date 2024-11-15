package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func parse(str string) []string {
	var result []string
	priority := map[string]int{"+": 1, "-": 1, "*": 2, "/": 2, "(": 0}
	currentNumber := ""
	var tmp strings.Builder

	for i, char1 := range str {
		char := string(char1)

		if unicode.IsDigit(char1) || char == "." {
			currentNumber += char
		} else {
			if currentNumber != "" {
				result = append(result, currentNumber)
				currentNumber = ""
			}

			if char == "-" && (i == 0 || str[i-1] == '(' || priority[string(str[i-1])] > 0) {
				currentNumber = "-"
			} else if priority[char] >= 0 {
				if tmp.Len() > 0 {
					result = append(result, tmp.String())
					tmp.Reset()
				}
				result = append(result, char)
			}
		}
	}

	if currentNumber != "" {
		result = append(result, currentNumber)
	}

	for i := 0; i < len(result); i++ {
		if result[i] == "" || result[i] == " " {
			result = append(result[:i], result[i+1:]...)
			i--
		}
	}
	return result
}

func Operation(operation string, first float64, second float64) (float64, error) {

	switch operation {
	case "+":
		return first + second, nil
	case "-":
		return first - second, nil
	case "*":
		return first * second, nil
	case "/":
		if second == 0 {
			return 0, errors.New("ошибка - деление на ноль")
		}
		return first / second, nil
	default:
		return 0, errors.New("unknown err")
	}
}

func Calc(expression string) (float64, error) {
	priority := map[string]int{"+": 1, "-": 1, "*": 2, "/": 2, "(": 0}
	parseExpression := parse(expression)
	reverseRecord := make([]string, 0)
	operators := make([]string, 0)

	for _, a := range parseExpression {

		if a == "(" {
			operators = append([]string{a}, operators...)
		} else if priority[a] > 0 || a == ")" {
			if len(operators) == 0 {
				operators = append([]string{a}, operators...)
			} else if a == ")" {
				for {
					q := operators[0]
					operators = operators[1:]
					if q == "(" {
						break
					}
					reverseRecord = append(reverseRecord, q)
				}
			} else if priority[operators[0]] < priority[a] {
				operators = append([]string{a}, operators...)
			} else {
				for {
					if len(operators) == 0 {
						break
					}
					q := operators[0]
					reverseRecord = append(reverseRecord, q)
					operators = operators[1:]
					if priority[q] == priority[a] {
						break
					}
				}
				operators = append([]string{a}, operators...)
			}
		} else {
			//if len(parseExpression) > 0 {
			//	if (a == "(" || a == ")") && (parseExpression[len(parseExpression)-1] == "(" || parseExpression[len(parseExpression)-1] == ")") {
			//		return 0, errors.New("лол")
			//	}
			//}
			reverseRecord = append(reverseRecord, a)
		}
	}
	if len(operators) != 0 {
		for len(operators) > 0 {
			q := operators[0]
			reverseRecord = append(reverseRecord, q)
			operators = operators[1:]
		}
	} else {
		return 0, fmt.Errorf("отсутсвуют операторы")
	}

	var solutions []float64

	for _, record := range reverseRecord {
		if priority[record] > 0 {
			if len(solutions) < 2 {
				return 0, fmt.Errorf("неверный ввод (проверьте ввод операторов)")
			}
			left := solutions[len(solutions)-2]
			right := solutions[len(solutions)-1]
			solutions = solutions[:len(solutions)-2]

			res, err := Operation(record, left, right)
			if err != nil {
				return 0, err
			}
			solutions = append(solutions, res)
		} else {
			num, err := strconv.ParseFloat(record, 64)
			if err != nil {
				return 0, fmt.Errorf("введено неккоректное значение")
			}
			solutions = append(solutions, num)
		}
	}
	if len(solutions) != 1 {
		return 0, fmt.Errorf("лишние операнды")
	}
	return solutions[0], nil
}
