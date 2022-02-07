package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Employee struct {
	Name        string `json:"fullname"`
	Employee_id string `json:"employee_id"`
	Position    string `json:"designation"`
	Salary      int    `json:"sal"`
}

var Employee_slice []Employee

func main() {
	fmt.Println("Welcome to the CRUD application")
	Employee_slice = append(Employee_slice, Employee{Name: "Keshav", Employee_id: "1", Position: "Golang Trainee", Salary: 450000})
	Employee_slice = append(Employee_slice, Employee{Name: "Vatsal", Employee_id: "2", Position: "Java Trainee", Salary: 350000})
	r := mux.NewRouter()
	r.HandleFunc("/", StartPage).Methods("GET")
	r.HandleFunc("/employees", DisplayAllEmployees).Methods("GET")
	r.HandleFunc("/employee/{id}", DisplaySingleEmployee).Methods("GET")
	r.HandleFunc("/employee", CreateEmployee).Methods("POST")
	r.HandleFunc("/employee/{id}", DeleteEmployee).Methods("DELETE")
	r.HandleFunc("/employee/{id}", UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employee", DeleteAllEmployee).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":4000", r))
}
func StartPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API with all CRUD operations (Employee system)</h1>"))
}
func DisplayAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatioan/json")
	json.NewEncoder(w).Encode(Employee_slice)
}
func DisplaySingleEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatioan/json")
	param := mux.Vars(r)
	for _, employee := range Employee_slice {
		if employee.Employee_id == param["id"] {
			json.NewEncoder(w).Encode(employee)
			return
		}
	}
	json.NewEncoder(w).Encode("No Employee with given employee id")
	return
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatioan/json")
	var employee Employee
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	_ = json.NewDecoder(r.Body).Decode(&employee)
	if employee.Employee_id == "" {
		json.NewEncoder(w).Encode("employee id cannot be empty")
		return
	}
	for _, singlemeployee := range Employee_slice {
		if employee.Employee_id == singlemeployee.Employee_id {
			json.NewEncoder(w).Encode("Employee id already exists")
			return
		}
	}
	Employee_slice = append(Employee_slice, employee)
	json.NewEncoder(w).Encode("Successfully added in DB")
	json.NewEncoder(w).Encode(employee)
	return
}
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	for index, employee := range Employee_slice {
		if employee.Employee_id == param["id"] {
			Employee_slice = append(Employee_slice[:index], Employee_slice[index+1:]...)
			json.NewEncoder(w).Encode("Employee Deleted")
			return
		}
	}
	json.NewEncoder(w).Encode("No such employee id present")
}
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	for index, employee := range Employee_slice {
		if employee.Employee_id == param["id"] {
			Employee_slice = append(Employee_slice[:index], Employee_slice[index+1:]...)
			fmt.Println("(SUCCESS)Employee found to be updated")
			fmt.Println(employee)
			x := employee.Employee_id
			var updatedemployee Employee
			json.NewDecoder(r.Body).Decode(&updatedemployee)
			if updatedemployee.Name == "" || updatedemployee.Position == "" || updatedemployee.Salary == 0 {
				json.NewEncoder(w).Encode("No field can be nil except Employee id")
				return
			}
			updatedemployee.Employee_id = x
			Employee_slice = append(Employee_slice, updatedemployee)
			json.NewEncoder(w).Encode("Employee has been updated successfully")
			return
		}
	}
	json.NewEncoder(w).Encode("No employee with the given id is present in the database to be updated")

}
func DeleteAllEmployee(w http.ResponseWriter, r *http.Request) {
	Employee_slice = append(Employee_slice[:0], Employee_slice[:0]...)
	json.NewEncoder(w).Encode("All the employees have been deleted")
}
