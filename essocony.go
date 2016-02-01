package esso

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var App = mux.NewRouter()

var ArticleSlice Articles
var ArticleHash ArticleMap

func init() {
	var err error
	ArticleSlice, err = LoadArticles("articles/*.html")
	if err != nil {
		panic(err)
	}
	ArticleHash = ArticleSlice.ArticleMap()
	App.HandleFunc("/", ArticlesHandler)
	App.HandleFunc("/articles/", ArticlesHandler)
	App.HandleFunc("/articles/{slug}", ArticleHandler)
	App.Handle("/static/{page:.*}", http.FileServer(http.Dir("public")))
}

var baseTpl = template.Must(template.ParseFiles("templates/base.html"))

var articlesTpl = template.Must(template.Must(baseTpl.Clone()).ParseFiles("templates/article.html"))

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	data := Page{Articles: ArticleSlice, Title: "Essocony: All Articles"}
	articleTpl.Execute(w, data)
}

var articleTpl = template.Must(template.Must(baseTpl.Clone()).ParseFiles("templates/article.html"))

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	article, found := ArticleHash[slug]
	if !found {
		http.NotFound(w, r)
		return
	}
	data := Page{Articles: Articles{article}, Title: "Essocony: " + article.Title}
	articlesTpl.Execute(w, data)
}
