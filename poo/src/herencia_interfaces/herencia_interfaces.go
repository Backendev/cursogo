package main

import "fmt"

//Concepto de composocion
type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}
type FullTimeEmployee struct {
	Person
	Employee
	enddate string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}
func (tEmployee temporalEmployee) getMessage() string {
	return "temporal Employee"
}

type temporalEmployee struct {
	Person
	Employee
	taxrate int
}

//interfaces
type printInfo interface {
	getMessage() string
}

func getMessage(p printInfo) {
	fmt.Printf("%s desde interface\n", p.getMessage())
}

func main() {
	println("h3erencia")
	fte := FullTimeEmployee{}
	fmt.Println(fte)
	fte.name = "name f"
	fte.age = 3
	fte.id = 1
	fmt.Println(fte)
	getMessage(fte)
	te := temporalEmployee{}
	getMessage(te)
}
