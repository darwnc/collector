package exercises

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestBinaryWatch(t *testing.T) {
	//1<<N
	fmt.Println("pow math", math.Pow(2, 4))
	fmt.Println("pow1<<N", 1<<4)
	fmt.Println("combination=", combination(1, 6))
	// result := make([][]string, 0)
	// fmt.Printf("return addr %#[1]v\n", len(&result[0]))
	// combinationDetail(num, 0, 3, &result)
	// fmt.Println(result)
	resultTime := binaryWatch(2)

	fmt.Println("acqure=", resultTime)

	// result0 := returnValueA(10)
	// fmt.Printf("return addr%p %#[1]v\n", result0)

	// result1 := make([]string, 0)
	// returnValueB(10, &result1)
	// fmt.Printf("result1 addr %p %#[1]v\n", &result1)
}
func returnValueA(b int) (all []string) {
	if b == 1 {
		all = append(all, strconv.Itoa(1))
		return
	}
	if all == nil {
		all = make([]string, 0)
	}
	all = append(all, strconv.Itoa(b))
	fmt.Printf("addr %p\n", &all)
	return append(all, returnValueA(b-1)...)
}

func returnValueB(a int, result *[]string) {
	if a == 1 {
		*result = append(*result, strconv.Itoa(1))
		return
	}
	*result = append(*result, strconv.Itoa(a))
	fmt.Printf("---addr %p\n", result)
	returnValueB(a-1, result)
}

func TestRemoveDigits(t *testing.T) {
	result := removeKdigits("1432219", 3)
	fmt.Println(result)
}
