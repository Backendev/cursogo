package main

import (
	//Alias pk
	pk "curso_golang_structs/src/mypackage"
	"fmt"
)

//Solucion a package is not in GOROOT
//go mod init

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) ping() {
	fmt.Println(myPc.brand, "Pong")
}

//acceder a struct puntero
func (myPc *pc) duplicateRam() {
	myPc.ram = myPc.ram * 2
}

//Stringer imprimir valores de struct
func (myPc pc) String() string {
	pct := fmt.Sprintf("Tengo %d Gb de Ram, %d gb de disco y soy un %s", myPc.ram, myPc.disk, myPc.brand)
	return pct

}

//interfaces
//structs
type cuadrado struct {
	base float64
}
type rectangulo struct {
	base   float64
	altura float64
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}
func (c cuadrado) area() float64 {
	return c.base * c.base
}

//Interface caso area
type figuras2d interface {
	area() float64
}

func calculate(f figuras2d) {
	fmt.Println("Area: ", f.area())
}

func main() {
	myCar := pk.CarPublic{Brand: "Ford", Year: 2020}
	fmt.Println(myCar)
	var otherCar pk.CarPublic
	otherCar.Brand = "Ferrari"
	otherCar.Year = 1999
	fmt.Println(otherCar)
	pk.PrintMessage("hola")

	//Punteros
	a := 50
	//variable como puntero = & accede a direccion de memoria de a
	b := &a
	fmt.Println(b)
	//* acceder al valor de la direccion de memoria
	fmt.Println(*b)
	*b = 100
	fmt.Println(*b)
	fmt.Println(a)

	myPc := pc{ram: 16, disk: 250, brand: "msi"}
	fmt.Println(myPc)
	myPc.ping()
	myPc.ram = 20
	myPc.duplicateRam()
	fmt.Println(myPc)

	//implementar interface
	myCuadrado := cuadrado{base: 4}
	myrectangulo := rectangulo{base: 4, altura: 3}

	calculate(myCuadrado)
	calculate(myrectangulo)

}
