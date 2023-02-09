package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

type LoginPageData struct {
	Header    string
	Image     string
	Button    string
	Subheader string
	Text      string
}

//go:embed tmpl
var tmplDir embed.FS

//go:embed assets/*
var assets embed.FS

func main() {
	assetsFs := http.FileServer(http.FS(assets))

	mux := http.NewServeMux()
	mux.Handle("/assets/", assetsFs)
	mux.HandleFunc("/", indexPage)

	fmt.Println("ready")

	http.ListenAndServe(":8080", mux)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	data := LoginPageData{
		Header:    "Gopher's Journey",
		Image:     "assets/img/gopher-journey.png",
		Text:      `The Go programming language has a mascot, a gopher, which has been designed by the talented artist Renee French. The gopher has inspired a variety of merchandise such as T-shirts, posters, stickers, and home decor items, which are all designed and sold by independent artists. The gopher has become a symbol of the Go community, and there is an initiative called GopherSource that aims to strengthen and diversify the ecosystem of the language. The gopher has become a "bat signal" for both beginner and experienced members of the Go community. For more information on the gopher, you can read the blog post on the Go programming language's website.`,
		Subheader: "Gopher Journey is me",
		Button:    "Say Hello",
	}
	printPage(w, "index.html", data)
}

func printPage(w http.ResponseWriter, fileName string, data interface{}) error {
	//t, _ := template.ParseFiles(fmt.Sprintf("./web/%s", fileName))
	t, err := template.ParseFS(tmplDir, fmt.Sprintf("tmpl/%s", fileName))
	if err != nil {
		panic(err)
	}
	return t.Execute(w, data)
}
