package views

import (
	"dx/models"
	"html/template"
	"net/http"
)

type HomeIndexView struct {
	Name     string
	Machines []models.Machine
}

func RenderHomeIndex(w http.ResponseWriter, data HomeIndexView) error {
	name := "./views/home_index.html"
	tmpl, err := template.ParseFiles(name)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "text/html")
	return tmpl.Execute(w, data)
}
