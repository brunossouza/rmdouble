package main

import (
	"container/list"
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

type object struct {
	name    string
	path    string
	hashSum string
	count   int
}

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

func listDir(dir string) (values *list.List) {
	files, err := ioutil.ReadDir(dir)
	checkError(err)

	values = list.New()

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
			if verbose {
				fmt.Println("File:", "\t", file.Name(), "\t", md5sum(diretory))
			}
			obj := object{name: file.Name(), path: diretory, hashSum: md5sum(diretory), count: 1}
			values.PushBack(obj)
		}
	}

	return values
}

func init() {
	flag.StringVar(&path, "p", "./", "Path do diretório a ser verificado.")
	flag.BoolVar(&recursive, "r", false, "Recursive")
	flag.BoolVar(&verbose, "v", false, "Verbose")
	flag.Parse()
}

func main() {
	fmt.Println(recursive, path)
	_ = listDir(path)

	// Loop over container list.
	// for temp := files.Front(); temp != nil; temp = temp.Next() {
	// 	fmt.Println(temp.Value)
	// }
}
