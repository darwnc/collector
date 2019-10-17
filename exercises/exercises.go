package exercises

import (
	"math"
)

var hours = [...]float64{0, 1, 2, 3}
var mins = [...]float64{0, 1, 2, 3, 4, 5}

//练习  n分布
//n>6两个部分都有hours mins都有 n<=6可能有两个部分有，一个部分可能有
func binaryWatch(num int) {

	for i := 0; i < num; i++ {

	}
}

//遍历每一个可能的位置位数n 代表个数，两个
func hoursCount() {
	for _, value := range hours {
		math.Pow(2, value)
	}
}
func minsCount() {
	for _, value := range hours {
		math.Pow(2, value)
	}
}
