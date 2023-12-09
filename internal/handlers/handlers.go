package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MariosKokmo/bookings/internal/config"
	"github.com/MariosKokmo/bookings/internal/models"
	"github.com/MariosKokmo/bookings/internal/render"
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
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
func (m *Repository) HomeTmpl(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	// every time I access the home page for the first time, I store the
	// user IP in the session
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) SingleRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "singleroom.page.tmpl", &models.TemplateData{})
}

func (m *Repository) DoubleRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "doubleroom.page.tmpl", &models.TemplateData{})
}

func (m *Repository) CheckDates(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "checkdates.page.tmpl", &models.TemplateData{})
}

// Post CheckDates
func (m *Repository) PostCheckDates(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
}

type jsonResponse struct {
	OK bool `json:"ok"`
	Message string `json:"message"`
}
// Handles request for check dates and returns JSON response
func (m *Repository) CheckDatesJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil{
		log.Println(err)
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "makereservation.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}