package servidor

import (

	"fmt"
	"net"
	"bufio"
)

func Escuchar()

func Echo(conexion net.Conn, canal chan string) {

	for {

		reader := bufio.NewReader(conexion)
		data, _ := reader.ReadBytes('\n')
		if len(data) == 0 {

			conexion.Close()
			fmt.Println("Conexion cerrada")
			break
		}
		conexion.Write(data)
		fmt.Println(data)
		fmt.Print(string(data))
	}
}
