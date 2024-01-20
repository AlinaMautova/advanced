package main

import (
	"github.com/bmizerany/pat"
	_ "github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() *pat.PatternServeMux {

	mux := pat.New()

	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/news/create", http.HandlerFunc(app.createNewsForm))
	mux.Post("/news/create", http.HandlerFunc(app.createNews))
	mux.Get("/snippet/:id", http.HandlerFunc(app.showNews))
	mux.Get("/contact", http.HandlerFunc(app.contact))
	mux.Get("/news/staff", http.HandlerFunc(app.staffNewsHandler))
	mux.Get("/news/applicants", http.HandlerFunc(app.applicantsNewsHandler))
	mux.Get("/news/researches", http.HandlerFunc(app.researchesNewsHandler))
	mux.Get("/news/students", http.HandlerFunc(app.studentsNewsHandler))
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return mux
}
