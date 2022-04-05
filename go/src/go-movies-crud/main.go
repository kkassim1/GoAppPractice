package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)


type movie struct{
ID string `json:"id"`
Isbn string `json:"isbn"`
Title string `json:"title"`
Director *Director `json:"director"`

}

type Director{
Firstname string `json:"firstname"`
Lastname string `json:"lastname"`

}


var movies []movie

func main(){

r:= mux.NewRouter("/movies" , getMovies).Methods("GET")

r.HandleFunc("/movies/{id}" , getMovie).Methods("GET")

r.HandleFunc("/movies", createMovie).Methods("POST")

r.HandleFunc("/movies/{id}" , updateMovie).Methods("PUT")

r.HandleFunc("/movies/{id}" , deleteMovie).Methods("DELETE")


}