package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ObtenerArchivos(ruta, validos string) map[int]Archivo {
	archivos := make(map[int]Archivo)

	err := filepath.Walk(ruta, procesarRuta(validos, &archivos))
	if err != nil {
		log.Println(err)
	}

	return archivos
}

func procesarRuta(validos string, archivos *map[int]Archivo) filepath.WalkFunc {
	c := 0

	return func(ruta string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		ext := filepath.Ext(ruta)
		if !info.IsDir() && ext != ".go" && TipoValido(validos, ext) {
			(*archivos)[c] = Archivo{
				ruta:    ruta,
				archivo: info,
			}
			c++
		}
		return nil
	}
}

func ExisteRuta(ruta string) bool {
	if _, err := os.Stat(ruta); os.IsNotExist(err) {
		return false
	}
	return true
}

func RenombrarArchivo(archivo Archivo) bool {
	nombreViejo := archivo.ruta
	baseDir := filepath.Dir(nombreViejo)
	nombreNuevo := strings.ReplaceAll(filepath.Base(nombreViejo), " ", "*")
	rutaNueva := baseDir + "/" + nombreNuevo
	fmt.Println(nombreViejo, rutaNueva)
	return true
}

func TipoValido(validos, ext string) bool {
	if strings.TrimSpace(validos) == "all" {
		return true
	} else {
		partes := strings.Split(validos, ",")
		for _, tipo := range partes {
			if ("." + strings.TrimSpace(tipo)) == ext {
				return true
			}
		}
		return false
	}
}
