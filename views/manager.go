package views

import (
	"bytes"
	"html/template"
	"io/fs"
)

// Manager is a wrapper struct for *template.Template
type Manager struct {
	t *template.Template // Tem
}

// Executes template corresponding to the given name and with the gicen ViewContext,
// if everything is fine it returns a byte slice with the executed template
// and a nil error, otherwise it will return a nil byte slice and an error
// indicating what failed in the execution.
func (m *Manager) ExecuteView(name string, viewContext *ViewContext) ([]byte, error) {
	var output bytes.Buffer
	err := m.t.ExecuteTemplate(&output, name, viewContext)

	if err != nil {
		return []byte{}, err
	}

	return output.Bytes(), nil
}

func CreateManager(fsys fs.FS) *Manager {
	manager := Manager{
		t: template.Must(template.New("templates").ParseFS(fsys, "src/templates/**/*.go.html")),
	}
	return &manager
}
