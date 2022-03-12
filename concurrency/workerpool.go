package main

import "fmt"

//worker pool
func main() {
	tasks := []int{2, 3, 4, 5, 7, 10}
	nWorkers := 3
	jobs_c := make(chan int, len(tasks))
	results_c := make(chan int, len(tasks))

	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs_c, results_c)

	}

	for _, value := range tasks {
		jobs_c <- value
	}
	close(jobs_c)

	for r := 0; r < len(tasks); r++ {
		<-results_c
	}
}

func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker whit id %v started fib with %v\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Worker with id %v,job %v and fib %v\n", id, job, fib)
		results <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
