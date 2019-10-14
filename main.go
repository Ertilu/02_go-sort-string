package main

import (
	"fmt"
	"sort"
	s "strings"
)

// define function for sort string
func sortVowels(sentence string) string {
	vowels := []string{"a", "e", "i", "o", "u"}

	// lowercase user input and then make them an slice
	sLower := s.ToLower(sentence)
	arrInput := s.Split(sLower, "")

	// this var used to put the new vowels alphabet
	newArr := []string{}

	//this var for notyfiying where index to delete the vowels alphabet
	iVowels := []int{}

	// loop for add vowels to new aray
	for i := 0; i < len(arrInput); i++ {
		for v := 0; v < len(vowels); v++ {
			if arrInput[i] == vowels[v] {
				newArr = append(newArr, arrInput[i])
				iVowels = append(iVowels, i)
			}
		}
	}

	// loop for delete vowels alphabet so the array input includes all of consonant alpahabet
	for j := 0; j < len(iVowels); j++ {
		copy(arrInput[j:], arrInput[j+1:])
		arrInput[len(arrInput)-1] = ""
		arrInput = arrInput[:len(arrInput)-1]
	}

	// sort new array include vowels alphabet
	sort.Strings(newArr)
	// sort array input include consonant alphabet
	sort.Strings(arrInput)
	// merged two array one is vowels another is consonant and the join them
	arrayMerged := append(newArr, arrInput...)
	output := s.Join(arrayMerged, "")
	return output
}

func main() {
	// implement the function
	fmt.Println(sortVowels("osama"))
}
