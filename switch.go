package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// reader=bufio.NewWriter(io.Stdin)
	reader := bufio.NewReader(os.Stdin)
	day, _ := reader.ReadString('\n')
	//day :=
	input,err:= strconv.ParseFloat(strings.TrimSpace(day),64)

	if(err!=nil) {
		fmt.Println(err)
	} else {
		switch input {
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wednesday")
		case 4:
			fmt.Println("Thursday")
		case 5:
			fmt.Println("Friday")
		case 6:
			fmt.Println("Saturday")
		case 7:
			fmt.Println("Sunday")
		}

		

		switch input {
		case 1,3,5:
			fmt.Println("Odd weekday")
		case 2,4:
			fmt.Println("Even weekday")
		case 6,7:
			fmt.Println("Weekend")
		default:
			fmt.Println("Invalid day of day number")
		}
		
	}

}
