package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает
// две числовых переменных a,b, значение которых > 2^20.

func calculate(num1, num2 *big.Int, operator string) *big.Int {
	// Для операций с большими числами используем пакет math/big
	var ans big.Int
	switch operator {
	case "+":
		return ans.Add(num1, num2)
	case "-":
		return ans.Sub(num1, num2)
	case "*":
		return ans.Mul(num1, num2)
	case "/":
		return ans.Div(num1, num2)
	default:
		return nil
	}
}

func main() {
	// Большие числа задаются с помощью строк
	num1,_ := new(big.Int).SetString("858192767401055658192", 10)
	num2,_ := new(big.Int).SetString("1258124665061271231", 10)
	
	fmt.Println(calculate(num1, num2, "+"))
	fmt.Println(calculate(num1, num2, "-"))
	fmt.Println(calculate(num1, num2, "*"))
	fmt.Println(calculate(num1, num2, "/"))
}
