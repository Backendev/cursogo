package main

import "fmt"

func suma(a, b int) int {
	return a + b
}
func getValues(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = x * 4
	return
}

//funciones variadicas
func sumav(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}
func main() {
	x := 5
	//Funciones Anonimas
	//Cuando se va a usar una sola vez
	y := func() int {
		return x * 2
	}
	z := func() int {
		return x * 2
	}()
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println(suma(1, 2))
	fmt.Println(sumav(1, 2, 3, 4, 5))
	fmt.Println(getValues(3)["quad"])

}
