package crypto

import (
	"encoding/base64"
)

var (
	//StdBase64 base64标准流bytes to string
	StdBase64 = base64.RawStdEncoding.EncodeToString
	//StdDeBase64 string to byte
	StdDeBase64 = base64.RawStdEncoding.DecodeString
	//StdDeByteBase64 byte to byte decode to base64
	StdDeByteBase64 = base64.RawStdEncoding.Decode
	//StdByteBase64 byte to byte encode to base64
	StdByteBase64 = base64.RawStdEncoding.Encode
)
