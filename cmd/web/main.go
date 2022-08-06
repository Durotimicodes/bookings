package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Durotimicodes/bookings/pkg/config"
	"github.com/Durotimicodes/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":9092"

var session *scs.SessionManager

var app config.AppConfig

func main() {

	//change this to true when in production
	app.InProduction = false

	session = scs.New()                            //Instantiate a session
	session.Lifetime = 24 * time.Hour              //it should last 24hours
	session.Cookie.Persist = true                  // store session in cookie and make it persist
	session.Cookie.SameSite = http.SameSiteLaxMode //strict rules to define what site you want to apply this cookies/session to
	session.Cookie.Secure = app.InProduction       //ensures that the cookie is encypted i.e site is https not http

	app.Session = session

	//creating template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	//store the gotten template tc into the app template
	app.TemplateCache = tc
	app.UseCache = false

	//render the new template: gives the render access to the config variable
	render.NewTemplates(&app)

	//start a web server
	fmt.Printf("Starting application at port %s", portNumber)

	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
