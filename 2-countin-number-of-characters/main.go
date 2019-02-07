package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func count(name string) int {
	numOfCharacters := []byte(name)
	return len(numOfCharacters)
}

const inputQuestion string = "What is the input string? "

func main() {
	fmt.Printf(inputQuestion)
	var name string
	_, err := fmt.Scanln(&name)

	if err != nil && err == io.EOF {
		panic(io.ErrUnexpectedEOF)
	} else if err != nil && name == "" {
		fmt.Println("User must enter something into the program.")
		return
	}
	fmt.Printf("%s has %d characters.\n", name, count(name))

	fmt.Println("Another Version")

	newVersion()

}

// Write a new version of the program without using any
// variables.

func newVersion() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf(inputQuestion)

	for scanner.Scan() {
		if scanner.Text() == "\n" {
			break
		}
		if scanner.Text() == "" {
			fmt.Println("User must enter something into the program.")

			break
		}
		fmt.Printf("%s has %d characters.\n", scanner.Text(), count(scanner.Text()))
		break
	}

	if scanner.Err() != nil {
		panic(errors.New("Scanner Error"))
	}
}
