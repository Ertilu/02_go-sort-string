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
}
