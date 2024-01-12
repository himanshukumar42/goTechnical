package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/himanshuk42/sampleapi/models"
)

type MoviesInput struct {
	Title string `json:"title"`
	Director string `json:"director"`
}

type MoviesOutput struct {
	ID string `json:"_id,omitempty"`
	Title string `json:"title,omitempty"`
	Director string `json:"director,omitempty"`
}

func MovieGetHandler(c *gin.Context) {

}


var (
	movies = make(chan models.Movies, 5)
)

// post request value in channel , and in GET request get values from channel

func MoviesGetHandler(c *gin.Context) {
	// client, err := models.NewClient()
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err})
	// }
	// mydata := <-movies
	defer close(movies)
	var mydata []interface{}
	
	for dt := range movies {
		fmt.Println("Movie values: ", dt)
		// mydata = append(mydata, dt)

	}
	
	
	c.JSON(http.StatusOK, gin.H{"data": mydata})
	return
	// db := client.Database("mydb")
	// collection := db.Collection("mycollection")
	// // var filter map[string]string 
	// projection := map[string]int{
	// 	"title": 1,
	// }
	// cursor, err := collection.Find(context.Background(), bson.M{}, options.Find().SetProjection(projection))
	// if err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": err})
	// 	return
	// }

	// defer cursor.Close(context.Background())

	// var moviesData[] MoviesOutput
	// err = cursor.All(context.Background(), &moviesData)
	// if err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	// 	return
	// }
	// fmt.Println("Movies Data: ", moviesData)
	// c.JSON(http.StatusOK, gin.H{"data": moviesData})
}

func MoviesPostHandler(c *gin.Context) {
	var movie MoviesInput
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
		return
	}

	movies <- models.Movies(movie)
	// client, err := models.NewClient()
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err})
	// }
	// db := client.Database("mydb")
	// collection := db.Collection("mycollection")
	// result, err := collection.InsertOne(context.Background(), movie)
	// if err != nil {
	// 	c.JSON(http.StatusCreated, gin.H{"error": err})
	// }
	
	c.JSON(http.StatusOK, gin.H{"InsertedID": movie})
}