package main
import "fmt"

func main()  {
	defer fmt.Println("WORLD")

	//  (LIFO)  and defer happens to be
	// executed at the last of the program 
	defer fmt.Println("THREE")
	defer fmt.Println("\nONE")
	defer myDefer()
	fmt.Println("\nHello")
	myDefer()
}

func myDefer()  {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i," ")
	}
}