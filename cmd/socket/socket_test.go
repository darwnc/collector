package main

import (
	"fmt"
	"hash/crc32"
	"strconv"
	"testing"
)

func TestSocket(t *testing.T) {
	crc := crc32.ChecksumIEEE([]byte("test"))
	// fmt.Println("crc32", crc/64, crc>>8)
	fmt.Println(strconv.FormatUint(uint64(crc), 2))
	// //验证crc 分割成4个8byte数组验证
	prix := a(crc)
	bytes := [...]byte{prix(), prix(), prix(), prix()}
	fmt.Println(bytes)
	fmt.Println("-->", strconv.FormatUint(uint64(bytes[0]), 2))
	fmt.Println("-->", strconv.FormatUint(uint64(bytes[1]), 2))
	fmt.Println("-->", strconv.FormatUint(uint64(bytes[2]), 2))
	fmt.Println("-->", strconv.FormatUint(uint64(bytes[3]), 2))
	// var result uint32
	// for i := 0; i < 4; i++ {
	// 	result += uint32(bytes[i]) << (i * 8)
	// }
	// // result := uint32(bytes[3])<<24 + uint32(bytes[2])<<16 + uint32(bytes[1])<<8 + uint32(bytes[0])
	// fmt.Println(strconv.FormatUint(uint64(result), 2))
	// fmt.Println("crc32", result)
	// fmt.Println(strconv.FormatUint(uint64(crc>>24), 2))      //前8位
	// fmt.Println(strconv.FormatUint(uint64(crc>>16&0xFF), 2)) //上面的后8位
	// crc16 := uint16((crc>>24))*256 + uint16((crc>>16)&0xFF)
	// result := (crc >> 16) & 0xFFFF
	// fmt.Println(strconv.FormatUint(uint64(crc16), 2))
	// fmt.Println(strconv.FormatUint(uint64(result), 2))
	// fmt.Println(( >> 16) & 0xFFFF)
	// dial()
	// <-sign

}

func a(p uint32) func() byte {
	i := 0
	return func() byte {
		b := (p >> (i * 8)) & 0xFF
		i++
		return byte(b)
	}
}
