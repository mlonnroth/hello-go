// input1.go
// Read keyboard input

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
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

	//Using keyboard
	fmt.Println("Read keystroke, q to quit")
	// disable input buffering (this shit doesn't work on Windows/Cygwin)
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		fmt.Println("I got the byte", b, "("+string(b)+")")
		if string(b) == "q" {
			break
		}
	}
}
