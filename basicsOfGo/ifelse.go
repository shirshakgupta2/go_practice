package main
import "fmt"
func main()  {
	fmt.Println("If - Else in Golang")

	loginCount := 23
	var res string
	if loginCount<10{
		 res="Regular user"
	} else if loginCount>10{
		res="watch out making alot use of the application"
	}else{
		res="Exactly 10"
	}
	fmt.Println(res)


	if num:=3;num<10{
		fmt.Println("num is less than 10")
	}else{
		fmt.Println("num is greater than 10")
	}

}