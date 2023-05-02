package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi"
	"github.com/rafteodoro/bookings/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		//do nothing
	default:
		t.Error(fmt.Sprintf("Type is not chi.mux, type is %T", v))
	}
}
