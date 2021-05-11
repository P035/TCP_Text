package main

import (

	"fmt"
	"github.com/P035/TCP_Text/servidor"
)

func main(){

	fmt.Println("Hola Mundo")
	Listener := servidor.Inicializar()
	servidor.Manejar_Func = servidor.Echo
	servidor.Escuchar(Listener)
}