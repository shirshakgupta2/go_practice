package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter rating between 1 to 5")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("rating in string format is:", input)
	fmt.Printf("rating is in %T the rating is %v", input, input)

	// numRating, err := strconv.ParseFloat(input, 64)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Printf("rating is in %T the rating is %v \n", numRating, numRating)

		fmt.Println("rating subtracted by one")
		fmt.Println("numRating is", numRating-1)
	}
}
