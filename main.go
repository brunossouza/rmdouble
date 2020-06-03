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
var count bool

var totalDuplicatos = 0

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
				fmt.Print("\nDiretory:", "\t", diretory)
			}
			if recursive {
				listDir(diretory)
			}
		} else {
			var hashValue = md5sum(diretory)
			if verbose {
				fmt.Print("\nFile:", "\t", file.Name(), "\t")
			}
			_, duplicado := filesMap[hashValue]

			if duplicado {
				if verbose {
					fmt.Print("DUPLICADO\t")
				}
				if count {
					totalDuplicatos++
				}

				if delete {
					os.Remove(diretory)
					fmt.Print("DELETADO!!!")
				}
			} else {
				filesMap[hashValue] = diretory
			}
		}
	}
}

func init() {
	flag.StringVar(&path, "p", "./", "Path do diretório a ser verificado.")
	flag.BoolVar(&recursive, "r", false, "Recursive")
	flag.BoolVar(&delete, "d", false, "Delete")
	flag.BoolVar(&count, "c", false, "Contar duplicados")
	flag.BoolVar(&verbose, "v", false, "Verbose")
	flag.Parse()
}

func main() {
	listDir(path)
	fmt.Print("\n", "Total de arquivos duplicados:\t", totalDuplicatos, "\n\n")
}
