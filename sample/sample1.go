package main

import (
	"fmt"
	"github.com/simp7/wordReminder-core/voca/standard"
	"sort"
)

func main() {
	a := standard.Word("apple", standard.Noun("사과"))
	z := standard.Word("zebra", standard.Noun("얼룩말"))
	g := standard.Word("go", standard.Verb("가다"), standard.Noun("Go 언어"))

	words := standard.Words("hello", a, z, g)
	fmt.Printf("%+v\n", words)
	sort.Sort(words)
	fmt.Printf("%+v\n", words)
	fmt.Printf("%+v\n", sort.Reverse(words))

}
