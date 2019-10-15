package main

import (
	"fmt"
	"sort"
	s "strings"
)

// define function for sort string
func sortVowels(sentence string) string {
	// lowercase user input and then make them an slice
	sLower := s.ToLower(sentence)
	arrInput := s.Split(sLower, "")

	// this var used to put the new vowels alphabet
	vowels := []string{}

	//this var for notyfiying where index to delete the vowels alphabet
	consonants := []string{}

	for _, alphabet := range arrInput {
		switch alphabet {
		case "a", "i", "u", "e", "o":
			vowels = append(vowels, alphabet)
		default:
			consonants = append(consonants, alphabet)
		}
	}

	// sort new array include vowels alphabet
	sort.Strings(vowels)
	// sort array input include consonant alphabet
	sort.Strings(consonants)
	// merged two array one is vowels another is consonant and the join them
	arrayMerged := append(vowels, consonants...)
	output := s.Join(arrayMerged, "")
	return "Hasil sorting: " + output
}

func main() {
	// implement the function
	input := ""
	fmt.Print("Masukkan Kata: ")
	fmt.Scanln(&input)
	fmt.Println(sortVowels(input))
}
