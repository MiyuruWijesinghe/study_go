package main

import "fmt"

func printNumbers(num int) {
	if num == 1 {
		fmt.Println("One")
	} else if num == 2 {
		fmt.Println("Two")
	} else {
		fmt.Println("None")
	}
}

func switchGender(num int) {
	switch num {
	case 1:
		fmt.Println("A boy")
	case 2:
		fmt.Println("A girl")
	default:
		fmt.Println("Alien")
	}
}

func main() {

	num := 6

	if num < 5 {
		fmt.Println("Hi ", num)
	} else {
		fmt.Println("Bye ", num)
	}

	printNumbers(num)
	switchGender(num)
}
