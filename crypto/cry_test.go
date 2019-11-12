package crypto

import (
	"fmt"
	"testing"
)

func TestDes(t *testing.T) {
	desKey := "12345678"
	result := p5DesEncrypt("123456", desKey)
	fmt.Println("result", result)
	decrypt := p5DesDecrypt(result, desKey)
	fmt.Println("result", decrypt)
	aesKey := "1234567812345678"
	aesResult := aesEncrypt("1234569", aesKey)
	fmt.Println("result", aesResult)
	fmt.Println("result", aesDecrypt(aesResult, aesKey))
}
