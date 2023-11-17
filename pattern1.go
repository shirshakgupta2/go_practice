package main

import "fmt"

func main() {
	fmt.Println("print the pattern ")

	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
