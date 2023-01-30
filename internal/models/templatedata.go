package models

import "github.com/Ayan-Ansar/bookings/internal/forms"

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string // cross site request forgery token
	/* it's nothing more that when you build a Web page, you have a hidden with a form on it.
	You have a hidden field in that form, which is a long string of random numbers, and they change every single time somebody goes to the page
	*/
	Flash   string
	Warning string
	Error   string
	Form    *forms.Form
}
