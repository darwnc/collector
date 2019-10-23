package verify

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//User 用户信息
type User struct {
	Name     string `form:"userName" json:"userName" xml:"userName"  binding:"required"`
	Passwrod string `form:"passwrod" json:"passwrod" xml:"passwrod" binding:"required"`
}

//SessionKey session的键
const SessionKey = "SessionID"

//UserKey 获取用户的信息的Key
const UserKey = "user"

func (user User) empty() bool {
	return len(user.Name) == 0 || len(user.Passwrod) == 0
}

//Verify 需要验证的一组统一用verify.***来使用
type Verify struct {
	user      User
	userGourp *gin.RouterGroup
}

var defaultVerify = &Verify{}

//Default 验证模块
func Default() *Verify {
	return defaultVerify
}

func (verify *Verify) verifyLogin(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errmsg": err.Error()})
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
	} else {
		hexSessValue = sessValue.(string)
	}
	verify.user = user
	sess.Set(SessionKey, hexSessValue)
	sess.Set(UserKey, user)
	sess.Save()
	c.JSON(http.StatusOK, gin.H{"user": user, "session": hexSessValue})
}

//过滤不符合的 获取session 以及用户信息
func (verify *Verify) userFilter(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(UserKey)
	sessionID := session.Get(SessionKey)
	if sessionID == nil || user == nil { //
		//302转发
		c.Redirect(http.StatusFound, "/login")
		// fmt.Println("user group")
	} else {
		// u := user.(verify.User)
		// c.JSON(200, gin.H{"user": u, "SessionID": sessionID})
		fmt.Println("user verify")
	}

}
func (verify *Verify) getUserInfo(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(UserKey)
	sess := session.Get(SessionKey)
	c.JSON(200, gin.H{"user": user, "session": sess})
	// verify.userGourp.GET("/info", func(c *gin.Context) {
	// })
}

//RegistUserGourp 注册用户组 以/user开头
func (verify *Verify) registUserGourp(gourp string, engine *gin.Engine) {
	verify.userGourp = engine.Group(gourp, verify.userFilter)
}

//登出系统
func (verify *Verify) logout(c *gin.Context) {
	session := sessions.Default(c)
	name := session.Get(UserKey)
	session.Delete(SessionKey)
	session.Delete(UserKey)
	session.Save()
	c.JSON(200, gin.H{
		"logout": "success",
		"name":   name,
	})
}

// func (verify *Verify) GetUser() User {
// 	return verify.user
// }

//Login 验证用户登录
func Login(c *gin.Context) {
	defaultVerify.verifyLogin(c)
}

//Logout 登出系统
func Logout(c *gin.Context) {
	defaultVerify.logout(c)
}

//RegistUserGourp 注册需要验证的信息
func RegistUserGourp(gourp string, engine *gin.Engine) {
	defaultVerify.registUserGourp(gourp, engine)
}

//UserInfo 返回用户信息 url为/user/$path
func UserInfo(path string) {
	defaultVerify.userGourp.GET(path, defaultVerify.getUserInfo)
}

//GetVerify 获取验证模块
// func GetVerify() []gin.HandlerFunc {
// 	return []gin.HandlerFunc{deaultVerify.VerifyLogin}
// }
