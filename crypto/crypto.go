package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"
)

// des 秘钥长度为 8
func p5DesEncrypt(data, key string) string {
	block, err := des.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	buff := pcks5Padding([]byte(data), block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, []byte(key))
	encryptBuff := make([]byte, len(buff))
	blockMode.CryptBlocks(encryptBuff, buff)
	resultString := base64.StdEncoding.EncodeToString(encryptBuff)
	return resultString
}
func pcks5Padding(buff []byte, blockSize int) []byte {
	padding := blockSize - len(buff)%blockSize
	paddingBuff := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(buff, paddingBuff...)
}
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func p5DesDecrypt(data, key string) string {
	buff, base64Err := base64.StdEncoding.DecodeString(data)
	if base64Err != nil {
		fmt.Println(base64Err)
		return ""
	}
	block, cipherErr := des.NewCipher([]byte(key))
	if cipherErr != nil {
		fmt.Println(cipherErr)
		return ""
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(key))
	origin := make([]byte, len(buff))
	blockMode.CryptBlocks(origin, buff)
	return string(pkcs5UnPadding(origin))
}

func aesEncrypt(data, key string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}
	buff := []byte(data)
	encrytyBuff := pcks5Padding(buff, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, []byte(key))
	result := make([]byte, len(encrytyBuff))
	blockMode.CryptBlocks(result, encrytyBuff)
	return base64.StdEncoding.EncodeToString(result)
}

func aesDecrypt(data, key string) string {
	buff, encodErr := base64.StdEncoding.DecodeString(data)
	if encodErr != nil {
		fmt.Println(encodErr)
		return ""
	}
	block, cipherErr := aes.NewCipher([]byte(key))
	if cipherErr != nil {
		fmt.Println(cipherErr)
		return ""
	}

	blockMode := cipher.NewCBCDecrypter(block, []byte(key))
	result := make([]byte, len(buff))
	blockMode.CryptBlocks(result, buff)
	return string(pkcs5UnPadding(result))
}
