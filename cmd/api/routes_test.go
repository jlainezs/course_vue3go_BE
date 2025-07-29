package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"testing"
)

func Test_Routes_Exist(t *testing.T) {
	testRoutes := testApp.routes()
	chiRoutes := testRoutes.(chi.Routes)

	// these routes must exist
	routeExists(t, chiRoutes, "/users/login")
	routeExists(t, chiRoutes, "/users/logout")
	routeExists(t, chiRoutes, "/admin/users/get/{id}")
	routeExists(t, chiRoutes, "/admin/users/save")
	routeExists(t, chiRoutes, "/admin/users")
	routeExists(t, chiRoutes, "/admin/users/delete")
}

func routeExists(t *testing.T, routes chi.Routes, route string) {
	// assume that the route does not exists
	found := false

	// walk through all the registered routes
	_ = chi.Walk(routes, func(method string, foundRoute string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		// if we find the route we're looking for, set found to true
		if route == foundRoute {
			found = true
		}
		return nil
	})

	if !found {
		t.Errorf("Route %s does not exist", route)
	}
}
