package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func handler(c *gin.Context) {
	jsonData, err := c.GetRawData()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%d bytes received", len(jsonData))

	c.String(http.StatusOK, "Success")
}
