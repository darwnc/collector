package exercises

import (
	"fmt"
	"strconv"
	"strings"
)

var timeTable = [...]int{1, 2, 4, 8, 1, 2, 4, 8, 16, 32}
var indexTable = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

//统计个数
var count = 0

//练习
//https://leetcode-cn.com/problems/binary-watch/
//找出indexTable给定指数的组合，去timeTable内的值即可得到，然后判断是否合适规则
//合并成一个从里面任意取第二部分的组合[1,2,4,8,1,2,4,8,16,32] [0,0,0,0,0,0,0,0,0,0]
func binaryWatch(num int) (resultTime []string) {
	reslutContainer := make(map[int][]int)
	for i := 1; i <= len(indexTable); i++ {
		result := make([]int, 0)
		acquire(indexTable[:], i, &result, reslutContainer)
	}
	for _, v := range reslutContainer {
		if len(v) == num { //符合条件
			// fmt.Println(v)
			mins := 0 //<=59
			hour := 0 //<=11
			for _, v := range v {
				if v > 3 { //超过三，表示分钟
					mins += timeTable[v]
				} else { //小时
					hour += timeTable[v]
				}

			}
			if mins > 59 || hour > 11 { //不符合规范

			} else {
				hourString := strconv.Itoa(hour)
				minsString := strconv.Itoa(mins)
				if len(minsString) == 1 { //加零
					minsString = "0" + minsString
				}
				resultTime = append(resultTime, hourString+":"+minsString)
			}

			mins = 0
			hour = 0

		}
	}
	return
}

//num包含个数

func acquire(all []int, num int, result *[]int, reslutContainer map[int][]int) {
	if *result == nil {
		*result = make([]int, 0)
	}

	if num == 0 {
		reslutContainer[count] = append(reslutContainer[count], (*result)...)
		// fmt.Println("acquire=", result)
		count++
		return
	}

	if len(all) == 0 {
		return
	}
	// fmt.Println("all=", all)
	*result = append(*result, all[0])
	acquire(all[1:], num-1, result, reslutContainer)

	// fmt.Println("result befor=", result)
	// if len(*result) >= 1 {
	// 	*result = (*result)[1:]
	// }
	// fmt.Println("result after=", result)
	*result = (*result)[:len(*result)-1]
	acquire(all[1:], num, result, reslutContainer)
}

//计算排列的个数 n取出个数 m基数 n<=m
func combination(n, m int) (com int64) {
	com = 1
	if n == 0 || m == 0 {
		//
		return
	}
	resultM := factorial(int64(m))
	resultN := factorial(int64(n))
	resultDiv := factorial(int64(m - n))
	com = resultM / (resultN * resultDiv)
	return
}

//num>0
func factorial(num int64) (result int64) {
	result = 1
	for i := int64(1); i <= num; i++ {
		result = result * int64(i)
	}
	return
}

// 给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。
func removeKdigits(num string, k int) (result string) {
	if len(num) < k {
		return result
	}
	if len(num) == k {
		result = "0"
		return
	}
	// //剩余数字位数
	// resultLen = len(num) - k

	//选取前k+1位 移除最大的k个数字，剩下的即位最大的首字母位0去掉
	// remove := num[0:k]
	index := []int{}
	for key := range num {
		index = append(index, key)
	}
	reslutMap := make(map[int][]int)
	for i := 1; i <= len(index); i++ {
		result := make([]int, 0)
		acquire(index, i, &result, reslutMap)
	}
	allSatisfy := []string{}
	for _, v := range reslutMap {
		if len(v) == len(num)-k { //符合条件
			//移除相应的位数
			tempResult := ""
			for _, v := range v {
				tempResult = tempResult + string(num[v])
			}
			allSatisfy = append(allSatisfy, tempResult)
		}
	}

	resultNumber, _ := strconv.Atoi(num)

	for _, v := range allSatisfy {
		//移除头部的0
		resultTirm := strings.TrimLeftFunc(v, func(r rune) bool {
			if r == rune('0') {
				return true
			}
			return false
		})
		resultTirmNum, _ := strconv.Atoi(resultTirm)
		//获取最小值
		if resultTirmNum < resultNumber {
			resultNumber = resultTirmNum
		}
	}

	// acqureOne(num[0:k+1], &result)
	// // fmt.Println("removeKdigits1", result)
	// result = result + num[k+1:]
	// // fmt.Println("removeKdigits2", result)
	// //去掉头部的0
	// // result = strings.TrimPrefix(result, "0")
	// result = strings.TrimLeftFunc(result, func(r rune) bool {
	// 	if r == rune('0') {
	// 		return true
	// 	}
	// 	return false
	// })
	return strconv.Itoa(resultNumber)
}
func acqureOne(num string, result *string) {
	if result == nil {
		result = new(string)
	}
	if len(num) == 1 { //最后一位，算完成
		*result = num
		fmt.Printf("%#v\n", *result)
		return
	}
	bigest := 0
	for _, v := range num {
		// item := strconv.QuoteRune(v)
		if value, _ := strconv.Atoi(string(v)); value >= bigest {
			bigest = value
		}
	}
	// Replace(s, old, new string, n int) string
	isDel := false
	num = strings.Map(func(r rune) rune {
		if value, _ := strconv.Atoi(string(r)); value == bigest && !isDel { //只删一次
			isDel = true
			return -1
		}
		return r
	}, num)
	// fmt.Println("string=", num)
	// *result = num
	acqureOne(num, result)
}
