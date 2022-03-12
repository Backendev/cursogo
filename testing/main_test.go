package main

import "testing"

//go test
//go test -cover = coverage
//go test -coverprofile=coverage.out = info coverage
//go tool cover -func=coverage.out = resumen de coverprofile
//go tool cover -html=coverage.out  = resumen en html
//profiling
//go test -cpuprofile=cpu.out
//go tool pprof cpu.out | despues | top -> muestra tiempos de ejecuciones | list Fibonacci -> tiempo por lineas | web -> info en html | pdf -> info en pdf
func TestSum(t *testing.T) {
	// total := Sum(5, 5)
	// if total != 10 {
	// 	t.Errorf("Sum was incorrect, got %d expect %d", total, 10)
	// }
	tables := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}
	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.c {
			t.Errorf("Sum was incorrect, got %d expect %d", total, item.c)
		}
	}

}
func TestGetMax(t *testing.T) {
	res := GetMax(4, 6)
	if res != 6 {
		t.Errorf("GetMax was incorrect, got %d expect %d", res, 6)
	}
	res = GetMax(8, 6)
	if res != 8 {
		t.Errorf("GetMax was incorrect, got %d expect %d", res, 8)
	}
}
func TestFib(t *testing.T) {
	tables := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}
	for _, item := range tables {
		fib := Fibonacci(item.a)
		if fib != item.n {
			t.Errorf("Fibonacci was incorrect, got %d expect %d", fib, item.n)
		}
	}
}
