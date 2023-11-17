package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	
)

func main() {
	fmt.Println("Welcome to file handling")
	content :="hello shirshak how are you"

	file,err := os.Create("./myfilehandling.txt")// for creating  a file handle os package or module is used
	ErrorHandling(err)
	fmt.Println(file.Name())
	len,err := io.WriteString(file,content)

	ErrorHandling(err)
	fmt.Println("Length is ",len)

	
	defer file.Close()

	readFile(file.Name())
}


func readFile(filename string){
	data,err := ioutil.ReadFile(filename)
	ErrorHandling(err)
	fmt.Println("Text in file ",string(data))
}
func ErrorHandling(err error) {
	if(err!=nil){
		panic(err)
	}
}
