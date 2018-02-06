// input1.go
// Read keyboard input

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Using fmt Scanln
	fmt.Print("Enter text: ")
	var input string
	fmt.Scanln(&input)
	fmt.Printf("\"%s\"\n", input)

	// Using bufio scanner
	fmt.Println("using NewScanner(), type \"quit\" to end")
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		log.Println("line", s.Text())
		if s.Text() == "quit" {
			break
		}
	}
}
