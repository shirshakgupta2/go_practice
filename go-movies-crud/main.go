package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Movie struct {
	ID       string    `json:"id"`
	isbn     string    `json:"isbn"`
	title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie //slice for movies
func getMovie(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(movies)

}

func deleteMovie(resp http.ResponseWriter, req *http.Request)  {
	resp.Header().Set("Content-Type", "application/json")
	
}
func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", isbn: "425166", title: "Movie One", Director: &Director{Firstname: "firstdirector", Lastname: "firstdirectorLastname"}})
	movies = append(movies, Movie{ID: "2", isbn: "516521", title: "Movie Two", Director: &Director{Firstname: "Seconddirector", Lastname: "SeconddirectorLastname"}})
	movies = append(movies, Movie{ID: "3", isbn: "656516", title: "Movie Three", Director: &Director{Firstname: "Thirddirector", Lastname: "ThirddirectorLastname"}})
	movies = append(movies, Movie{ID: "4", isbn: "654564", title: "Movie Four", Director: &Director{Firstname: "Fourthdirector", Lastname: "FourthdirectorLastname"}})

	r.HandleFunc("/movies", getMovies).method("GET") //if we dont write method then also default  cosidered as get
	r.HandleFunc("/movies/{id}", getMovie).method("GEt")
	r.HandleFunc("/movies", createMovie).method("POST")
	r.HandleFunc("movies/{id}", updateMovie).method("POST")
	r.HandleFunc("movies/{id}", deleteMovie).method("DELETE")

	fmt.Println("STARTING SERVER AS PORT 9090")
	log.Fatal(http.ListenAndServe(":9090", nil))

}
