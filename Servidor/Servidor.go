package servidor

import (
	"fmt"
	"log"
	"net"
)

const datos = "192.168.1.5:3000"

var Conexiones int = 0
var Canal_General chan []byte
var Manejar_Func func(conexion *net.TCPConn, canal chan []byte, conexiones *[]*net.TCPConn) = nil
var Conexiones_Server []*net.TCPConn
// La función manejar se puede redefinir segun el valor que este en Manejar_Func.
func Manejar(conexion *net.TCPConn) {

	Manejar_Func(conexion, Canal_General, &Conexiones_Server)
}

// Creara un listener y lo devolverá
func Inicializar() *net.TCPListener {

	fmt.Println("Inicializando Listener.")
	Direccion, err := net.ResolveTCPAddr("tcp", datos)
	if err != nil {

		fmt.Println(err)
		log.Fatal("Error declarando la dirección!")
	}
	Listener, err := net.ListenTCP("tcp", Direccion)
	if err != nil {

		fmt.Println(err)
		log.Fatal("Error creando listener.")
	} else {

		fmt.Println("Listener inicializado perfectamente.")
	}
	return Listener
}

// Creara y aceptará conexiones
func Escuchar(Listener *net.TCPListener) bool {

	defer Listener.Close()
	for {

		conexion, err := Listener.AcceptTCP()
		if err != nil {

			fmt.Println("Error aceptando conexion.")
		} else {

			fmt.Println("Conexión aceptada. Porcediendo a manejarla")
			Conexiones += 1
			Conexiones_Server = append(Conexiones_Server, conexion)
			fmt.Println("Número de conexiones: ", Conexiones)
			go Manejar(conexion)
		}
	}
}
