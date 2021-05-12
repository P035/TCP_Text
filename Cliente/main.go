package main

import (

	"os"
	"github.com/P035/TCP_Text/Cliente/Cliente"
)

func main() {

	Cliente.Conectar(os.Args)
}