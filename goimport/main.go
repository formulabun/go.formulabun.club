package main

import (
	"html/template"
	"log"
	"net/http"
)

type templateData struct {
	Path string
}

const htmlPage = `<html>
<head>
<meta name="go-import" content="go.formulabun.club{{.Path}} git https://github.com/formulabun{{.Path}}">
</head>
<body>
<a href="https://godoc.formulabun.club{{.Path}}">Redirecting to documentation.</a>
</body>
</html>`

var templ *template.Template

func handler(w http.ResponseWriter, r *http.Request) {
	data := templateData{r.URL.EscapedPath()}
	err := templ.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	var err error
	templ, err = template.New("redirect").Parse(htmlPage)
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8001", nil))
}
