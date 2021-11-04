package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//pathToUrls maps websites to keywords.
var pathToUrls map[string]string = map[string]string{
	"dogs":     "https://www.google.com/search?q=dogs&source=lnms&tbm=isch&sa=X&ved=2ahUKEwi5kJu0pfzzAhVtzDgGHXuLAQwQ_AUoAXoECAEQAw&biw=1536&bih=746&dpr=1.25",
	"rickroll": "https://www.youtube.com/watch?v=dQw4w9WgXcQ&ab_channel=RickAstley",
}

func main() {
	startServer(&pathToUrls)
}

//starts the servers
func startServer(path *map[string]string) {
	router := gin.Default()
	router.GET("/", hello)
	router.GET("/:path", urlShortener)
	//router.NoRoute(handleURL)

	router.Run()
}

func urlShortener(c *gin.Context) {
	path := c.Param("path")
	if dest, ok := pathToUrls[path]; ok {
		c.Redirect(http.StatusFound, dest)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "no url found ",
	})
}

func hello(c *gin.Context) {
	c.String(http.StatusOK, "try adding /rickroll to the url.\nOther keyword includes dog.")
}
