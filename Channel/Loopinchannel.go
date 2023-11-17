

package main

import "fmt"


func main() {

	
	mychnl := make(chan string)

	// Anonymous goroutine
	go insert(mychnl)
	/*
	go func() {
        mychnl <- "GFG"
        mychnl <- "gfg"
        mychnl <- "Geeks"
        mychnl <- "GeeksforGeeks"
        close(mychnl)
    }()
	
	*/

	// Using for loop
	for res := range mychnl {
		fmt.Println(res)
	}
}
func insert(mychnl chan   string) {
	mychnl <- "GFG"
	mychnl <- "gfg"
	mychnl <- "Geeks"
	mychnl <- "GeeksforGeeks"
	close(mychnl)
}
