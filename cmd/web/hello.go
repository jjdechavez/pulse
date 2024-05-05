package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Health struct {
	Server string `json:"server"`
}

var (
	TEN_MEGABYTES_MAX_SIZE_UPLOAD = 10 << 20
)

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

	var res map[string]interface{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error failed unmarshal: %e", err)
	}

	component := HealthPost(res)
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in HelloWebHandler: %e", err)
	}
}

type EndpointFile struct {
	Endpoint string
}

func EndpointUploadWebHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(int64(TEN_MEGABYTES_MAX_SIZE_UPLOAD))

	file, header, err := r.FormFile("endpoints")
	if err != nil {
		log.Fatalf("Error parsing file: %e", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	defer file.Close()

	jsonFile, err := header.Open()
	if err != nil {
		log.Fatalf("Failed to open the uploaded file: %e", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	content, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Failed to read the uploaded file: %e", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	fmt.Print(string(content))

	var listOfEndpoints []EndpointFile
	errU := json.Unmarshal(content, &listOfEndpoints)
	if errU != nil {
		log.Fatalf("Failed to json econde the uploaded file content: %e", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	fmt.Printf("All listOfEndpoints: %+v", listOfEndpoints)
	// fmt.Printf("total count of listOfEndpoints: %d", len(listOfEndpoints))

	for _, value := range listOfEndpoints {
		fmt.Printf("each endpoint: %+v", value.Endpoint)
	}

	component := InitialResultPost(listOfEndpoints)
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in EndpointUploadWebHandler: %e", err)
	}
}
