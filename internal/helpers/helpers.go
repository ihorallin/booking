package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"github.com/ihorallin/bookings/internal/config"
)

var app *config.AppConfig

// NewHelpers sets up new config for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

// ClientError 
func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status of", status)
	http.Error(w, http.StatusText(status), status)
}

// ServerError
func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Panicln(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}