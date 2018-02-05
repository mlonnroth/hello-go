// Testing file operations

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println("Open file...")

	fd, err := os.Open("the_tempest.txt")
	if err != nil {
		log.Fatal(err)
	}

	//	buf := make(
	err = fd.Close()
	if err != nil {
		log.Fatal(err)
	}

}
