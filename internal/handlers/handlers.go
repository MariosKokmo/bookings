package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MariosKokmo/bookings/internal/config"
	"github.com/MariosKokmo/bookings/internal/forms"
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
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.RenderTemplate(w, r, "makereservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// PostReservation handles the posting of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName: r.Form.Get("last_name"),
		Email: r.Form.Get("email"),
	}

	form := forms.New(r.PostForm)

	
	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3, r)
	form.IsEmail("email")
	
	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(w, r, "makereservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	// route to reservation summary
	// first we store the data into the session
	m.App.Session.Put(r.Context(), "reservation", reservation)

	// redirect to summary page
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		log.Println("cannot get item from session")
		m.App.Session.Put(r.Context(), "error", "Can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	// at this point we have taken the reservation, don't need it anymore
	m.App.Session.Remove(r.Context(), "reservation")
	
	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.RenderTemplate(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data: data,
	})
}