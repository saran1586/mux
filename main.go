package main

import (
"log"
"net/http"

"github.com/gorilla/mux"
)

func main() {

	
ConnectDB()
r := mux.NewRouter()

	
r.HandleFunc("/students", CreateStudent).Methods("POST")
r.HandleFunc("/students", GetAllStudents).Methods("GET")
r.HandleFunc("/students/{id}", GetStudentByID).Methods("GET")
r.HandleFunc("/students/{id}", UpdateStudent).Methods("PUT")
r.HandleFunc("/students/{id}", DeleteStudent).Methods("DELETE")

log.Println("Server Running On :8080")

err := http.ListenAndServe(":8080", r)
if err != nil {
log.Fatal(err)
}
}
