package main

type monthlySalary struct {
	Salary float64 `json:"salary" binding:"required,min=1.0"`
	Month  int     `json:"month" binding:"required,max=12,min=1"`
}
