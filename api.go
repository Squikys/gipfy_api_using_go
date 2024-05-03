package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type meme struct {
	ID  int    `json:"id"`
	URL string `json:"URL"`
}

// albums slice to seed record album data.
var memes = []meme{}

var query string

func getMemes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, memes)
}

func getMemeByID(c *gin.Context) {
	query = c.Param("id")

	d := colly.NewCollector(
		colly.AllowedDomains("giphy.com"),
	)
	d.OnHTML(".giphy-grid", func(e *colly.HTMLElement) {
		links := e.ChildAttrs("a", "href")
		for i := 0; i < 5; i++ {
			var newMeme meme
			newMeme.ID = i
			newMeme.URL = links[i]
			memes = append(memes, newMeme)
		}
	})

	d.Visit("https://giphy.com/search/" + query)
	time.Sleep(1 * time.Second)
	c.IndentedJSON(http.StatusOK, memes)

	memes = []meme{}
}
func main() {
	router := gin.Default()
	router.GET("/meme", getMemes)
	router.GET("/meme/:id", getMemeByID)
	router.Run("localhost:8080")
}
