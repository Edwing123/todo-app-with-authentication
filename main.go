package main

import (
	"log"
	"net/http"

	"github.com/Edwing123/todo-app-with-authentication/env"
	"github.com/Edwing123/todo-app-with-authentication/handlers"
	"github.com/Edwing123/todo-app-with-authentication/views"
)

func main() {
	views := views.CreateManager(templates)
	env := env.CreateEnv(views, nil)
	handlersFuncMap := handlers.CreateHandlerFuncsMap(env)
	serveMux := http.NewServeMux()

	// register handler functions
	for route, handlerFunc := range handlersFuncMap {
		serveMux.HandleFunc(route, handlerFunc)
	}

	// Set route for static assets
	serveMux.Handle("/assets/", createAssetsHandler())

	// run server
	addr := ":8080"

	log.Printf("Running on %s", addr)
	log.Fatalln(http.ListenAndServe(addr, serveMux))
}
