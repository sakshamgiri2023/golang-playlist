package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("files in golang")
	content := "this needs to go in a file -sakshamhgiri500@gmail.com"
	file, err := os.Create("./mylocalfile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length is:", length)
	defer file.Close()
	readFile("mylocalfile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("text data inside the file is \n", databyte)

}
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
