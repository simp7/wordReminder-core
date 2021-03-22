package main

import (
	"fmt"
	"github.com/simp7/wordReminder-core/voca/class"
	"github.com/simp7/wordReminder-core/voca/meaning"
	"github.com/simp7/wordReminder-core/voca/word"
	"github.com/simp7/wordReminder-core/voca/wordSet"
	"sort"
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
