package exercises

import (
	"fmt"
	"strings"
)

// https://leetcode-cn.com/problems/number-of-matching-subsequences/
func numMatchingSubseq(s string, words []string) int {
	// total := word{string: s}
	// total.initField()
	// fmt.Println(total)
	count := 0
	tempS := new(string)
	for _, v := range words {
		success := true
		*tempS = s
		for _, wv := range v {
			sub := strings.SplitN(*tempS, string(wv), 2)
			if len(sub) == 2 {
				*tempS = sub[1]
				continue
			}
			success = false
			break

		}
		if success {
			count++
		}
	}
	return count
}

type word struct {
	string
	count map[string]int
}

func (w *word) initField() {
	if w.count == nil {
		w.count = make(map[string]int, 24)
	}
	for _, v := range alpha {
		if count := strings.Count(w.string, v); count > 0 {
			w.count[v] = count
		}
	}
}
func (w word) String() string {
	return fmt.Sprintln(w.count)
}
func (w word) countSingle(single string) (v int, ok bool) {
	v, ok = w.count[single]
	return
}

var alpha = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i",
	"j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v",
	"w", "x", "y", "z"}

type countString string

func (cs countString) remove(m rune) (result string, del bool) {
	once := false
	del = false
	result = strings.Map(func(r rune) rune {
		if m == r && !once { //只移除一次
			once = true
			del = true
			return -1
		}
		return r
	}, string(cs))
	return
}
