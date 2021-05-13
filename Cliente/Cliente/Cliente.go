package Cliente

import (

	"fmt"
	"net"
	"log"
	"bufio"
	"os"
)

func Escuchar(Conexion *net.TCPConn) {

	reader := bufio.NewReader(Conexion)
	for {

		datos, err := reader.ReadBytes('\n')
		if err != nil {

			fmt.Println("Error reciviendo información con el servidor!")
			Conexion.Close()
			log.Fatal("")
		}else if string(datos) == ""{

			Conexion.Close()
			log.Fatal("Conexion perdida con el servidor.")
		}
		fmt.Print(string(datos))
	}	
}

func Conectar(argumentos []string){

	argumentos = argumentos[1:]
	var addr string
	index := 0
	var bytes []byte
	reader := bufio.NewReader(os.Stdin)
	for _, i := range argumentos {

		addr += i
		addr += ":"
		index += 1
	}
	if index != 0 {

		addr = addr[:len(addr) - 1]
	}
	direccion_remota, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {

		log.Fatal("Error creando dirección remota.")
	}
	fmt.Println("Dirección remota creada perfectamente!")
	Conexion, err := net.DialTCP("tcp", nil, direccion_remota)
	if err != nil {
	
		log.Fatal("Error estableciendo la conexion con el servidor")	
	}
	fmt.Println("Conexión establecida con el servidor")
	defer Conexion.Close() 
	go Escuchar(Conexion)
	for {
	
		bytes, _ = reader.ReadBytes('\n')
		Conexion.Write(bytes)	
	}	
}