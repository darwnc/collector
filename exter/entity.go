package exter

// type Response struct {
// 	BaseResp
// 	Data interface{}
// }

//Resp 基础报文实体
type Resp struct {
	//返回码
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	T       int64       `json:"t"`
	Service string      `json:"service"`
	Data    interface{} `json:"data"`
}

//TestRequest 测试
type TestRequest struct {
	Header  Header      `json:"header"`
	Payload BasePayload `json:"payload"`
}

// Header token version osType servce
type Header struct {
	AppToken string `json:"appToken"`
	Version  string `json:"version"`
	OsType   string `json:"osType"`
	Service  string `json:"service"`
	T        int64  `json:"t"`
}

//BasePayload 请求实体
type BasePayload struct {
	AppToken string `json:"appToken"`
}
