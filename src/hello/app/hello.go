package app

import (
	"fmt"
	"hello/templates"
	"net/http"
	"time"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	nowTime := time.Now()
	m := fmt.Sprintf("Hello, the time is %s", nowTime.String())

	t := templates.HelloTemplate(m)

	err := t.Render(r.Context(), w)

	if err != nil {
		http.Error(w, "failed to render", http.StatusInternalServerError)
	}
}
