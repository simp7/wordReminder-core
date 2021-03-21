package main

import (
	"fmt"
	"sort"
	"wordReminder-core/voca/class"
	"wordReminder-core/voca/meaning"
	"wordReminder-core/voca/word"
	"wordReminder-core/voca/wordSet"
)

func main() {
	a := word.New("apple", meaning.New(class.Noun, "사과"))
	z := word.New("zebra", meaning.New(class.Noun, "얼룩말"))
	g := word.New("go", meaning.New(class.Verb, "가다"), meaning.New(class.Noun, "Go 언어"))

	words := wordSet.New("hello", a, z, g)
	fmt.Printf("%+v\n", words)
	sort.Sort(words)
	fmt.Printf("%+v\n", words)
	fmt.Printf("%+v\n", sort.Reverse(words))

}
