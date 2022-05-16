package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var Articles = []Article{
	{Id: "1", Title: "Hello", Content: "Content of hello"},
	{Id: "2", Title: "Hello 2", Content: "Content of hello 2"},
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome !"))
}

func handleArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Articles)
}

func handleArticleDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, article := range Articles {
		if article.Id == id {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func handleCreateArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var article Article
	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

func handleDeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
			w.Write([]byte("success"))
		}
	}
}

func handleUpdateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	reqBody, _ := ioutil.ReadAll(r.Body)
	var updatedArticle Article
	json.Unmarshal(reqBody, &updatedArticle)

	for index, article := range Articles {
		if article.Id == id {
			Articles[index] = updatedArticle
			json.NewEncoder(w).Encode(updatedArticle)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handleHome)
	router.HandleFunc("/articles", handleCreateArticle).Methods("POST")
	router.HandleFunc("/articles", handleArticles)
	router.HandleFunc("/articles/{id}", handleDeleteArticle).Methods("DELETE")
	router.HandleFunc("/articles/{id}", handleUpdateArticle).Methods("PUT")
	router.HandleFunc("/articles/{id}", handleArticleDetail)

	http.ListenAndServe(":3000", router)
}
