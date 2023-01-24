package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// carpeta a recorrer
	carpeta := "C:/PRUEBA_T/enron_mail_20110402/enron_mail_20110402"
	// archivo de salida
	f, err := os.Create("salidax.ndjson")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// recorrer la carpeta
	err = recorrerCarpeta(carpeta, f)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func recorrerCarpeta(carpeta string, f *os.File) error {
	// obtener contenido de la carpeta
	archivos, err := ioutil.ReadDir(carpeta)
	if err != nil {
		return err
	}

	// recorrer los archivos
	for _, archivo := range archivos {
		// obtener ruta completa del archivo
		ruta := carpeta + "/" + archivo.Name()

		// si es una carpeta, se recorre
		if archivo.IsDir() {
			err = recorrerCarpeta(ruta, f)
			if err != nil {
				return err
			}
		} else {
			// leer contenido del archivo
			contenido, err := ioutil.ReadFile(ruta)
			if err != nil {
				return err
			}

			header := map[string]interface{}{
				"index": map[string]string{
					"_index": "correos",
				},
			}

			data := map[string]string{"archivo": archivo.Name(), "contenido": string(contenido)}

			jsonDataHeader, err := json.Marshal(header)
			if err != nil {
				return err
			}

			_, err = f.Write(append(jsonDataHeader, byte('\n')))
			if err != nil {
				return err
			}

			jsonData, err := json.Marshal(data)
			if err != nil {
				return err
			}

			_, err = f.Write(append(jsonData, byte('\n')))
			if err != nil {
				return err
			}

		}
	}
	return nil
}
