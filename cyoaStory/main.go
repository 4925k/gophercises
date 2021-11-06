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

	// tpl := template.Must(template.New("").Parse(customTemplate))
	// h := cyoa.NewHandler(story, cyoa.WithTemplate(tpl), cyoa.WithPathFunc(pathFn))
	//mux:= http.NewServeMux()
	//mux.Handle("/story/", h)

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting CYOA server on port %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

}

//using custom handlers
// func pathFn(r *http.Request) string {
// 	path := r.URL.Path
// 	if path == "/story" || path == "/story/" {
// 		path = "/story/intro"
// 	}
// 	return path[len("/story/"):]
// }

// var customTemplate = `<!DOCTYPE html>
// <html>
//   <head>
//     <meta charset="utf-8">
//     <title>Choose Your Own Adventure</title>
//   </head>
//   <body>
//     <section class="page">
//       <h1>{{.Title}}</h1>
//       {{range .Paragraphs}}
//         <p>{{.}}</p>
//       {{end}}
//       {{if .Options}}
//         <ul>
//         {{range .Options}}
//           <li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
//         {{end}}
//         </ul>
//       {{else}}
//         <h3>The End</h3>
//       {{end}}
//     </section>
//     <style>
//       body {
//         font-family: helvetica, arial;
//       }
//       h1 {
//         text-align:center;
//         position:relative;
//       }
//       .page {
//         width: 80%;
//         max-width: 500px;
//         margin: auto;
//         margin-top: 40px;
//         margin-bottom: 40px;
//         padding: 80px;
//         background: #FFFCF6;
//         border: 1px solid #eee;
//         box-shadow: 0 10px 6px -6px #777;
//       }
//       ul {
//         border-top: 1px dotted #ccc;
//         padding: 10px 0 0 0;
//         -webkit-padding-start: 0;
//       }
//       li {
//         padding-top: 10px;
//       }
//       a,
//       a:visited {
//         text-decoration: none;
//         color: #6295b5;
//       }
//       a:active,
//       a:hover {
//         color: #7792a2;
//       }
//       p {
//         text-indent: 1em;
//       }
//     </style>
//   </body>
// </html>`
