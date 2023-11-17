package main
import "fmt"

func main(){
	var a int
	var ptr *int
	var pntr **int
	a=100
	ptr=&a
	pntr=&ptr

	fmt.Println("The value of a is: ",a)
	fmt.Println("The value of a is: ",ptr)
	fmt.Println("The value of a is: ",*ptr)
	fmt.Println("The value of a is: ",pntr)
	fmt.Println("The value of a is: ",*pntr)
	fmt.Println("The value of a is: ",**pntr)
}