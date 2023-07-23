package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/ihorallin/bookings/internal/config"
)

func TestRoots(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// Test passed
	default:
		t.Error(fmt.Sprintf("type is not h*chi.Mux, but is %T", v))
	}
}