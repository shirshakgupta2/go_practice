// Go program to illustrate the
// concept of select statement
package main

import "fmt"

// main function
func main() {

// creating channel
mychannel:= make(chan int)
select{
case <- mychannel:

default:fmt.Println("Not found")
}
}
/*
The blocking of select statement means when there is no case statement is ready and
the select statement does not contain any default statement, then the select statement
block until at least one case statement or communication can proceed.
*/