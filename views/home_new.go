package views

import (
	"html/template"
	"net/http"
)

func RenderHomeNew(w http.ResponseWriter) error {
	name := "./views/home_new.html"
	tmpl, err := template.ParseFiles(name)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "text/html")
	return tmpl.Execute(w, nil)
}
