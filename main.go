package main

import (
	"net/http"
	"path/filepath"
	"os"
	"time"
	"io"
	"fmt"
)

func createDownloadsDirectoryIfNotExists(downloadsDirectory string){
	_,err := os.Stat(downloadsDirectory)
	if err != nil {
		//
	}
	os.MkdirAll(downloadsDirectory, os.ModePerm)
}

func createFile(downloadsDirectory string ,path string) *os.File{

	fileName := downloadsDirectory+filepath.Base(path)
	println("Guardando "+fileName+" en disco...")

	file, err2 := os.Create(fileName)
	if err2 != nil {
		panic("Error al guardar el archivo en disco")
	}
	return file
}

func downloadFile (url string,file *os.File){
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		panic("Error al obtener el archivo de la ruta")
	}

	println("Descargando archivo desde "+url)
	io.Copy(file,resp.Body)
}

func main() {
	start := time.Now()
	// "http://ubuntu.unc.edu.ar/releases/bionic/ubuntu-18.04-live-server-amd64.iso"
	path := "https://pbs.twimg.com/profile_images/554798224154701824/mWd3laxO_400x400.png"
	downloadsDirectory := "Downloads/"

	createDownloadsDirectoryIfNotExists(downloadsDirectory)
	file := createFile(downloadsDirectory,path)

	downloadFile(path,file)

	elapsed := time.Since(start)
	fmt.Printf("Descarga completada en %v \n",elapsed)
}

