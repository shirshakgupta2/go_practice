package main

import (
	"fmt"
	"net/http"
)

func formHandler(resp http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Printf("response parse error  and the error : %v ", err)
	}
	fmt.Fprintf(resp, "Postrequest \n")
	name := req.FormValue("name") // similar to req.getParameter("name") in java
	address := req.FormValue("address")

	fmt.Fprintf(resp, "response for name : %v \n", name)
	fmt.Fprintf(resp, "response for address : %v \n", address)
}

func helloHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	if req.URL.Path != "/hello" {
		http.Error(resp, "Invalid URL", http.StatusNotFound)
		return
	}

	if req.Method != "GET" {
		http.Error(resp, "method not found", http.StatusNotFound)
		return
	}

	fmt.Println("HELLO  this is my intial to web server")
	fmt.Fprintf(resp, "HELLO  this is my intial to web server")

}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	fmt.Println("Going to fileserver")
	http.Handle("/", fileServer)
	fmt.Println("going to form handler")
	http.HandleFunc("/form", formHandler) // this is same as creating a web servlet
	fmt.Println("going to hello handler")
	http.HandleFunc("/hello", helloHandler)
	
	fmt.Println("Starting server")
	http.ListenAndServe(":9090", nil)

}
