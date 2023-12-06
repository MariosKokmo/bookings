package handlers

import (
	"net/http"

	"github.com/MariosKokmo/bookings/pkg/config"
	"github.com/MariosKokmo/bookings/pkg/models"
	"github.com/MariosKokmo/bookings/pkg/render"
)

// Repository is used by Handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) AboutTmpl(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	stringMap["text"] = `Lorem Ipsum
	This is the best text that you'll find here.`

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
func (m *Repository) HomeTmpl(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	// every time I access the home page for the first time, I store the
	// user IP in the session
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) SingleRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "singleroom.page.tmpl", &models.TemplateData{})
}

func (m *Repository) DoubleRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "doubleroom.page.tmpl", &models.TemplateData{})
}

func (m *Repository) CheckDates(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "checkdates.page.tmpl", &models.TemplateData{})
}

func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "makereservation.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}