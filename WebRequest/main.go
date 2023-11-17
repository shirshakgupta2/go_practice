package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://lco.dev"

func main() {
	fmt.Println("WEB REQUEST")

	resp, err := http.Get(url)// to get request
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type %T ", resp) //

	defer resp.Body.Close() //it is caller responsibility to close the connection

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}	

	content := string(data)

	fmt.Println(content)

}
