package handlers

import (
	"html"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"github.com/tomlister/ragg/rss"
)

func sanitise(feed []*rss.AttributedItem) {
	p := bluemonday.StrictPolicy()
	for _, el := range feed {
		for i, dsc := range el.Desc {
			(*el).Desc[i] = html.UnescapeString(p.Sanitize(dsc))
		}
	}
}

func sortOldNew(items []*rss.AttributedItem, last time.Time) (old []*rss.AttributedItem, new []*rss.AttributedItem) {
	old = []*rss.AttributedItem{}
	new = []*rss.AttributedItem{}
	for _, el := range items {
		if el.Item.PublishedParsed.After(last) {
			new = append(new, el)
		} else {
			old = append(old, el)
		}
	}
	return
}

func setLastVisit(c *gin.Context) {
	c.SetCookie("last-visit", strconv.Itoa(int(time.Now().Unix())), 60*60*24*365*100, "/", "", false, true)
}

func handleIndex(c *gin.Context) {
	epochStr, err := c.Cookie("last-visit")
	setLastVisit(c)
	sources := getSources(c)
	if err == http.ErrNoCookie {
		feed := rss.GetFeed(sources)
		sanitise(feed)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"News": feed,
			"FontFamily": getFontFamily(c),
			"BgColor": getBgColor(c),
			"Color": getTextColor(c),
		})
		return
	}
	epoch, err := strconv.Atoi(epochStr)
	if err != nil {
		c.Error(err)
		return
	}
	lastVisit := time.Unix(int64(epoch), 0)
	feed := rss.GetFeed(sources)
	sanitise(feed)
	old, new := sortOldNew(feed, lastVisit)
	c.HTML(http.StatusOK, "index-new.html", gin.H{
		"New": new,
		"Old": old,
		"FontFamily": getFontFamily(c),
		"BgColor": getBgColor(c),
		"Color": getTextColor(c),
	})
}