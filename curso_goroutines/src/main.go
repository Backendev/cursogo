package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)

}

func main() {
	//acumular conjunto de go ro y acumularla poco a poco
	var wg sync.WaitGroup
	fmt.Println("Hello")
	//añadir una go routine
	wg.Add(1)
	//se envia puntero a func
	go say("World", &wg)
	//esperar hasta que todas la goroutines se ejecuten
	wg.Wait()

	go func(text string) {
		fmt.Println(text)
	}("Adios")
	time.Sleep(time.Second * 1)

	//Chanels
	//creacion del canal (Tipo tipo_que_recibe cuantos cuantos_datos_simultaneos)
	c := make(chan string, 1)
	fmt.Println("hello")
	//llamada de funcion con go routine y se pasa el canal que guardara el valor
	//genberado en la funcion
	go say2("bye", c)
	//se extrae el valor generado en la funcion y guardado en el canal
	fmt.Println(<-c)

	//metodos chanels
	c2 := make(chan string, 2)
	c2 <- "mensaje 1"
	c2 <- "mensaje 2"
	// ya se cumplio con la cantidad del canal no va a recibir mas datos
	fmt.Println("Tamaño ", len(c2), "capacidad ", cap(c2))
	//Range y close
	//cerrar canal no va a recibir mas datos adicionales
	close(c2)
	//si no nse cierra el canal el for seguira esperando datos de este
	for message := range c2 {
		fmt.Println(message)
	}
	//select
	email1 := make(chan string)
	email2 := make(chan string)
	go message("msg 1", email1)
	go message("msg 2", email2)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email 1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email 2", m2)
		}
	}
}

//
// func say2(text string, c chan string) {
// 	c <- text
// }

//Por mejor practica definir si el canal va a ser de entrada o de salida
//Entrada =  c chan<- | salida = c <-chan
func say2(text string, c chan<- string) {
	c <- text
}

func message(text string, c chan string) {
	c <- text
}
