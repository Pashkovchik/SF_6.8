package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b, c float32
	var d string
	println("Введите первое значение:")
	fmt.Fscan(os.Stdin, &a)
	println("Введите знак:")
	fmt.Fscan(os.Stdin, &d)
	println("Введите второе значение:")
	fmt.Fscan(os.Stdin, &b)

	switch d {
	case "*":
		c = a * b
		println("Результат умножения:")
		fmt.Printf("%.2f\n", c)
	case "/":
		if b != 0 {
			c = a / b
			println("Результат деления:")
			fmt.Printf("%.2f\n", c)
		} else {
			println("на ноль делить нельзя")
		}
	case "+":
		c = a + b
		println("Результат сложения:")
		fmt.Printf("%.2f\n", c)
	case "-":
		c = a - b
		println("Результат вычитания:")
		fmt.Printf("%.2f\n", c)
	default:
		println("некорректные данные")
	}
}
