package exter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/darwnc/collector/crypto"

	"github.com/gin-gonic/gin"
)

//cryptoJSON writes data with custom ContentType.
// Render(http.ResponseWriter) error
// WriteContentType(w http.ResponseWriter)WriteContentType writes custom ContentType.
type cryptoJSON struct {
	data interface{}
}

func (json cryptoJSON) Render(writer http.ResponseWriter) error {
	return writeCryptoJSON(writer, json.data)
}

func (json cryptoJSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

var jsonContentType = []string{"application/json; charset=utf-8"}

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	//写入新的header
	header["Context-Track"] = []string{"header context"}
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}
func writeCryptoJSON(w http.ResponseWriter, obj interface{}) error {
	writeContentType(w, jsonContentType)
	var buff []byte
	var err error
	switch obj.(type) {
	case string:
		w.Write([]byte(obj.(string)))
		return nil
		// buff, _ = crypto.StdDeBase64(obj.(string))
	default:
		buff, err = json.Marshal(obj)
		// crypto.  加密相关操作
		if err != nil {
			return err
		}
	}
	// var encrypt []byte
	// encrypt, err = crypto.AESEncrypt(buff, crypto.IV)
	// if err != nil {
	// 	return err
	// }
	_, err = w.Write(buff)
	return err
}

// type Binding interface {
// 	Name() string
// 	Bind(*http.Request, interface{}) error
// }

type decryptBody struct{}

func (d decryptBody) Name() string {
	return "decryptBody"
}
func (d decryptBody) Bind(reqeust *http.Request, data interface{}) error {
	buff, err := ioutil.ReadAll(reqeust.Body)
	if err != nil {
		return err
	}
	// fmt.Println(string(buff))
	base64Decode := make([]byte, len(buff))
	n, _ := crypto.StdDeByteBase64(base64Decode, buff)
	// base64Decode, _ := crypto.StdDeBase64(string(buff))
	// crypto.StdDeByteBase64(base64Decode, buff)
	dbody, decErr := crypto.AESDecrypt(base64Decode[:n], crypto.IV)
	if decErr != nil {
		return decErr
	}
	if jsonErr := json.Unmarshal(dbody, data); jsonErr != nil {
		return jsonErr
	}
	// body := data.(*string)
	// *body = string(dbody)
	return nil
}

// type Cast interface {
// 	Cast()
// }

//Process 处理报问题
//返回interface{} 返回报文
type Process func() interface{}

//Wrap 赋值data Process接收返回的结构体
// 请求结构体{header:{},payload:{}} 与其绑定失败则不执行param
// 返回结构体 BaseResp like {data{},code,msg}
func Wrap(data interface{}, param Process) gin.HandlerFunc {
	return func(c *gin.Context) {
		//header相关操作
		// ShouldBindWith(obj interface{}, b binding.Binding)
		// var body testRequest
		if err := c.ShouldBindWith(data, decryptBody{}); err != nil {
			fmt.Println(err.Error())
			c.Render(http.StatusOK, cryptoJSON{err.Error()})
			return
			//解析错误
		}
		// buff, _ := ioutil.ReadAll(c.Request.Body)
		// c.Bind(obj)
		respData := param()
		c.Render(http.StatusOK, cryptoJSON{respData})
	}
}
