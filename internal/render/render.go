package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/MariosKokmo/bookings/internal/config"
	"github.com/MariosKokmo/bookings/internal/models"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig
var pathToTemplates = "./templates"
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) error{
	// create a template cache or better get it from app config
	// so that I do not create it over and over again
	var tc map[string]*template.Template
	// using cache might mean that I don't reload the changes from disk
	// while the server is running. I would probably need that
	// in development mode
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	// get requested template from template cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
		return errors.New("can't get template from cache")
	}

	// buffer will just hold bytes
	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	_ = t.Execute(buf, td)

	// render template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	//myCache := make(map[string]*template.Template) or the same as
	myCache := map[string]*template.Template{}
	// populate all at once named *.page.tmpl
	// get all of the files that end in *.page.tmpl from ./templates
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return myCache, err
	}

	// range through all pages
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// look if any layouts exist in the directory
		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			fmt.Println("No templates found")
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
