package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type ChapterItem struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

const hp_base = "https://www.housepetscomic.com"

func main() {

	c := colly.NewCollector()

	c.OnHTML("select#chapter.postform", func(h *colly.HTMLElement) {
		fmt.Println(h.ChildAttrs("option.level-0", "value"))
	})

	c.Visit(hp_base + "/archive/")
}
