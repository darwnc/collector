package main

import (
	"encoding/gob"

	"github.com/darwnc/collector/exercises"

	"github.com/darwnc/collector/verify"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	//需要注册，否则无法获取到该结构体
	gob.Register(verify.User{})

	gin.SetMode(gin.DebugMode)
	engine := gin.New()
	store := cookie.NewStore([]byte("gcookie"))
	engine.Use(sessions.Sessions("gsession", store))

	engine.Use(gin.Logger(), gin.Recovery())
	engine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"pong": "hello"})
	})
	//需要验证的模块以/user为开头
	verify.RegistUserGourp("/user", engine)
	verify.UserInfo("/info")

	//无需验证的模块
	engine.GET("/login", verify.Login)
	engine.POST("/login", verify.Login)
	engine.GET("/logout", verify.Logout)

	engine.GET("/plot", plotTest)
	engine.GET("/image", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, `<img src="/plot"/>`)
	})

	engine.GET("/exercises/watch/", exercises.BinWatch)
	engine.GET("/exercises/removeDigits/", exercises.RemoveDigits)
	// engine.GET("/login", func(c *gin.Context) {
	// 	session := sessions.Default(c)
	// 	name := c.Query("name")
	// 	pwd := c.Query("password")
	// 	if len(name) == 0 || len(pwd) == 0 {
	// 		c.JSON(200, gin.H{"user": "empty", "password": "empty"})
	// 		return
	// 	}
	// 	sessionID := session.Get("sessionID")
	// 	if sessionID == nil {
	// 		buff := make([]byte, 16)
	// 		n, err := rand.Read(buff)
	// 		if err != nil {
	// 			c.JSON(200, gin.H{"login": "gen session err", "err": err})
	// 			return
	// 		}
	// 		sessionID = hex.EncodeToString(buff)
	// 		session.Set(sessionKey, sessionID)
	// 		session.Save()
	// 		c.JSON(200, gin.H{"login": "gen session success",
	// 			"count": n, sessionKey: sessionID})
	// 		return
	// 	}
	// 	session.Set(sessionKey, sessionID)
	// 	session.Set("name", name)
	// 	session.Save()
	// 	c.JSON(200, gin.H{
	// 		"name": name,
	// 	})

	// })

	engine.Run(":8080")
}

func plotTest(c *gin.Context) {
	groupA := plotter.Values{20, 35, 30, 35, 27}
	groupB := plotter.Values{25, 32, 34, 20, 25}
	groupC := plotter.Values{12, 28, 15, 21, 8}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Bar chart"
	p.Y.Label.Text = "Heights"

	w := vg.Points(20)

	barsA, err := plotter.NewBarChart(groupA, w)
	if err != nil {
		panic(err)
	}
	barsA.LineStyle.Width = vg.Length(0)
	barsA.Color = plotutil.Color(0)
	barsA.Offset = -w

	barsB, err := plotter.NewBarChart(groupB, w)
	if err != nil {
		panic(err)
	}
	barsB.LineStyle.Width = vg.Length(0)
	barsB.Color = plotutil.Color(1)

	barsC, err := plotter.NewBarChart(groupC, w)
	if err != nil {
		panic(err)
	}
	barsC.LineStyle.Width = vg.Length(0)
	barsC.Color = plotutil.Color(2)
	barsC.Offset = w

	p.Add(barsA, barsB, barsC)
	p.Legend.Add("Group A", barsA)
	p.Legend.Add("Group B", barsB)
	p.Legend.Add("Group C", barsC)
	p.Legend.Top = true
	p.NominalX("One", "Two", "Three", "Four", "Five")
	wt, _ := p.WriterTo(5*vg.Inch, 3*vg.Inch, "png")
	//有其他接口调用的方式
	wt.WriteTo(c.Writer)
	//	第二种显示图片方式
	// c.Header("Content-Type", "text/html; charset=utf-8")
	// img := `<img src="data:image/png;base64,`
	// buff := &bytes.Buffer{}
	// wt.WriteTo(buff)
	// base64.NewEncoder(base64.StdEncoding, buff).Write(buff.Bytes())
	// png := base64.StdEncoding.EncodeToString(buff.Bytes())
	// c.Writer.Write([]byte(img + png + `"/>`))
	// if err := p.Save(5*vg.Inch, 3*vg.Inch, "barchart.png"); err != nil {
	// 	panic(err)
	// }
}
