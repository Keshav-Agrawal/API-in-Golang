package model

type Employee struct {
	Name        string `json:"fullname"`
	Employee_id string `json:"employee_id"`
	Position    string `json:"designation"`
	Salary      int    `json:"sal"`
}

var Employee_slice []Employee
