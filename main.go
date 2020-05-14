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

	var filesMap map[string]string
	filesMap = make(map[string]string)

	for _, file := range files {
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
				fmt.Print("File:", "\t", file.Name(), "\t")
			}
			_, duplicado := filesMap[hashValue]

			if duplicado {
				fmt.Print("DUPLICADO\t")
				if delete {
					os.Remove(diretory)
					fmt.Println("DELETADO!!!")
				}
			} else {
				filesMap[hashValue] = diretory
			}
			fmt.Println("")
		}
	}
}

func init() {
	flag.StringVar(&path, "p", "./", "Path do diretório a ser verificado.")
	flag.BoolVar(&recursive, "r", false, "Recursive")
	flag.BoolVar(&delete, "d", false, "Delete")
	flag.BoolVar(&verbose, "v", false, "Verbose")
	flag.Parse()
}

func main() {
	fmt.Println(recursive, path)
	listDir(path)
}
