package main

import "fmt"

func printCharacters(s string) {
	for _, char := range s {
		fmt.Print(string(char))
	}
}

func main() {
	input := "Aaslema , Baba!"
	printCharacters(input)
}
