package app

import (
	"container/templates"
	"net/http"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	uri := r.RequestURI

	if uri == "/" && uri == "" {
		uri = ""
	} else {
		uri = "/apps" + uri
	}

	t := templates.ContainerTemplate(uri)

	err := t.Render(r.Context(), w)

	if err != nil {
		http.Error(w, "failed to render", http.StatusInternalServerError)
	}
}
