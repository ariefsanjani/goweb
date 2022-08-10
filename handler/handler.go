package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Yah ada yang error, keep calm! sedang diperbaiki", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Yah ada yang error, keep calm! sedang diperbaiki", http.StatusInternalServerError)
		return
	}

}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("views", "contact.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Yah ada yang error, keep calm! sedang diperbaiki", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Yah ada yang error, keep calm! sedang diperbaiki", http.StatusInternalServerError)
		return
	}
}
