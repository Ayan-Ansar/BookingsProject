package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Ayan-Ansar/bookings/internal/config"
	"github.com/Ayan-Ansar/bookings/internal/forms"
	"github.com/Ayan-Ansar/bookings/internal/models"
	"github.com/Ayan-Ansar/bookings/internal/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}


// Home renders the Landing/Home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP) // storing the users remote ip as a string in the session
	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}


// About renders the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "HELLO WORLD"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation renders the reservation page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "reservations.page.html", &models.TemplateData{
		Form: forms.New(nil),
	})
}

// PostReservation handles the posting of a reservation form 
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {

}

//Suite renders the suite page
func (m *Repository) Suite(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "suite.page.html", &models.TemplateData{})
}

// BasicRoom renders the basic room page
func (m *Repository) BasicRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "basicroom.page.html", &models.TemplateData{})
}

// Availablility render the book now page 
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.html", &models.TemplateData{})
}

// contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{})
}

// PostAvailability posts the start and end received from the book now page 
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles requests for availability and sends json response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {

	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
