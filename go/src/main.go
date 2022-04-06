package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Student struct {
	ID          string   `json:"id"`
	Class       string   `json:"class"`
	StudentName string   `json:"studentname"`
	Teacher     *Teacher `json:"teacher"`
}

type Teacher struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var Students []Student

func getStudents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Students)

}

func deleteStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range Students {

		if item.ID == params["id"] {
			Students = append(Students[:index], Students[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Students)
}

func getStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range Students {
		if item.ID == params["id"] {

			json.NewEncoder(w).Encode(item)
			return

		}

	}

}

func createStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var student Student

	_ = json.NewDecoder(r.Body).Decode(&student)
	student.ID = strconv.Itoa(rand.Intn(100000000))
	Students = append(Students, student)
	json.NewEncoder(w).Encode(student)

}

func updateStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range Students {

		if item.ID == params["id"] {
			Students = append(Students[:index], Students[index+1:]...)
			var student Student

			_ = json.NewDecoder(r.Body).Decode(&student)
			student.ID = params["id"]
			Students = append(Students, student)
			json.NewEncoder(w).Encode(student)
		}

	}

}

func main() {

	r := mux.NewRouter()

	Students = append(Students, Student{ID: "1", Class: "32122", StudentName: "stu 1", Teacher: &Teacher{Firstname: "john", Lastname: "doe"}})

	Students = append(Students, Student{ID: "2", Class: "32122", StudentName: "stu 2", Teacher: &Teacher{Firstname: "kkn", Lastname: "yup"}})
	r.HandleFunc("/students", getStudents).Methods("GET")
	r.HandleFunc("/students/{id}", getStudent).Methods("GET")

	r.HandleFunc("/students", createStudent).Methods("POST")

	r.HandleFunc("/students/{id}", updateStudent).Methods("PUT")

	r.HandleFunc("/students/{id}", deleteStudent).Methods("DELETE")

	fmt.Print("starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
