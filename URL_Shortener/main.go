package main

import (
	"net/http"
	"url_shortener/db"

	"github.com/gin-gonic/gin"
)

// func init() {
// 	if err := db.SetupDB(); err != nil {
// 		fmt.Printf("ERROR setting up DB: %v", err)
// 	}
// }

//starts the servers
func startServer() {
	router := gin.Default()
	router.GET("/", mainpage)
	router.GET("/:path", redirect)
	router.Run()
}

//mainpage of the server.
func mainpage(c *gin.Context) {
	c.String(http.StatusOK, "try adding /rickroll to the url.\nOther keyword includes dog.")
}

//redirects the user to given urls if the keyword matches
func redirect(c *gin.Context) {
	path := c.Param("path") //getting the keyword from the url
	//find key in map
	// if dest, ok := pathToUrls[path]; ok {
	// 	c.Redirect(http.StatusFound, dest)
	// 	return
	// }

	//connect to database
	conn, err := db.DB("./db/urlShortener.db")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "databse error",
		})
		return
	}
	defer conn.Close()

	//check bucket for key
	url := db.ViewKey(conn, path)
	if string(url) != "" {
		c.Redirect(http.StatusFound, string(url))
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "no url found",
	})
}

func main() {
	// //yaml flag to take in yaml files.
	// yamlFile := flag.String("yaml", "", "path to your yaml list of urls and keyword")
	// jsonFile := flag.String("json", "", "path to your json list of urls and keyword")
	// flag.Parse()

	// fmt.Println(*yamlFile, *jsonFile)
	// //fetches list of urls from given file to a map
	// if *yamlFile != "" {
	// 	data := readFile(*yamlFile)
	// 	parseYAML(data)
	// } else if *jsonFile != "" {
	// 	data := readFile(*jsonFile)
	// 	parseJSON(data)
	// } else {
	// 	fmt.Println("ERROR no file input given.")
	// 	os.Exit(1)
	// }

	//starts the url shortener server
	startServer()
}

// //parses yaml contents to a map[string]string
// func parseYAML(listdata []byte) {
// 	//unmarshal file contents to []shortURL
// 	var data []shortURL
// 	yaml.Unmarshal(listdata, &data)

// 	//convert struct to map[string]string
// 	for _, v := range data {
// 		pathToUrls[v.Path] = v.Url
// 	}
// }

// //parses json contents to map[string]string
// func parseJSON(listdata []byte) {
// 	//unmarshal file contents into []shortURL
// 	var data []shortURL
// 	yaml.Unmarshal(listdata, data)

// 	//convert struct to map[string]string
// 	for _, v := range data {
// 		pathToUrls[v.Path] = v.Url
// 	}
// }

// //read []byte contents of a file
// func readFile(filepath string) []byte {
// 	data, err := os.ReadFile(filepath)
// 	if err != nil {
// 		fmt.Println("ERROR reading file ", err)
// 		os.Exit(1)
// 	}
// 	return data
// }

// type shortURL struct {
// 	Path string `yaml: path`
// 	Url  string `yaml: url`
// }

//pathToUrls maps websites to keywords.
// var pathToUrls map[string]string = map[string]string{
// 	"dogs":     "https://www.google.com/search?q=dogs&source=lnms&tbm=isch&sa=X&ved=2ahUKEwi5kJu0pfzzAhVtzDgGHXuLAQwQ_AUoAXoECAEQAw&biw=1536&bih=746&dpr=1.25",
// 	"rickroll": "https://www.youtube.com/watch?v=dQw4w9WgXcQ&ab_channel=RickAstley",
// }

// var pathToUrls = make(map[string]string)

// // //yaml contains an array of the keyword and url.
// // var yamlList string = `
// // - path: dogs
// //   url: https://www.google.com/search?q=dogs&source=lnms&tbm=isch&sa=X&ved=2ahUKEwi5kJu0pfzzAhVtzDgGHXuLAQwQ_AUoAXoECAEQAw&biw=1536&bih=746&dpr=1.25
// // - path: rickroll
// //   url: https://www.youtube.com/watch?v=dQw4w9WgXcQ&ab_channel=RickAstley
// // `
