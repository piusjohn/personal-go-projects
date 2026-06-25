package main
import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)
//using gin framework
// album represents data about a record album.
type album struct{
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Price  float64 `json:"price"`
}
// albums slice to seed record album data.

var albums = []album{
	{ID: "1", Title: "f2f", Artist: "piuspraise", Price: 10.0},
	{ID: "2", Title: "Open Up", Artist: "Dunsin Oyekan", Price: 20.0},
	{ID: "3", Title: "Mi-Corazon", Artist: "Don Moen", Price: 15.0},
}

//getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)	
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context){
	var newalbum album 
	if err := c.BindJSON(&newalbum); err != nil{
		return
	}
	albums = append(albums, newalbum)
	c.IndentedJSON(http.StatusCreated, newalbum)
}

func getAlbumById(c *gin.Context){
	id := c.Param("id")

	for _,  a := range albums{
		if a.ID == id{
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main(){
	r := gin.Default()
	r.GET("/albums", getAlbums)
	r.POST("/albums", postAlbums)
	r.GET("/albums/:id", getAlbumById)
	fmt.Println("server is running on port :8080")
	r.Run(":8080")
}