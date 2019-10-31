package exercises

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/
func findSubstring(s string, words []string) []int {
	result := []int{}
	for i := 0; i < len(words); i++ {
		if !strings.Contains(s, words[i]) { //不包含，没有
			return result
		}
	}
	//确认包含，确认顺序，从头开始遍历s,符合match中任何一个字符串的 ，继续拼组合，否则不匹配
	// utf8.RuneCountInString(s)
	// wordsCopy := make([]string, len(words))
	// copy(wordsCopy, words)
	// fmt.Println("copy=", wordsCopy)
	matchString := strings.Join(words, "")
	matchStringLen := utf8.RuneCountInString(matchString)
	r := []rune(s)
	count := 0
	for k, v := range r {
		if len(r)-k < matchStringLen {
			break
		}
		// fmt.Println(k, "--", string(v))
		//判断是否匹配match中任何一个中的第一个，有则继续，无责break跳出
		// value := string(v)

		// wordsCopy := make([]string, len(words))
		// wordsCopy = wordsCopy[0:len(words)]
		wordsCopy := make([]string, len(words))
		copy(wordsCopy, words)
		// fmt.Println(wordsCopy)
		moveTo := 0
		for {
			if k+moveTo >= len(r) || len(r)-k < matchStringLen { //剩余的长度不够真个字符串的长度，则不符合
				break
			}
			count++
			// value = string(r[k+moveTo])
			v = r[k+moveTo]
			// fmt.Println("value=", value)
			// fmt.Println("word=", wordsCopy)
			// fmt.Println("moveTo=", moveTo)
			have := false
			for i := 0; i < len(wordsCopy); i++ {
				wordRune := []rune(wordsCopy[i])
				if len(wordRune)+k+moveTo > len(r) {
					break
				}
				// fmt.Println("length==", string(wordRune[0]))
				// fmt.Println("wordsCopy==", wordsCopy[i])
				if wordRune[0] == v { //第一个字符相等 跳过[i]字符验证下一组
					wordLen := len(wordRune)
					if string(r[k+moveTo:wordLen+k+moveTo]) == wordsCopy[i] { //符合，下一个match的item
						// 删除当前item进行下一次迭代 ，下一次匹配的是当前s[k:len(words[i])+k+1]的值
						moveTo += wordLen
						wordsCopy = append(wordsCopy[:i], wordsCopy[i+1:]...)
						// fmt.Println("substring s=", s[k:len(words[i])+k])
						// fmt.Println("match=", words[i])
						have = true
						break
					}
				}
			}
			//如果wordsCopy无变化，则当前不匹配，跳出循环
			if len(wordsCopy) == len(words) || !have {
				wordsCopy = nil
				break
			}
			if len(wordsCopy) == 0 { //全符合
				// 记录index
				result = append(result, k)
				break
			}
		}
		count++
	}
	fmt.Println("count=", count)
	return result
}

func match(source []string, first string) bool {
	fullString := strings.Join(source, "")
	return strings.Contains(fullString, first)
}
