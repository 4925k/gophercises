package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type urlPath struct {
	Path string `yaml: path`
	Url  string `yaml: url`
}

//pathToUrls maps websites to keywords.
// var pathToUrls map[string]string = map[string]string{
// 	"dogs":     "https://www.google.com/search?q=dogs&source=lnms&tbm=isch&sa=X&ved=2ahUKEwi5kJu0pfzzAhVtzDgGHXuLAQwQ_AUoAXoECAEQAw&biw=1536&bih=746&dpr=1.25",
// 	"rickroll": "https://www.youtube.com/watch?v=dQw4w9WgXcQ&ab_channel=RickAstley",
// }

var pathToUrls = make(map[string]string)

//yaml contains an array of the keyword and url.
var yamlList string = `
- path: dogs
  url: https://www.google.com/search?q=dogs&source=lnms&tbm=isch&sa=X&ved=2ahUKEwi5kJu0pfzzAhVtzDgGHXuLAQwQ_AUoAXoECAEQAw&biw=1536&bih=746&dpr=1.25
- path: rickroll
  url: https://www.youtube.com/watch?v=dQw4w9WgXcQ&ab_channel=RickAstley
`

func main() {
	parseYAML()
	fmt.Println(pathToUrls)
	startServer(&pathToUrls)
}

//starts the servers
func startServer(path *map[string]string) {
	router := gin.Default()
	router.GET("/", mainpage)
	router.GET("/:path", redirect)
	router.Run()
}

//redirects the user to given urls if the keyword matches
func redirect(c *gin.Context) {
	path := c.Param("path") //getting the keyword from the url
	if dest, ok := pathToUrls[path]; ok {
		c.Redirect(http.StatusFound, dest)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "no url found",
	})
}

//parses yaml string to a map[string]string
func parseYAML() {
	//parse yaml string to []struct
	var data []urlPath
	yamlByte := []byte(yamlList)
	yaml.Unmarshal(yamlByte, &data)

	//convert yaml []struct to map[string]string
	for _, v := range data {
		pathToUrls[v.Path] = v.Url
	}

}

//mainpage of the server.
func mainpage(c *gin.Context) {
	c.String(http.StatusOK, "try adding /rickroll to the url.\nOther keyword includes dog.")
}
