package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Ayan-Ansar/bookings/pkg/config"
	"github.com/Ayan-Ansar/bookings/pkg/handlers"
	"github.com/Ayan-Ansar/bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig

func main() {

	// change this to true when in production
	app.InProduction = false

	session := scs.New()                           // creating a new session
	session.Lifetime = 24 * time.Hour              // declare duration of session
	session.Cookie.Persist = true                  // true if you want cookies to stay after the user closes the browser window or quit the window, false otherwise
	session.Cookie.SameSite = http.SameSiteLaxMode // how strict you want to be about what site this cookie applies to
	session.Cookie.Secure = app.InProduction       // this will insist that the cookies be encrypted and that the connection is from HTTPS instead HTTP, in production you want this to be true

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	// to decide whether the code is in dev or production, if dev set false , else if prod set true
	app.UseCache = app.InProduction

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplate(&app)

	//http.HandleFunc("/", handlers.Home)
	//http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	// to start a local web server that listens to requests
	//_ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
