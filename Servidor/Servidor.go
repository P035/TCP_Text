package servidor

import (

	"fmt"
	"net"
	"log"
)

const datos = "127.0.0.1:3000"
var Conexiones int = 0
var Manejar_Func func(conexion net.Conn) = nil

// La función manejar se puede redefinir segun el valor que este en Manejar_Func.
func Manejar(conexion net.Conn) {

	Manejar_Func(conexion)
}

func Limpiar() {

}

// Creara un listener y lo devolverá
func Inicializar() net.Listener{

	fmt.Println("Inicializando Listener.")
	Listener, err := net.Listen("tcp", datos)
	if err != nil {
	
		log.Fatal("Error creando listener.")
	}else {

		fmt.Println("Listener inicializado perfectamente.")
	}
	return Listener
}

// Creara y aceptará conexiones
func Escuchar(Listener net.Listener) bool {

	defer Listener.Close()
	for {

		conexion, err := Listener.Accept()		
		if err != nil {
		
			fmt.Println("Error aceptando conexion.")
		}else {
		
			fmt.Println("Conexión aceptada. Porcediendo a manejarla")
			if Manejar_Func == nil{

				fmt.Println("La función para manejar las conexiones no ha sido redefinida.")
				fmt.Println("Cerrando las conexiones!")
				conexion.Close()
				log.Fatal("Conexiones cerradas!")
			}
			Conexiones += 1
			fmt.Println("Número de conexiones: ", Conexiones)
			go Manejar(conexion)
		}
	}	
}