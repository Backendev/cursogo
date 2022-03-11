package main

import (
	"github.com/donvito/hellomod"
	hellomod2 "github.com/donvito/hellomod/v2"
)

//go get modulo
//go mod tidy - Elimina dependenbcias no usadas
//go mod verify -- verificar modulos
//go mod download -json -Info de los modulos cacheados
func main() {
	hellomod.SayHello()
	hellomod2.SayHello("ed ")
}
