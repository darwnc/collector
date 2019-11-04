package exercises

import (
	"fmt"
	"sort"
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
	elementCount := make(map[string]int)
	// breakPoint := false
	element := make([]rune, 0)
	// formulaRune := []rune(formula)
	for k, v := range formula {
		// fmt.Println("k", k, "v=", v, ",string=", string(v))
		//遇到括号，数字，确定当前元素
		if rangeNumber(v) || v == rightBracket || v == leftBracket {
			if len(element) > 0 {
				elementCount[string(element)] = 0
			}
			element = make([]rune, 0)
			continue
		}
		if rangeaz(v) { //小写字母
			element = append(element, v)
			continue
		}
		// fmt.Println(k, "---", string(v))
		if rangeAZ(v) { //A~Z之间统计下一个大写字符看前面一个元素
			// breakPoint = true
			if len(element) > 0 { //包含有，结束，统计下一个
				//后面一个元素，不是数字则为1
				// if rangeNumber(formulaRune[k-1]) {
				// 	count, _ := strconv.Atoi(string(formulaRune[k-1]))
				// 	elementCount[string(element)] = count
				// } else {
				// }
				elementCount[string(element)] = 0
				element = make([]rune, 0)
			}
			element = append(element, v)
			if k == len(formula)-1 { //最后一个是大写字母的话，加入
				elementCount[string(element)] = 0
			}
		}
	}

	fmt.Println(elementCount)

	copyFormula := formula
	// leftTrubo := 1
	rightTurbo := 1
	for strings.Index(copyFormula, "(") != -1 {

		start := strings.Index(copyFormula, "(")
		end := strings.LastIndex(copyFormula, ")")
		// fmt.Println("copyFormula=", copyFormula[start+1:end])
		if start == 0 {
			if len(copyFormula) > end+1 {

			}
			// copyFormula[end]
		}
		// rightTurbo = 1
		fmt.Println("start=", copyFormula[0:start])

		countAtom(copyFormula[0:start], elementCount, rightTurbo)

		if len(copyFormula) >= end+2 {
			var err error
			var pars int
			pars, err = strconv.Atoi(copyFormula[end+1 : end+2])
			if err != nil {
				rightTurbo = 1
			} else {
				rightTurbo = pars * rightTurbo
			}
		}
		fmt.Println("trubo=", rightTurbo)
		copyFormula = copyFormula[start+1 : end]
		fmt.Println("after=", copyFormula)
	}
	countAtom(copyFormula, elementCount, rightTurbo)
	result := make([]string, 0)
	for k, v := range elementCount {
		if v == 1 {
			result = append(result, k)
		} else {
			result = append(result, k+strconv.Itoa(v))
		}

	}
	// fmt.Printf("result=%#v\n", stringSlice(result))
	sort.Sort(stringSlice(result))
	// result = sort.StringSlice(result)
	// fmt.Println(result)
	return strings.Join(result, "")
}

func split(formula string) {
	count := 0
	countBracket := 0
	outer := make([]string, 0)
	for k, v := range formula {

		if v == leftBracket { //寻找最近的一个最右括号，
			count++
			fmt.Println("count=", count)
			for i := 1; i+k < len(formula); i++ {
				if formula[i+k] == leftBracket { //(
					countBracket++
					fmt.Println("leftBracketIndex=", i+k)
				}
				if formula[i+k] == rightBracket {
					countBracket--
					fmt.Println("rightBracketIndex=", i+k)
				}

				if countBracket == -1 { //两个相邻的括号）>( 闭合
					fmt.Println(formula[k+1 : i+k])
					fmt.Println(formula[i+k : i+k])
					countBracket = 0
					break
				}
			}
		} else {
			outer = append(outer, string(v))
		}
	}
}

type stringSlice []string

func (ss stringSlice) Len() int {
	return len(ss)
}
func (ss stringSlice) Less(i, j int) bool {
	return ss[i][0] < ss[j][0]
}
func (ss stringSlice) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func countAtom(f string, ele map[string]int, turbo int) {
	// buff := []rune(f)
	// buff := bytes.Runes([]byte(f))
	for len(f) != 0 {
		for k := range ele {
			// b := []byte(k)[0]
			// index := bytes.IndexByte(buff, b)
			//k 表达为\x00**这样的字符，去掉前面的00
			// fmt.Printf("element=%#v len(k)=%v\n", k, len(k))
			fmt.Printf("%#v\n", f)
			index := strings.Index(f, k)
			// strings.Fields(s)
			// index := strings.IndexAny(f, k)
			if index == -1 {
				continue
			}
			if len(f) >= index+len(k)+1 {
				// fmt.Println("index=", index, "atom=", k, "ele=", f, "countAtom=", string(f[index]))
				n := 0
				// var err error
				i := 1
				for ; len(f) >= index+len(k)+i; i++ {
					if in, err := strconv.Atoi(f[index+len(k) : index+len(k)+i]); err == nil { //出错，部位nil,记录前一位数字
						n = in
					} else {
						break
					}
				}
				if n > 0 {
					ele[k] = ele[k] + n*turbo
					f = f[0:index] + f[index+len(k)+i-1:]
				} else {
					ele[k] = ele[k] + 1*turbo
					f = f[0:index] + f[index+len(k):]
				}
			} else {
				ele[k] = ele[k] + 1*turbo
			}
			if len(f) == 0 {
				break
			}
			// if index == 0 {
			// 	f = f[len(k):]
			// } else {
			// 	f = f[0:index] + f[index+len(k):]
			// }
		}
	}

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
