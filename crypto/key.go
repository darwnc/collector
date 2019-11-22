package crypto

import (
	"crypto/rand"
)

//16位的key
func genKey() []byte {
	buff := make([]byte, 16)
	rand.Read(buff)
	return buff
}
