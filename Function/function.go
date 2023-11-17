
package main

import (
	"fmt"
)

func myMessage() {
	fmt.Println("I just got executed!")
}
/*signature of a method which is user defined */
// func method_name(parameter para_type) return_type{

// }



func add() {
	fmt.Println(5+2)
}
func addition(val1 int, val2 int) int {  //method signature  is as such 
	return val1+val2
}

func proAdder(values ... int) int {//all values are int (... int)
	total := 0
	for _,val := range values {
		total += val
	}
	return total
}

func myFunction(x int,y int) int{
	return x + y
}
	// func myFunc(x int,y int) int{  
	// 	return x+y;
	// }
	
func main()  {
	

		
	myMessage() // call the function
	add()
	
	res := addition(2,7)
	fmt.Println(res)
	
	res1 := proAdder(6,8,9,6,9)
	fmt.Println("res1= ",res1)
	
	 fmt.Println(myFunction(1, 2))

}
