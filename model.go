package main

type Student struct {
	ID     uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Salary float64 `json:"salary"`
}
