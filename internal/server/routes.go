package server

import (
	"encoding/json"
	"log"
	"net/http"

	"pulse/cmd/web"

	"github.com/a-h/templ"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	// mux.HandleFunc("/", s.HelloWorldHandler)
	mux.HandleFunc("/name", s.ParamHandler)

	mux.HandleFunc("/health", s.healthHandler)

	fileServer := http.FileServer(http.FS(web.Files))
	mux.Handle("/assets/", fileServer)
	mux.Handle("/", templ.Handler(web.HelloForm()))
	mux.HandleFunc("/hello", web.HelloWebHandler)
	mux.HandleFunc("/endpoints/upload", web.EndpointUploadWebHandler)

	return mux
}

// func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
// 	resp := make(map[string]string)
// 	resp["message"] = "Hello World"
//
// 	jsonResp, err := json.Marshal(resp)
// 	if err != nil {
// 		log.Fatalf("error handling JSON marshal. Err: %v", err)
// 	}
//
// 	_, _ = w.Write(jsonResp)
// }

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) ParamHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Param World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
