package servidor

import (
	"bufio"
	"fmt"
	"net"
)
/*
func Respuestas(conexion *net.TCPConn, canal *chan []byte) {

	var dato []byte

	for {

		fmt.Println("Hola")
		dato = <-*canal
		fmt.Println(string(dato))
		conexion.Write(dato)
	}
}
*/

func Echo(conexion *net.TCPConn, canal chan []byte, conexiones *[]*net.TCPConn) {

	//go Respuestas(conexion, &canal)
	index := 0
	var nombre string
	var nombre_real string
	for {

		
		reader := bufio.NewReader(conexion)
		data, _ := reader.ReadBytes('\n')
		if index == 0 {

			nombre = string(data)
		}else {

			if len(data) == 0 {

				conexion.Close()
				fmt.Println("Conexion cerrada")
				break
			}
			for indice, i := range nombre {

				if indice == len(nombre) - 1 {

					break
				}else if string(i) != "\r"{
				
					nombre_real += string(i)
				}
			}
			fmt.Println(nombre_real)
			mensaje := nombre_real + "> " + string(data) + "\n"
			fmt.Print(mensaje)
			for _, value := range *conexiones {

 				if value != conexion {
					value.Write([]byte(mensaje))
				}else if value == conexion{

					mensaje = "Tu" + "> " + string(data) + "\n"
				}
			}
		}
		fmt.Println(conexion.RemoteAddr())
		fmt.Println(data)
		fmt.Print(string(data))
		index++
	}
}
