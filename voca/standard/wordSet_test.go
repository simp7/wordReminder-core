package standard

import (
	"github.com/simp7/wordReminder-core/voca"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

var dictionary map[string]voca.Word
var once = sync.Once{}

func wordOf(s string) voca.Word {
	once.Do(func() {
		dictionary = make(map[string]voca.Word)
		dictionary["apple"] = Word("apple", Noun("사과"))
		dictionary["go"] = Word("go", Noun("Go 언어"), Verb("가다"))
		dictionary["and"] = Word("and", Conjunction("그리고"))
	})
	return dictionary[s]
}

func TestSet_Len(t *testing.T) {

	scenario := []struct {
		desc   string
		set    voca.WordSet
		output int
	}{
		{"empty", Words("1"), 0},
		{"one", Words("2", wordOf("apple")), 1},
		{"many", Words("3", wordOf("apple"), wordOf("go"), wordOf("and")), 3},
	}

	for _, v := range scenario {
		assert.Equal(t, v.output, v.set.Len(), v.desc)
	}

}

func TestSet_Get(t *testing.T) {

	scenario := []struct {
		desc   string
		set    voca.WordSet
		idx    int
		output voca.Word
	}{
		{"get 3 from empty", Words("1"), 3, nil},
		{"get 0 from one-word set", Words("2", wordOf("apple")), 0, wordOf("apple")},
		{"get 1 from one-word set", Words("3", wordOf("apple")), 1, nil},
		{"get 1 from three-word set", Words("4", wordOf("apple"), wordOf("go"), wordOf("and")), 1, wordOf("go")},
		{"get 100 from three-word set", Words("5", wordOf("apple"), wordOf("go"), wordOf("and")), 100, nil},
	}

	for _, v := range scenario {
		assert.Equal(t, v.output, v.set.Get(v.idx), v.desc)
	}

}

func TestSet_Less(t *testing.T) {

	scenario := []struct {
		desc   string
		set    voca.WordSet
		idx1   int
		idx2   int
		output bool
	}{
		{"compare two from empty", Words("1"), 1, 2, false},
		{"compare self one-word set", Words("2", wordOf("apple")), 1, 1, false},
		{"compare 1, 2 from three-word set", Words("3", wordOf("apple"), wordOf("go"), wordOf("and")), 1, 2, false},
		{"compare 2, 1 from three-word set", Words("4", wordOf("apple"), wordOf("go"), wordOf("and")), 2, 1, true},
		{"compare over the index from three-word set", Words("5", wordOf("apple"), wordOf("go"), wordOf("and")), 100, 150, false},
	}

	for _, v := range scenario {
		assert.Equal(t, v.output, v.set.Less(v.idx1, v.idx2), v.desc)
	}

}

func TestSet_Swap(t *testing.T) {

	scenario := []struct {
		desc   string
		set    voca.WordSet
		idx1   int
		idx2   int
		output voca.WordSet
	}{
		{"compare two from empty", Words("1"), 1, 2, Words("1")},
		{"compare self one-word set", Words("2", wordOf("apple")), 1, 1, Words("2", wordOf("apple"))},
		{"compare 1, 2 from three-word set", Words("3", wordOf("apple"), wordOf("go"), wordOf("and")), 1, 2, Words("3", wordOf("apple"), wordOf("and"), wordOf("go"))},
		{"compare over the index from three-word set", Words("4", wordOf("apple"), wordOf("go"), wordOf("and")), 100, 150, Words("4", wordOf("apple"), wordOf("go"), wordOf("and"))},
	}

	for _, v := range scenario {
		v.set.Swap(v.idx1, v.idx2)
		assert.Equal(t, v.output, v.set, v.desc)
	}

}
