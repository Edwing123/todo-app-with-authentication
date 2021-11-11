package env

import (
	"log"

	"github.com/Edwing123/todo-app-with-authentication/views"
)

type Env struct {
	Views  *views.Manager
	Logger *log.Logger
}

func CreateEnv(views *views.Manager, logger *log.Logger) *Env {
	return &Env{
		Views:  views,
		Logger: logger,
	}
}
