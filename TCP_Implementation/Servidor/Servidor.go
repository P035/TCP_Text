package Servidor

import (

	"fmt"
//	"log"
	"net"
)

const Datos = "192.168.1.5:5000"

func Inicializar(){

	Direccion, err := net.ResolveTCPAddr("tcp", Datos)
	if err == nil{

		fmt.Println(Direccion.Network())
	}

}