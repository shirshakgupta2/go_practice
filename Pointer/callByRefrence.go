package main
import "fmt"
func main() {
	var a int =100
	var b int =300

	fmt.Printf(" before swap  a=%d b=%d\n", a, b)
	//swapByRef(&a, &b)
	//fmt.Printf(" After swap  a=%d b=%d\n", a, b)
	swapByvalue(a,b)
	fmt.Printf(" After swap  a=%d b=%d\n", a, b)

}

func swapByRef(x *int, y *int) {// ghar chnge karna 
	temp := *x
	*x=*y
	*y=temp
}
func swapByvalue(x int, y int) {// apni position chnge karna
	temp := x
	x=y
	y=temp
}