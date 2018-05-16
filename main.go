package main

import (
	"net/http"
	"io/ioutil"
	"path/filepath"
	"os"
)

func main() {
	path := "https://pbs.twimg.com/profile_images/554798224154701824/mWd3laxO_400x400.png"
	downloadsDirectory := "Downloads/"

	_,err := os.Stat(downloadsDirectory)
	if err != nil {
		//
	}
	os.MkdirAll(downloadsDirectory, os.ModePerm)

	//resp, err := http.Get("http://ubuntu.unc.edu.ar/releases/bionic/ubuntu-18.04-live-server-amd64.iso")
	resp, err := http.Get(path)
	defer resp.Body.Close()
	if err != nil {
		panic("Error al obtener el archivo de la ruta")
	}

	println("Descargando archivo desde "+path)
	body, err := ioutil.ReadAll(resp.Body)

	fileName := downloadsDirectory+filepath.Base(path)
	println("Guardando "+fileName+" en disco...")

	err2 := ioutil.WriteFile(fileName, body, 0644)
	if err2 != nil {
		panic("Error al guardar el archivo en disco")
	}

	println("Descarga completada!")
}

