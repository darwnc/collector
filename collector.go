package main

import (
	"encoding/gob"
	"fmt"
	"net/http"

	"github.com/darwnc/collector/exercises"
	"github.com/darwnc/collector/exter"

	"github.com/darwnc/collector/verify"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var htmlFiles = []string{"./static/html/index.html", "./static/html/temp/header.html",
	"./static/html/test.html", "./static/html/temp/footer.html"}

func main() {
	//需要注册，否则无法获取到该结构体
	gob.Register(verify.User{})
	// dir, _ := os.Getwd()
	// fmt.Println("currentdir=", dir)
	gin.SetMode(gin.DebugMode)
	engine := gin.New()
	store := cookie.NewStore([]byte("gcookie"))
	engine.Use(sessions.Sessions("gsession", store))
	engine.StaticFS("/resources", http.Dir("./static/resources"))

	// engine.HTMLRender.Instance(string, interface{})
	// testTemp := template.Must(template.ParseFiles("./static/html/test.html", "./static/html/temp/header.html", "./static/html/temp/footer.html"))
	// engine.SetHTMLTemplate(testTemp)
	// engine.LoadHTMLGlob("static/html/*/*")
	// engine.StaticFile("/favicon.ico", "/Users/Jack/Documents/golearn/collector/favicon.ico")
	engine.LoadHTMLFiles(htmlFiles...)
	engine.Use(gin.Logger(), gin.Recovery())
	engine.GET("/", func(c *gin.Context) {
		// c.JSON(200, gin.H{"pong": "hello"})
		// indexTemp := template.Must(template.ParseFiles())
		// engine.SetHTMLTemplate(indexTemp)
		//define "home/index.html" 否则默认index.html即文件名加后缀
		c.HTML(http.StatusOK, "home/index.html", struct{ Title string }{"首页"})
	})
	engine.GET("/test", func(c *gin.Context) {
		// testTemp := template.Must(template.ParseFiles("./static/html/test.html", "./static/html/temp/header.html", "./static/html/temp/footer.html"))
		// engine.SetHTMLTemplate(testTemp)
		c.HTML(200, "test.html", struct{ Title string }{"测试"})
	})
	//需要验证的模块以/user为开头
	verify.RegistUserGroup("/user", engine)
	verify.UserInfo("/info")

	// engine.Any("/login", verify.Login)
	// engine.Any("/logout", verify.Logout)
	engine.GET("/login", verify.Login)
	engine.GET("/logout", verify.Logout)
	//无需验证的模块
	engine.GET("/jsonp", func(c *gin.Context) {
		c.JSONP(200, gin.H{"foo": "bar"})
	})
	engine.GET("/plot", plotTest)
	engine.GET("/image", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, `<img src="/plot"/>`)
	})
	//转发，浏览器无变化，服务器内部
	engine.GET("/t1", func(c *gin.Context) {
		c.Request.URL.Path = "/t2"
		engine.HandleContext(c)
		// reader := strings.NewReader("string reader")
		// reader.WriteTo(w)
		// ioutil.ReadAll(r)

	})
	engine.GET("/t2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"redirect": "form t1"})
	})
	verify.RegistAuthGroup("/auth", engine)
	verify.AuthUser("/look")

	engine.GET("/exercises/watch/", exercises.BinWatch)
	engine.GET("/exercises/removeDigits/", exercises.RemoveDigits)
	var tr exter.TestRequest
	engine.GET("/test_crypto", exter.Wrap(&tr, func() interface{} {
		fmt.Println("service=", tr.Header.Service)
		fmt.Printf("%#v", tr)
		return tr
	}))

	engine.Run(":8080")
	// autotls.Run(engine, "172.16.0.31")
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
