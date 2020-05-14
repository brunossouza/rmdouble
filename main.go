package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var path string
var recursive bool
var verbose bool
var delete bool

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func verifyDirPath(stringPath string) (stringResult string) {

	if strings.HasSuffix(stringPath, "/") {
		stringResult = stringPath
	} else {
		stringResult = stringPath + "/"
	}

	return
}

func md5sum(filePath string) (result string) {
	file, err := os.Open(filePath)
	checkError(err)
	defer file.Close()

	hash := md5.New()
	_, err = io.Copy(hash, file)
	checkError(err)

	result = hex.EncodeToString(hash.Sum(nil))
	return
}

func listDir(dir string) {
	files, err := ioutil.ReadDir(dir)
	checkError(err)

	var filesHash [len(files)]string

	for indice, file := range files {
		diretory := verifyDirPath(dir) + file.Name()

		if file.IsDir() {
			if verbose {
				fmt.Println("Diretory:", "\t", diretory)
			}
			if recursive {
				listDir(diretory)
			}
		} else {
			var hashValue = md5sum(diretory)
			if verbose {
				fmt.Println("File:", "\t", file.Name(), "\t")
			}
			filesHash[indice] = hashValue
			fmt.Println(len(filesHash))
		}
	}
}

func init() {
	flag.StringVar(&path, "p", "./", "Path do diretório a ser verificado.")
	flag.BoolVar(&recursive, "r", false, "Recursive")
	flag.BoolVar(&delete, "d", true, "Delete")
	flag.BoolVar(&verbose, "v", false, "Verbose")
	flag.Parse()
}

func main() {
	fmt.Println(recursive, path)
	listDir(path)
}
