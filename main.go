package main

import (
	"fmt"
	"io/ioutil"

	//	"reflect"
	"strings"
)

func main() {
	// fmt.Println("Program to implement the commands.")
	// content := "This needs to go in a file.\n"

	// file, err := os.Create("./mylog.txt")
	// checkNilErr(err)

	// io.WriteString(file, content)
	// length, err := io.WriteString(file, content)
	// checkNilErr(err)
	// fmt.Println("length is: ", length)

	// defer file.Close()

	implementWC("./mylog.txt")
}

func wc(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	// whatever we read through 'ioutil'
	// is usually in the form of bytes.
	checkNilErr(err)

	words := string(databyte)
	char_count := len(words)
	lineList := strings.Split(words, "\n")
	line_count := len(lineList) - 1
	wordsList := strings.Split(words, " ")
	words_count := len(wordsList) + line_count - 1

	// fmt.Println(reflect.TypeOf(wordsList))
	// fmt.Println(wordsList)

	fmt.Println(line_count, "		", words_count, "		", char_count)
}

func implementWC(filename string) {
	wc(filename)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
