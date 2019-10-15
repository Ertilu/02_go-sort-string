package main

import (
	"testing"
)

func TestSort(t *testing.T) {
	// test is the input is nil
	t.Run("Nil input", func(t *testing.T) {
		input := ""
		sort := sortVowels(input)
		expectedResult := "Hasil sorting: "
		if sort != expectedResult {
			t.Fatalf("return must be %s instead %s", expectedResult, sort)
		}
	})

	// test is the are the input of string
	t.Run("Input osama string", func(t *testing.T) {
		input := "osama"
		sort := sortVowels(input)
		expectedResult := "Hasil sorting: aaoms"
		if sort != expectedResult {
			t.Fatalf("return must be %s instead %s", expectedResult, sort)
		}
	})

	// call the main function so the coverage will 100%
	main()
}

func BenchmarkSortVowels(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := "aiueo"
		sortVowels(input)
	}
}
