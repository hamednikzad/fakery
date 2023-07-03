package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hamednikzad/fakery/pkg/color"
	"github.com/hamednikzad/fakery/pkg/mylog"
	"log"
	"os"
)

func registerControllers(e *gin.Engine) {
	mylog.Println("POST on /data/input-a", color.CYN)
	mylog.Println("POST on /data/input-b", color.CYN)

	routes := e.Group("/data")
	routes.POST("/input-a", handler)
	routes.POST("/input-b", handler)
}

func Run(addr string) {
	engine := gin.Default()
	registerControllers(engine)

	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	mylog.Println(fmt.Sprintf("Starting server on %s", addr), color.YEL)

	log.Fatal(engine.Run(addr))
}
