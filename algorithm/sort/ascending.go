package sort

import (
	"sort"
	"wordReminder-core/voca"
)

func Ascending(set voca.WordSet) {
	sort.Sort(set)
}

func Descending(set voca.WordSet) {
	sort.Sort(sort.Reverse(set))
}