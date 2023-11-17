package main

import (
	"bufio"
	"fmt"
	"os"
)

// interface definition
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

// MyString implements VowelsFinder
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {

	reader := bufio.NewReader(os.Stdin) // thos thing works as a Scnner object
	fmt.Println("Enter the String in which you want to find the Vowels")
	input,err := reader.ReadString('\n')
	if err!=nil {
		panic(err)}
		
	name := MyString(input)// type casting string to MyString 
	var v VowelsFinder
	v = name// possible since MyString implements VowelsFinder
	fmt.Printf("Vowels are %c", v.FindVowels())

}
