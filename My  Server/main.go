package main

import "net/http"

func main() {
	// http.HandleFunc("/serverstart", func(w http.ResponseWriter, r *http.Request) { //Responce writer is a interface and Request is a struct and http is a package
	// 	w.Header().Set("content-type", "text/html")

	// 	w.Write([]byte("<html><head><title>Server Start</title></head><body><h1>Server Start hello shirshak</h1></body></html>"))

	// 	//w.Write([]byte("<head><form><p>Name:</p><input type="text" name="shirshak"><p>Mobile Number:</p><input type="number" name="name"><br><input type="Submit" name="number"></form>"))
	// 	w.Write([]byte("helo"))
	// })

	http.HandleFunc("/hello",Hello)
	http.ListenAndServe(":9090", nil)

}
func Hello(resp http.ResponseWriter,req *http.Request){
	resp.Write([]byte("hello world"))
}
