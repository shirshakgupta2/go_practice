package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println(" enter the rating:")
	// comma ok syntax||err ok

	input,err := reader.ReadBytes('\n')// when we are expecting an input or if by any chance we have an error then 
	// we dont have try catch in this rather  we store it in variable as err 
	
	if err!=nil{
		panic(err)
	}
	fmt.Printf("thanks for rating %v",input)
	



}
