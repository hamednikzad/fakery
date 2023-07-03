package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hamednikzad/fakery/pkg/color"
	"log"
	"os"
)

func registerControllers(e *gin.Engine) {
	log.Println("POST on /data/input-a")
	log.Println("POST on /data/input-b")
	log.Printf("Red: %s %s None: %s %s", color.ColorBLU, "red string", color.ColorMAG, "colorless string")

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

	log.Printf("Statring server on %s\n", addr)

	engine.Run(addr)
}
