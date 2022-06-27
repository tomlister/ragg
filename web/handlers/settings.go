package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getFontFamily(c *gin.Context) string {
	ff, err := c.Cookie("ff")
	if err == http.ErrNoCookie {
		ff = "Arial"
		c.SetCookie("ff", ff, 60*60*24*365*100, "/", "", false, false)
	}
	return ff
}

func getTextColor(c *gin.Context) string {
	tc, err := c.Cookie("tc")
	if err == http.ErrNoCookie {
		tc = "#000000"
		c.SetCookie("tc", tc, 60*60*24*365*100, "/", "", false, false)
	}
	return tc
}

func getBgColor(c *gin.Context) string {
	bgc, err := c.Cookie("bgc")
	if err == http.ErrNoCookie {
		bgc = "#ffffff"
		c.SetCookie("bgc", bgc, 60*60*24*365*100, "/", "", false, false)
	}
	return bgc
}

func handleSettings(c *gin.Context) {
	c.HTML(http.StatusOK, "settings.html", gin.H{
		"FontFamily": getFontFamily(c),
		"BgColor": getBgColor(c),
		"Color": getTextColor(c),
	})
}