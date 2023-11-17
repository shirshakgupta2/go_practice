package main

import (
	"fmt"
)

var students string

// x:=25  should not be defined outside the function
const LoginToken = "login" //  as we make it capital letter for variable name it becomes PUBLIC
func main() {
	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x := 2                       //type is inferred
	var y string
	var i int
	var j bool
	students = "hello shirshak"
	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(students)

	fmt.Printf("the type of the variable is %T and the value is %v", LoginToken, LoginToken)
}
