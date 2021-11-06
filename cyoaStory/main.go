package main

import (
	"cyoaStory/cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	//setting up flags
	filename := flag.String("file", "cyoa/story.json", "JSON file with the CYOA story")
	port := flag.Int("port", 8080, "port to run cyoa story on")
	flag.Parse()

	//opening story json file
	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	//decoding json data into map
	story, err := cyoa.JsonStory(file)
	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting CYOA server on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

}
