package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/darwnc/collector/verify"

	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const sessionKey = "SessionID"

type user struct {
	name     string
	password string
}

func main() {
	gin.SetMode(gin.DebugMode)
	engine := gin.New()

	store := cookie.NewStore([]byte("gcookie"))
	engine.Use(sessions.Sessions("gsession", store))

	engine.Use(gin.Logger(), gin.Recovery())
	engine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"pong": "hello"})
	})
	engine.POST("/post", verify.GetVerify()...)
	group := engine.Group("/user", func(c *gin.Context) {
		session := sessions.Default(c)
		name := session.Get("name")
		sessionID := session.Get(sessionKey)
		if sessionID == nil || name == nil { //
			//302转发
			c.Redirect(http.StatusMovedPermanently, "/login")
			fmt.Println("user group")
		} else {
			fmt.Println("user verify")
		}

	})
	{
		group.GET("/cosor", func(c *gin.Context) {
			session := sessions.Default(c)
			name := session.Get("name")
			sess := session.Get(sessionKey)
			c.JSON(200, gin.H{"user/cosor": name, "session": sess})
		})
	}
	engine.GET("/login", func(c *gin.Context) {
		session := sessions.Default(c)
		name := c.Query("name")
		pwd := c.Query("password")
		if len(name) == 0 || len(pwd) == 0 {
			c.JSON(200, gin.H{"user": "empty", "password": "empty"})
			return
		}
		sessionID := session.Get("sessionID")
		if sessionID == nil {
			buff := make([]byte, 16)
			n, err := rand.Read(buff)
			if err != nil {
				c.JSON(200, gin.H{"login": "gen session err", "err": err})
				return
			}
			sessionID = hex.EncodeToString(buff)
			session.Set(sessionKey, sessionID)
			session.Save()
			c.JSON(200, gin.H{"login": "gen session success",
				"count": n, sessionKey: sessionID})
			return
		}
		session.Set(sessionKey, sessionID)
		session.Set("name", name)
		session.Save()
		c.JSON(200, gin.H{
			"name": name,
		})

	})
	engine.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		name := session.Get("name")
		session.Delete(sessionKey)
		session.Delete("name")
		session.Save()
		c.JSON(200, gin.H{
			"logout": "success",
			"name":   name,
		})
	})
	engine.Run(":8080")
}

func userVerify(c *gin.Context) {

}
