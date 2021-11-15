package handlers

import (
	"net/http"

	"github.com/Edwing123/todo-app-with-authentication/env"
	"github.com/Edwing123/todo-app-with-authentication/handlers/index"
)

// Used to map routes to handler functions
type handlersFuncMap map[string]func(w http.ResponseWriter, r *http.Request)

// Creates a handlers map where keys are paths (e.g. "/"), and values
// are the corresponding paths handler functions.
func CreateHandlerFuncsMap(env *env.Env) handlersFuncMap {
	return handlersFuncMap{
		"/": index.Get(env),
	}
}
