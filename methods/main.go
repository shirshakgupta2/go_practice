package main
import "fmt"

func main()  {
	fmt.Println("Structs in golang")

	
	c1 := Customer{"shirshak", "shirshak@go.dev", true,24,23}


	fmt.Println(c1)
	fmt.Printf("detail of customer 1 : %+v\n", c1)
	fmt.Printf("Name of customer 1 : %v\nEmail Of customer1: %v \n", c1.Name,c1.Email)
	fmt.Printf("oneAge :=%v",c1.oneAge)
	c1.GetStatus()
	c1.newMail()
	fmt.Printf("Name of customer 1 : %v\nEmail Of customer1: %v \n", c1.Name,c1.Email)
}


type Customer struct{
	Name string
	Email string
	Status bool
	Age int
	oneAge int 
}


/*SYNTAX TO DEFINE A METHOD*/
// func(reciver_name Type) method_name(parameter_list)(return_type){
// 	 Code
// 	}


func (c Customer) GetStatus(){
	fmt.Println("\nis Customer Active: ",c.Status)
}

func (c2 Customer) newMail(){
	c2.Email="Shirshak@googlemail.com"
	fmt.Println("New Mail: ",c2.Email)
}