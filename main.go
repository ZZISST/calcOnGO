package main

import "fmt"

func main() {
	var expression string
	fmt.Println("Калькулятор запущен \nВведите выражение без пробелов")

	_, errorik := fmt.Scanln(&expression)
	if errorik != nil {
		return
	}

	result, err := Calc(expression)
	if err != nil {
		fmt.Printf("Ошибка при вычислении выражения '%s': %s\n", expression, err)
	} else {
		fmt.Printf("Результат вычисления выражения '%s': %f\n", expression, result)
	}

}
