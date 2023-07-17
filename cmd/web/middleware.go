package main

import (
	"net/http"
	"github.com/justinas/nosurf"
)

// NoSurf adds CSRF protection to all POST requet
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad loads andsaves the session on very request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}