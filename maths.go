package main

import (
	"fmt"
	"math"
)

func main() {

	var num float64 = 12

	var result = math.Sqrt(num)

	fmt.Printf("The output is %.2g thanks", result) //%.2g used in Printf for formatting
	fmt.Println("")

	fmt.Printf("The output is %.2f thanks", result) //%.2f used in Printf for formatting
	fmt.Println("")

	var intRounded = math.Round(result)
	fmt.Printf("The rounded output is %.2f thanks", intRounded) //Round function
	fmt.Println("")

	var intCeil = math.Ceil(result)
	fmt.Printf("The ceil output is %.2f thanks", intCeil) //Ceil function
	fmt.Println("")

	var intFloor = math.Floor(result)
	fmt.Printf("The floor output is %.2f thanks", intFloor) //Floor function
}
