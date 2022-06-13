package handlers

import (
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("web/pages/*.html")
	r.GET("/settings", handleSettings)
	r.GET("/sources", handleSources)
	r.GET("/", handleIndex)
	return r
}