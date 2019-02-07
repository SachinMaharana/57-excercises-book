package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func concatenation(name string) string {
	return fmt.Sprintf("Hello, %s, nice to meet you!\n", name)
}

const greetingQuestion string = "What is your name? "

func main() {
	fmt.Printf(greetingQuestion)
	var name string
	_, err := fmt.Scanln(&name)

	if err != nil {
		panic(err)
	}

	output := concatenation(name)

	fmt.Println(output)

	fmt.Println("Another Version")

	newVersion()

}

// Write a new version of the program without using any
// variables.

func newVersion() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf(greetingQuestion)

	for scanner.Scan() {
		if scanner.Text() == "\n" {
			break
		}
		fmt.Println(concatenation(scanner.Text()))
		break
	}

	if scanner.Err() != nil {
		panic(errors.New("Scanner Error"))
	}
}
