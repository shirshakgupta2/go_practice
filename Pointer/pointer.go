package main
import "fmt"

func main()  {
	var pntr *int//astric is used to notify that its pointer and its going to store integer values
	fmt.Println("the value of pointer is ",pntr)

	myNumber:=20
	//ptr is automettically named as var ptr *int
	var ptr = &myNumber// here we are not just creating a pinter instead here  we are creating a pointer which is even refering to some value 

	//and sign means refrence

	fmt.Println("Value of new pointer is ",ptr)// pointer is a  refrence to its direct memory location 
	fmt.Println("Value of new pointer is ",*ptr)// but the value inside the pointer is given by astric

	*ptr = *ptr *2
	fmt.Println("new value is: ",myNumber)
	


}
