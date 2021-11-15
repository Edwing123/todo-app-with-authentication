package index

import (
	"fmt"
	"net/http"

	"github.com/Edwing123/todo-app-with-authentication/env"
	"github.com/Edwing123/todo-app-with-authentication/views"
)

var (
	defaultIndexPage = views.Page{
		Meta: views.Meta{
			Title:       "Home | UrTodos",
			Description: "Web app to manage your todos",
		},
	}

	defaultIndexData = map[string]interface{}{
		"IsAuthenticated": false,
	}
)

func createIndexViewContext() *views.ViewContext {
	context := views.ViewContext{
		Page: defaultIndexPage,
		Data: defaultIndexData,
	}

	return &context
}

func Get(env *env.Env) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		context := createIndexViewContext()
		view, err := env.Views.ExecuteView("index.go.html", context)

		if err != nil {
			http.Error(
				w,
				fmt.Sprintf("failed at executing template with error %s", err),
				http.StatusInternalServerError,
			)
			return
		}

		w.Write(view)
	}
}
