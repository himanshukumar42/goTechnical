// get request and post request  post receive data

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/himanshuk42/sampleapi/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"health": "ok"})
	})
	r.GET("/movies", controllers.MoviesGetHandler)
	r.GET("/movies/:id", controllers.MovieGetHandler)
	r.POST("/movies", controllers.MoviesPostHandler)
	r.Run()
}