package handlers

import (
	"github.com/Ayan-Ansar/bookings/pkg/config"
	"github.com/Ayan-Ansar/bookings/pkg/models"
	"github.com/Ayan-Ansar/bookings/pkg/render"
	"net/http"
)

var Repo *Repository 

type Repository struct{
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository){
	Repo = r
}


func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP) // storing the users remote ip as a string in the session 
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "HELLO WORLD"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}