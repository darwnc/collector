package exercises

import (
	"fmt"
	"strconv"
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

// https://leetcode-cn.com/problems/number-of-atoms/
// formula = "K4(ON(SO3)2)2"
// 输出: "K4N2O14S4"
// 解释:
// 原子的数量是 {'K': 4, 'N': 2, 'O': 14, 'S': 4}。
func countOfAtoms(formula string) string {
	//括号、分割大写字符分割
	str := strings.SplitAfterN(formula, "(", 2)
	fmt.Println(str)
	inside := strings.SplitAfterN(str[1], ")", 2)
	fmt.Println(inside)
	elementCount := make(map[string]int)
	// breakPoint := false
	element := make([]rune, 0)
	formulaRune := []rune(formula)
	for k, v := range formula {
		fmt.Println("k", k, "v=", v, ",string=", string(v))
		if rangeaz(v) { //小写字母
			element = append(element, v)
			continue
		}
		if rangeAZ(v) { //A~Z之间统计下一个大写字符看前面一个元素
			// breakPoint = true
			if len(element) > 0 {
				//后面一个元素，不是数字则为1
				if rangeNumber(formulaRune[k-1]) {
					count, _ := strconv.Atoi(string(formulaRune[k-1]))
					elementCount[string(element)] = count
				} else {
					elementCount[string(element)] = 1
				}
			}
			element = make([]rune, 1)
			element = append(element, v)
		}
	}
	for k := range elementCount {
		index := strings.IndexAny(formula, k)
		fmt.Println(index)
		if len(formulaRune) > index+len(k)-1 {
			if rangeNumber(formulaRune[index+len(k)-1]) {
				n, _ := strconv.Atoi(string(formulaRune[index+len(k)-1]))
				elementCount[k] = n
			}
		}
	}
	fmt.Println(elementCount)
	return ""
}
func rangeaz(r rune) bool {
	return r >= minLowLetter && r <= maxLowLetter
}
func rangeAZ(r rune) bool {
	return r >= minLetter && r <= maxLetter
}
func rangeNumber(r rune) bool {
	return r >= minNumber && r <= maxNumber
}

const (
	maxLetter    = 'Z'
	minLetter    = 'A'
	maxLowLetter = 'z'
	minLowLetter = 'a'
	leftBracket  = '('
	rightBracket = ')'
	maxNumber    = '9'
	minNumber    = '1'
)
