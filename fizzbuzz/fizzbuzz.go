package main

import "fmt"

func main() {

	for i := 1; i <= 25; i++ {
		printType := 0

		// maths
		if i%15 == 0 {
			printType = 15
		} else if i%5 == 0 {
			printType = 5
		} else if i%3 == 0 {
			printType = 3
		}

		// print
		switch printType {
		case 3:
			fmt.Println("Fizzles")
		case 5:
			fmt.Println("Buzzles")
		case 15:
			fmt.Println("FizzBuzzles")
		default:
			fmt.Println(i)
		}
	}
}
