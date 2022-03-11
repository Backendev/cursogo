package main

import "fmt"

type Employee struct {
	Id   int
	name string
}

//asignacion metodos a struct
func (e *Employee) SetId(id int) {
	e.Id = id
}
func (e *Employee) SetName(name string) {
	e.name = name
}
func (e *Employee) GetId() int {
	return e.Id
}
func (e *Employee) GetName() string {
	return e.name
}

//constructor
func NewEmployee(Id int, name string) *Employee {
	return &Employee{
		Id:   Id,
		name: name,
	}
}

func main() {
	//
	e := Employee{}
	fmt.Println(e)
	e.SetId(5)
	e.SetName("Nme")
	fmt.Println(e.GetId())
	fmt.Println(e.GetName())
	//2
	e2 := Employee{
		Id:   2,
		name: "nm2",
	}
	fmt.Println(e2)
	//3
	e3 := new(Employee)
	//devuelve referencia -> apuntador
	fmt.Println(e3)
	//devuelve valor
	fmt.Println(*e3)
	e3.Id = 3
	e3.name = "nm3"
	fmt.Println(*e3)
	//4
	e4 := NewEmployee(4, "nm4")
	fmt.Println(e4)
	fmt.Println(*e4)
}
