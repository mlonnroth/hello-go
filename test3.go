// Testing file operations

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println("Open file...")

	fd, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Close file...")
	err = fd.Close()
	if err != nil {
		log.Fatal(err)
	}

}
