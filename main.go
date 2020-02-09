package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]

	var ruta, validos string

	if len(args) > 1 {
		ruta = args[0]
		validos = args[1]
	} else {
		fmt.Println("uso: [ruta] [tipos validos: all o png,pdf]")
		return
	}

	if ExisteRuta(ruta) {

		archivos := ObtenerArchivos(ruta, validos)

		for _, f := range archivos {
			RenombrarArchivo(f)
		}
	} else {
		fmt.Printf("No existe el directorio \"%s\". Comprueba el directorio\n", ruta)
	}
}
