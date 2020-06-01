package main

import (
	"github.com/ekcode/gottto/cmd/lotto"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./views", true)))
	r.Static("assets", "assets")
	//r.LoadHTMLGlob("views/**")


	r.GET("/api", func(c *gin.Context) {
		//c.HTML(200, "index.html", gin.H{"당첨번호": lotto.Generate()})
		c.JSON(200, gin.H{"당첨번호": lotto.Generate()})
	})

	r.Run(":8000")
}
