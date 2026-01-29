package main

import "fmt"

func main() {
	var num1 int = 2
	var num2 int = 3
	var num3 int                     //when declared as int without assigning value then default is 0
	var result1 = num1 + num2 + num3 //data type is not mandatory for variables -> result1 variable

	var num4, num5 int //can declare variables in same line
	num4, num5 = 4, 5  //can assign values in same line
	num6 := 6          // same as var num6 = 6

	const num7 = 7 //constant variable

	var result2 = num4 + num5 + num6 + num7
	var result = result1 + result2

	fmt.Print(result)
}
