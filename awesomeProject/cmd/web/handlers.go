package main

import (
	"errors"
	"fmt"
	_ "fmt"
	"lina.net/aitunewstask/pkg/models"
	"log"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	newsList, err := app.news.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "home.page.tmpl", &templateData{
		NewsList: newsList,
	})
}

func (app *application) showNews(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	s, err := app.news.Get(id)
	if err != nil {
		log.Println(err) // or fmt.Println(err) for debugging
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	flash := app.session.PopString(r, "flash")

	app.render(w, r, "show.page.tmpl", &templateData{
		Flash: flash,
		News:  s,
	})
}
func (app *application) createNewsForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "createNews.page.tmpl", &templateData{})

}

func (app *application) createNews(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	// Use the r.PostForm.Get() method to retrieve the relevant data fields
	// from the r.PostForm map.
	title := r.FormValue("title")
	content := r.FormValue("content")
	expires := r.FormValue("expires")
	category := r.FormValue("category")
	switch category {
	case "staff":
		http.Redirect(w, r, "/news/staff", http.StatusSeeOther)
	case "applicants":
		http.Redirect(w, r, "/news/applicants", http.StatusSeeOther)
	case "researches":
		http.Redirect(w, r, "/news/researches", http.StatusSeeOther)
	case "students":
		http.Redirect(w, r, "/news/students", http.StatusSeeOther)
	default:
		// Handle unknown category
		http.NotFound(w, r)
	}

	// Create a new snippet record in the database using the form data.
	_, err = app.news.Insert(title, content, expires, category)
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.session.Put(r, "flash", "News successfully created!")
	http.Redirect(w, r, fmt.Sprintf("/news/%s", category), http.StatusSeeOther)

}
func (app *application) staffNewsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/news/staff" {
		app.notFound(w)
		return
	}

	s, err := app.news.LatestByCategory("staff")

	if err != nil {
		log.Printf("Error fetching latest news: %s", err)
		app.serverError(w, err)
		return
	}

	log.Printf("Latest news retrieved: %+v", s)

	app.render(w, r, "staff.page.tmpl", &templateData{
		NewsList: s,
	})
}

func (app *application) applicantsNewsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/news/applicants" {
		app.notFound(w)
		return
	}

	s, err := app.news.LatestByCategory("applicants")
	if err != nil {
		log.Printf("Error fetching latest news: %s", err)
		app.serverError(w, err)
		return
	}

	log.Printf("Latest news retrieved: %+v", s)
	app.render(w, r, "applicants.page.tmpl", &templateData{
		NewsList: s,
	})
}

func (app *application) researchesNewsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/news/researches" {
		app.notFound(w)
		return
	}

	s, err := app.news.LatestByCategory("researches")
	if err != nil {

		app.serverError(w, err)
		return
	}

	log.Printf("Latest news retrieved: %+v", s)
	app.render(w, r, "researches.page.tmpl", &templateData{
		NewsList: s,
	})
}
func (app *application) studentsNewsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/news/students" {
		app.notFound(w)
		return
	}

	s, err := app.news.LatestByCategory("students")
	if err != nil {
		log.Printf("Error fetching latest news: %s", err)
		app.serverError(w, err)
		return
	}

	log.Printf("Latest news retrieved: %+v", s)
	app.render(w, r, "students.page.tmpl", &templateData{
		NewsList: s,
	})
}
func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "contact.page.tmpl", &templateData{})
}
