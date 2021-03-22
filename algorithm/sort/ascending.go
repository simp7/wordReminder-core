package sort

import (
	"github.com/simp7/wordReminder-core/voca"
	"sort"
)

func Ascending(set voca.WordSet) {
	sort.Sort(set)
}

func Descending(set voca.WordSet) {
	sort.Sort(sort.Reverse(set))
}
