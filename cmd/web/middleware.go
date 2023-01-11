package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

/*func WriteToConsole(next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r) // move to the next part in code
	})
}*/

// NoSurf: Adds CSRF protection to every POST request 
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	/* We need to set the base cookie because it uses cookies to make sure that the token it generates is available
	on a per page basis.*/

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad: Load and Save the session on every request 
func SessionLoad(next http.Handler) http.Handler {
	return app.Session.LoadAndSave(next)
}
