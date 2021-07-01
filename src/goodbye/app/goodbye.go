package app

import (
	"fmt"
	"goodbye/templates"
	"net/http"
	"time"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	nowTime := time.Now()
	m := fmt.Sprintf("Goodbye, the time is %s", nowTime.String())

	t := templates.GoodbyeTemplate(m)

	err := t.Render(r.Context(), w)

	if err != nil {
		http.Error(w, http.StatusInternalServerError, "failed to render")
	}
}
