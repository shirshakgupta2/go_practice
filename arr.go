package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8} // this is the array with length  defined

	fmt.Println(arr1)
	fmt.Println(arr2)

	var arr3 = [...]int{1, 2, 3}
	arr4 := [...]int{4, 5, 6, 7, 8} // this is the array with length not defined

	fmt.Println(arr3)
	fmt.Println(arr4)

	var arr5 = [...]int{1: 10, 4: 40}
	fmt.Print(arr5," \n")

	fmt.Print("lenth of array",len(arr5),"\n")


 // var a =[...]byte{}
	var arr6 = [7]string{}
	reader := bufio.NewReader(os.Stdin)


  fmt.Println("Enter the digits in array")
	for i := 0; i < len(arr6); i++ {
    fmt.Println("Enter the ", i, "digit in array")
		x,_:= reader.ReadString('\n')
    arr6[i]=strings.TrimSpace(x)
    
	}
	fmt.Print(arr6)

}
