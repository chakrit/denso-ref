package controllers

import (
	"dx/models"
	"dx/views"
	"log"
	"net/http"
)

type HomeController struct{}

func (h HomeController) Index(w http.ResponseWriter, r *http.Request) {
	machines, err := models.GetAllMachines()
	if err != nil {
		log.Println("error loading machines", err)
	}

	err = views.RenderHomeIndex(w, views.HomeIndexView{
		Name:     "A4",
		Machines: machines,
	})
	if err != nil {
		log.Println("error rendering views", err)
	}
}

func (h HomeController) New(w http.ResponseWriter, r *http.Request) {
	err := views.RenderHomeNew(w)
	if err != nil {
		log.Println("error rendering views", err)
	}
}

func (h HomeController) Create(w http.ResponseWriter, r *http.Request) {
	err := models.AddNewMachine(models.Machine{
		ID:     r.PostFormValue("ID"),
		Name:   r.PostFormValue("Name"),
		Status: r.PostFormValue("Status"),
	})
	if err != nil {
		log.Println("error retrieving form data", err)
	}

	http.Redirect(w, r, "/", 302)
	// ...
}
