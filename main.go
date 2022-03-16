package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Rock", Artist: "Rock Kumar", Price: 55.25},
	{ID: "2", Title: "HipHop", Artist: "HipHop Kumar", Price: 155.25},
	{ID: "3", Title: "Jazz", Artist: "Jazz Kumar", Price: 255.25},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)

	router.Run("localhost:8000")
}

// getAlbums responds with the list of all albums as JSON.

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"detail": "album not found"})
}
