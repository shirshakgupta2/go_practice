package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	

	http.HandleFunc("/hello",GetBook)
	http.ListenAndServe(":9090", nil)

}

type Book struct {
	Title string `json:"title"`
	Author string `json:"author"`
	Page int `json:"page"`
}
func GetBook(resp http.ResponseWriter,req *http.Request)  {
	resp.Header().Set("Content-Type", "application/json")
	book := Book{"Charles the Hittler","shirshak",987,
		// Title :"the book",
		// Author: "shirshak",
		// Page : 450,
	}

	json.NewEncoder(resp).Encode(book)
}