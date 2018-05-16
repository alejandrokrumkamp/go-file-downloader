package main

import (
	"net/http"
	"io/ioutil"
	"path/filepath"
	"os"
	"time"
)

func createDownloadsDirectoryIfNotExists(downloadsDirectory string){
	_,err := os.Stat(downloadsDirectory)
	if err != nil {
		//
	}
	os.MkdirAll(downloadsDirectory, os.ModePerm)
}

func createFile(data []byte,downloadsDirectory string ,path string){

	fileName := downloadsDirectory+filepath.Base(path)
	println("Guardando "+fileName+" en disco...")

	err2 := ioutil.WriteFile(fileName, data, 0644)
	if err2 != nil {
		panic("Error al guardar el archivo en disco")
	}
}

func getFile (url string) []byte{
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		panic("Error al obtener el archivo de la ruta")
	}

	println("Descargando archivo desde "+url)
	body, err := ioutil.ReadAll(resp.Body)
	return body
}

func main() {
	start := time.Now()
	// "http://ubuntu.unc.edu.ar/releases/bionic/ubuntu-18.04-live-server-amd64.iso"
	path := "https://pbs.twimg.com/profile_images/554798224154701824/mWd3laxO_400x400.png"
	downloadsDirectory := "Downloads/"

	createDownloadsDirectoryIfNotExists(downloadsDirectory)
	data := getFile("https://pbs.twimg.com/profile_images/554798224154701824/mWd3laxO_400x400.png")
	createFile(data,downloadsDirectory,path)

	t := time.Now()
	elapsed := t.Sub(start)/1000/1000
	println("Descarga completada en ",elapsed,"ms!")
}

