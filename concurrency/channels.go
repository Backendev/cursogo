package main

import "fmt"

func main() {
	//canal con buffer
	// c := make(chan int, 5)
	// c <- 1
	// c <- 2
	// c <- 3
	// c <- 3
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	//pipelines
	generator_c := make(chan int)
	double_c := make(chan int)
	go generator(generator_c)
	go Double(generator_c, double_c)
	Print(double_c)

}
func generator(c chan<- int) {
	for i := 1; i <= 10; i++ {
		c <- i
	}
	//cerrar canal y desbloque cualquier espera de este
	close(c)
}
func Double(in <-chan int, out chan<- int) {
	for value := range in {
		out <- 2 * value
	}
	close(out)
}
func Print(c chan int) {
	for value := range c {
		fmt.Println(value)
	}
}
