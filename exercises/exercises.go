package exercises

import (
	"strconv"
)

var timeTable = [...]int{1, 2, 4, 8, 1, 2, 4, 8, 16, 32}
var indexTable = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

//统计几何个数
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
