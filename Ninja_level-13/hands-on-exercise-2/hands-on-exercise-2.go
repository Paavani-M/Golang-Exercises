package main

import (
	"fmt"

	"desktop/Ninja_level-13/hands-on-exercise-2/quote"
	"desktop/Ninja_level-13/hands-on-exercise-2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
