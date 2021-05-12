package Cliente

import (

	"fmt"
//	"net"
//	"log"
)

func Conectar(argumentos []string){

	var addr string
	for _, i := range argumentos {

		addr += i		
	}
	fmt.Println(addr)
}