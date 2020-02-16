package main

import "os"

type Archivo struct {
	ruta    string
	archivo os.FileInfo
}

type ArchivoMap map[int]Archivo
