package app

import (
	"container/templates"
	"fmt"
	"net/http"
	"time"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	nowTime := time.Now()
	m := fmt.Sprintf("Container, the time is %s", nowTime.String())

	t := templates.ContainerTemplate(m)

	err := t.Render(r.Context(), w)

	if err != nil {
		http.Error(w, "failed to render", http.StatusInternalServerError)
	}
}
