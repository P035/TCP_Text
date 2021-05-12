package servidor

import (
	"bufio"
	"fmt"
	"net"
)

/*
func Respuestas(canal chan string) {

	var dato string

	for {

		dato = <-canal
	}
}
*/

func Echo(conexion *net.TCPConn, canal chan string) {

	for {

		reader := bufio.NewReader(conexion)
		data, _ := reader.ReadBytes('\n')
		if len(data) == 0 {

			conexion.Close()
			fmt.Println("Conexion cerrada")
			break
		}
		conexion.Write(data)
		fmt.Println(conexion.RemoteAddr())
		fmt.Println(data)
		fmt.Print(string(data))
	}
}
