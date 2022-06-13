package rss

import (
	"context"
	"fmt"
	"html"
	"sort"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
)

type AttributedItem struct{
	Item *gofeed.Item
	Source string
	Desc []string
}

func getFeed(source string) (*gofeed.Feed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	fp := gofeed.NewParser()
	feed, err := fp.ParseURLWithContext(source, ctx)
	return feed, err
}

func attribute(feed *gofeed.Feed) []*AttributedItem {
	items := []*AttributedItem{}
	for _, item := range feed.Items {
		lines := strings.Split(strings.ReplaceAll(html.UnescapeString(item.Description), "</p>", ""), "<p>")
		items = append(items, &AttributedItem{
			Item: item,
			Source: feed.Title,
			Desc: lines,
		})
	}
	return items
}

func GetFeed(sources []string) []*AttributedItem {
	/*sources := []string{
		"https://www.9news.com.au/rss",
		"https://www.theguardian.com/au/rss",
	}*/
	items := []*AttributedItem{}
	for _, s := range sources {
		feed, err := getFeed(s)
		fmt.Printf("%v", feed)
		if err != nil {
			continue
		}
		attributed := attribute(feed)
		items = append(items, attributed...)
	}
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Item.PublishedParsed.After(*items[j].Item.PublishedParsed)
	})
	return items
}