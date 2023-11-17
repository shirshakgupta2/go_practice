package main

import "fmt"

func main() {
	fmt.Println("Channel in golang")
	/*In Go language, a channel is created using chan keyword and it can only transfer
	 data of the same type, different types of data are not allowed to transport from the same channel.*/
/****SYNTAX*****/
	 //channel_name:= make(chan Type)
	 var mychannel chan int
	 fmt.Println("Value of the channel: ", mychannel)
	 fmt.Printf("Type of the channel: %T ", mychannel)
   
	 // Creating a channel using make() function
	 mychannel1 := make(chan int)
	 fmt.Println("\nValue of the channel1: ", mychannel1)
	 fmt.Printf("Type of the channel1: %T ", mychannel1)
 }

}