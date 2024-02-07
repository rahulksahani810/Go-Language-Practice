package main

import (
	"fmt"
	"unicode/utf8"
)
func printBytes2(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars2(s string) {
	fmt.Printf("Characters: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}
func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func main() {
	name := "Hello World"
	fmt.Printf("String: %s\n", name)
	printBytes(name)
	fmt.Printf("String: %s\n", name)
	printChars2(name)
	fmt.Printf("\n")
	printBytes2(name)
	fmt.Println("**********************")
	word1 := "Señor"
	fmt.Printf("String: %s\n", word1)
	fmt.Printf("Length: %d\n", utf8.RuneCountInString(word1))
	fmt.Printf("Number of bytes: %d\n", len(word1))

	fmt.Printf("\n")
	word2 := "Pets"
	fmt.Printf("String: %s\n", word2)
	fmt.Printf("Length: %d\n", utf8.RuneCountInString(word2))
	fmt.Printf("Number of bytes: %d\n", len(word2))
}