package exercises

import "fmt"

var hours = [...]float64{0, 1, 2, 3}
var mins = [...]float64{0, 1, 2, 3, 4, 5}
var timeTable = [...]int{1, 2, 4, 8, 1, 2, 4, 8, 16, 32}
var indexTable = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

//练习
//https://leetcode-cn.com/problems/binary-watch/
//找出indexTable给定指数的组合，去timeTable内的值即可得到，然后判断是否合适规则
//合并成一个从里面任意取第二部分的组合[1,2,4,8,1,2,4,8,16,32] [0,0,0,0,0,0,0,0,0,0]
func binaryWatch(num int) {
	for i := 0; i < num; i++ {

	}
}

//num包含个数
func combinationDetail(all [3]string, begin, num int, result *[][]string) {
	//总个数有num个 从第一位开始去，取到满足有num个
	// container := [count]string{}
	// *result = append(*result, all[begin])

	count := combination(num, len(all))
	fmt.Println("count=", count)
	for i := int64(0); i < count; i++ {
		var temp string

		fmt.Println("index=", temp)
	}
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
