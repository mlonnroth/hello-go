// file3.go
// Copy file using ioutil package

package main

import (
	"io/ioutil"
)

func main() {
	// read the whole file at once
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// write the whole body at once
	err = ioutil.WriteFile("output.txt", b, 0644)
	if err != nil {
		panic(err)
	}
}
