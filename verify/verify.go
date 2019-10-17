package verify

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//Verifys 验证模块
var verifys []gin.HandlerFunc

//SessionKey session的键
const SessionKey = "SessionID"

//User 用户模块
type User struct {
	Name string `form:"userName" json:"userName" xml:"userName"  binding:"required"`
	PWD  string `form:"passwrod" json:"passwrod" xml:"passwrod" binding:"required"`
}

func (user User) empty() bool {
	return len(user.Name) == 0 || len(user.PWD) == 0
}

func verifyLogin(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if user.empty() {
		c.JSON(http.StatusOK, gin.H{"empty": "user name empty and password empty"})
		return
	}
	sess := sessions.Default(c)
	var hexSessValue string
	sessValue := sess.Get(SessionKey)
	if sessValue == nil {
		//赋值
		buff := make([]byte, 16)
		if _, err := rand.Read(buff); err != nil {
			c.JSON(http.StatusOK, gin.H{"session err": err.Error()})
			return
		}
		hexSessValue = hex.EncodeToString(buff)
	}
	hexSessValue = sessValue.(string)
	sess.Set(SessionKey, hexSessValue)
	sess.Set("user", &user)
	sess.Save()
	c.JSON(http.StatusOK, gin.H{"user": user, "session": hexSessValue})
}

//GetVerify 获取验证模块
func GetVerify() []gin.HandlerFunc {
	return []gin.HandlerFunc{verifyLogin}
}
