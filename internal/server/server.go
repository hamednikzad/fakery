package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hamednikzad/fakery/pkg/color"
	"github.com/hamednikzad/fakery/pkg/mylog"
	"log"
	"os"
)

func registerControllers(e *gin.Engine, addr string) {
	addrA := fmt.Sprintf("%s/data/input-a", addr)
	addrB := fmt.Sprintf("%s/data/input-b", addr)
	mylog.Println("POST on "+addrA, color.CYN)
	mylog.Println("POST on "+addrB, color.CYN)

	routes := e.Group("/data")
	routes.POST("/input-a", handler)
	routes.POST("/input-b", handler)
}

func Run(addr string) {
	engine := gin.Default()
	registerControllers(engine, addr)

	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	mylog.Println(fmt.Sprintf("Starting server on %s", addr), color.YEL)

	log.Fatal(engine.Run(addr))
}
