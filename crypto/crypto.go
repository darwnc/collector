package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

//IV 默认aes iv
var IV = []byte{0x59, 0x54, 0x42, 0x53, 0x42, 0x43, 0x42, 0x41, 0x4e, 0x4b, 0x5a, 0x58, 0x4b, 0x45, 0x59, 0x53}

//AESEncrypt ase加密
func AESEncrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	buff := pcks5Padding(data, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, IV)
	result := make([]byte, len(buff))
	blockMode.CryptBlocks(result, buff)
	return result, nil
}

//AESDecrypt ase 解密
func AESDecrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, IV)
	if data == nil || len(data) == 0 {
		return nil, errors.New("AESDecrypt data is nil or len=0")
	}
	buff := make([]byte, len(data))
	blockMode.CryptBlocks(buff, data)
	return pkcs5UnPadding(buff), nil
}

//RSAEncrypt rsa加密
func RSAEncrypt(data []byte) ([]byte, error) {
	publicKey, err := ioutil.ReadFile("./pub_key.cer")
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, data)
}

//RSADecrypt 解密
func RSADecrypt(data []byte) ([]byte, error) {
	privateKey, err := ioutil.ReadFile("./pri_key.cer")
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// dataBuff, _ := base64.StdEncoding.DecodeString(data)
	return rsa.DecryptPKCS1v15(rand.Reader, priv, data)
}
func verify() {
	// publicKey, err := ioutil.ReadFile("./pub_key.cer")
	// if err != nil {
	// 	return
	// }
	// block, _ := pem.Decode(publicKey)
	// if block == nil {
	// 	return
	// }
	// pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	// h := sha1.New()
	// h.Write(src)
	// hashed := h.Sum(nil)
	// return rsa.VerifyPKCS1v15(this.publicKey, hash, hashed, sign)
}

func genPemFile() {
	priKey, _ := os.Create("pri_key.pem")
	pubKey, _ := os.Create("pub_key.pem")
	privateKey, keyGenErr := rsa.GenerateKey(rand.Reader, 1024)
	if keyGenErr != nil {
		fmt.Println(keyGenErr)
		return
	}
	privateKeybuff := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{Type: "privatekey", Bytes: privateKeybuff}
	encodeErr := pem.Encode(priKey, block)
	if encodeErr != nil {
		fmt.Println(encodeErr)
		return
	}

	publicKey := &privateKey.PublicKey
	pubKeyBuff, pubKeyGenErr := x509.MarshalPKIXPublicKey(publicKey)
	if pubKeyGenErr != nil {
		fmt.Println(pubKeyGenErr)
		return
		// logger.Warning(pubKeyGenErr)
	}
	pubBlock := &pem.Block{Type: "publickey", Bytes: pubKeyBuff}
	pubEncodeErr := pem.Encode(pubKey, pubBlock)
	if pubEncodeErr != nil {
		fmt.Println(pubEncodeErr)
		// logger.Trace(pubEncodeErr)
		return
	}
}

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
