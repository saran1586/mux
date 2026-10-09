package main

import (
"encoding/json"
"net/http"

"github.com/gorilla/mux"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {

var student Student

err := json.NewDecoder(r.Body).Decode(&student)
if err != nil {
http.Error(w, err.Error(), http.StatusBadRequest)
return
}

err = DB.Create(&student).Error
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}

w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)

json.NewEncoder(w).Encode(student)
}


func GetAllStudents(w http.ResponseWriter, r *http.Request) {

var students []Student

err := DB.Find(&students).Error
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}

w.Header().Set("Content-Type", "application/json")

json.NewEncoder(w).Encode(students)
}


func GetStudentByID(w http.ResponseWriter, r *http.Request) {

var student Student

params := mux.Vars(r)
id := params["id"]

err := DB.First(&student, id).Error
if err != nil {
http.Error(w, "Student Not Found", http.StatusNotFound)
return
}

w.Header().Set("Content-Type", "application/json")

json.NewEncoder(w).Encode(student)
}


func UpdateStudent(w http.ResponseWriter, r *http.Request) {

var student Student

params := mux.Vars(r)
id := params["id"]

err := DB.First(&student, id).Error
if err != nil {
http.Error(w, "Student Not Found", http.StatusNotFound)
return
}


var updatedStudent Student

err = json.NewDecoder(r.Body).Decode(&updatedStudent)
if err != nil {
http.Error(w, err.Error(), http.StatusBadRequest)
return
}


student.Name = updatedStudent.Name
student.Age = updatedStudent.Age
student.Salary = updatedStudent.Salary


err = DB.Save(&student).Error
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}


w.Header().Set("Content-Type", "application/json")

json.NewEncoder(w).Encode(student)
}



func DeleteStudent(w http.ResponseWriter, r *http.Request) {

var student Student

params := mux.Vars(r)
id := params["id"]

err := DB.First(&student, id).Error
if err != nil {
http.Error(w, "Student Not Found", http.StatusNotFound)
return
}


err = DB.Delete(&student).Error
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}


w.WriteHeader(http.StatusOK)

w.Write([]byte("Student Deleted Successfully"))
}