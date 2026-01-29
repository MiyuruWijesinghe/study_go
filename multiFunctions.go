package main

import "fmt"

func calculatorBasic(x int, y int) (int, int) {
	var out1 = x + y
	var out2 = x - y
	return out1, out2
}

func calculatorNew(x int, y int) (out1, out2 int) {
	out1 = x + y
	out2 = x * y
	return
}

func main() {
	num1 := 4
	num2 := 5

	result1, result2 := calculatorBasic(num1, num2)
	result3, result4 := calculatorNew(num1, num2)

	fmt.Println(result1, result2)
	fmt.Println(result3, result4)
}
