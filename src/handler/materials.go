package handler

import "net/http"
import . "./template"

func MaterialsHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "materials")
}
