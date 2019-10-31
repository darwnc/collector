package exercises

import (
	"fmt"
	"strings"
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
	r := []rune(s)
	count := 0
	for k, v := range r {
		// fmt.Println(k, "--", string(v))
		//判断是否匹配match中任何一个中的第一个，有则继续，无责break跳出
		value := string(v)

		wordsCopy := make([]string, len(words))
		copy(wordsCopy, words)
		moveTo := 0
		for {

			if k+moveTo >= len(r) || len(r)-k < len(matchString) { //剩余的长度不够真个字符串的长度，则不符合
				break
			}
			count++
			value = string(r[k+moveTo])
			// fmt.Println("value=", value)
			// fmt.Println("word=", wordsCopy)
			// fmt.Println("moveTo=", moveTo)
			have := false
			for i := 0; i < len(wordsCopy); i++ {
				if len(words[i])+k+moveTo > len(r) {
					break
				}
				// fmt.Println("length==")
				// fmt.Println("wordsCopy==", wordsCopy[i])
				if wordsCopy[i][0:1] == value { //第一个字符相等 跳过[i]字符验证下一组
					if s[k+moveTo:len(wordsCopy[i])+k+moveTo] == wordsCopy[i] { //符合，下一个match的item
						// 删除当前item进行下一次迭代 ，下一次匹配的是当前s[k:len(words[i])+k+1]的值
						moveTo += len(wordsCopy[i])
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
	}
	fmt.Println("count=", count)
	return result
}

func match(source []string, first string) bool {
	fullString := strings.Join(source, "")
	return strings.Contains(fullString, first)
}
