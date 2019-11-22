package crypto

import (
	"fmt"
	"testing"
)

func TestDes(t *testing.T) {
	// desKey := "12345678"
	// result := p5DesEncrypt("123456", desKey)
	// fmt.Println("result", result)
	// decrypt := p5DesDecrypt(result, desKey)
	// fmt.Println("result", decrypt)
	// aesKey := "1234567812345678"
	// aesResult := aesEncrypt("1234569", aesKey)
	// fmt.Println("result", aesResult)
	// fmt.Println("result", aesDecrypt(aesResult, aesKey))
	var data = `{"header":{"_t":0,"service":"/v1/platform/home/index.do"},"payload":{"_t":1533621716093,"version":""}}`
	// var key = "J5wn8XWnQxXzTbxj"
	// iVEITNcxx9eOec6vBARm8w==
	// key := genKey()
	// fmt.Println("key=", base64.StdEncoding.EncodeToString(key))
	encrypt, _ := AESEncrypt([]byte(data), IV)
	base64Data := StdBase64(encrypt)
	fmt.Println("encrypt=", string(base64Data))
	decrypt, _ := AESDecrypt(encrypt, IV)
	fmt.Println(string(decrypt))

}
func TestRsa(t *testing.T) {
	var data = `{"header":{"_t":0,"service":"/v1/platform/home/index.do"},"payload":{"_t":1533621716093,"version":""}}`
	rsaencry, _ := RSAEncrypt([]byte(data))
	fmt.Println("rsa=\n", StdBase64(rsaencry))
	rsaDecry, _ := RSADecrypt(rsaencry)
	fmt.Println("decode rsa=\n", string(rsaDecry))
}
func TestGenPemFile(t *testing.T) {
	genPemFile()
}
