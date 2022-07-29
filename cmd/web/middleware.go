package main

import (
	"net/http"
	"github.com/justinas/nosurf"
)


//NoSurf adds CSRF protection to all POST requests 
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	//uses cookie to ensure the token it generate is available on a per page bases
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}


//Session loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler{

	return session.LoadAndSave(next)
}