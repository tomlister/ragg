package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getSources(c *gin.Context) []string {
	srcs, err := c.Cookie("sources")
	if err != nil {
		setSources(c, []string{})
		return []string{}
	}
	var sources []string
	err = json.Unmarshal([]byte(srcs), &sources)
	if err != nil {
		setSources(c, []string{})
		return []string{}
	}
	return sources
}

func setSources(c *gin.Context, sources []string) error {
	data, err := json.Marshal(&sources)
	if err != nil {
		return err
	}
	c.SetCookie("sources", string(data), 60*60*24*365*100, "/", "", false, false)
	return nil
}

func handleSources(c *gin.Context) {
	sources := getSources(c)
	c.HTML(http.StatusOK, "sources.html", gin.H{
		"Sources": sources,
		"FontFamily": getFontFamily(c),
		"BgColor": getBgColor(c),
		"Color": getTextColor(c),
	})
}