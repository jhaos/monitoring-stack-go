package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Healthcheck struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func getHealthcheck(con *gin.Context) {
	con.IndentedJSON(http.StatusOK, estado)
}

var estado = Healthcheck{Status: "green", Code: 200}

func main() {

	router := gin.Default()
	router.GET("/api/health", getHealthcheck)

	router.Run("localhost:8081")

}

/*type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}*/
