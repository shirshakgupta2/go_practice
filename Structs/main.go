package main
import "fmt"
func main()  {
	fmt.Println("Structs in golang")

	//no inheritance in golang; no super,no parent,no child
	// object_names := struct_name{variable 1,variable 2,...}
	c1 := Customer{"shirshak", "shirshak@go.dev", true,24}


	fmt.Println(c1)
	fmt.Printf("detail of customer 1 : %+v\n", c1)
	fmt.Printf("Name of customer 1 : %v\nEmail Of customer1: %v", c1.Name,c1.Email)
}

// type struct_name struct {
// 	member1 datatype;
// 	member2 datatype;
// 	member3 datatype;
// 	...
// }

type Customer struct{//structure struct 
	Name string
	Email string
	Status bool
	Age int
}