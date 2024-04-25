package web

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Health struct {
	Server string `json:"server"`
}

func HelloWebHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	endpoint := r.FormValue("endpoint")

	resp, err := http.Get(endpoint)

	if err != nil {
		log.Fatalf("Error fetching in HelloWebHandler: %e", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error readAll: %e", err)
	}

	var h Health
	err = json.Unmarshal(body, &h)
	if err != nil {
		log.Fatalf("Error failed unmarshal: %e", err)
	}

	component := HealthPost(h)
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in HelloWebHandler: %e", err)
	}
}
