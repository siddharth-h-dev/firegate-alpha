package main

import (
	"log"
	"net/http"
	"encoding/json"
	"fmt"	
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"firegate/internal/git"
	"firegate/internal/services"
)

type ApplyRequest struct {
	Message string `json:"message"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	
	r.Get("/api/{service}/status", func(w http.ResponseWriter, r *http.Request) {
		svc := chi.URLParam(r, "service")
		status, err := git.GetStatus(svc)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		if status == "" {
			status = "working tree clean"
		}
		json.NewEncoder(w).Encode(map[string]string{"status": status})
	})

	r.Post("/api/{service}/apply", func(w http.ResponseWriter, r *http.Request) {
		svc := chi.URLParam(r, "service")
		
		var req ApplyRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err !=nil {
			http.Error(w, "Invalid JSON Body", 400)
			return
		}
		
		if req.Message == "" {
			req.Message = fmt.Sprintf("Updated %s config in UI", svc)
		}
		
		var serviceType services.ServiceType
		switch svc {
			case  "nftables": serviceType = services.ServiceNftables
			case  "unbound": serviceType = services.ServiceUnbound
			case  "suricata": serviceType = services.ServiceSuricata
			case  "tor": serviceType = services.ServiceTor
			case "motd": serviceType = services.ServiceMOTD
			default:
				http.Error(w, "Unknown Service", 400)
				return
		}
		
		if  err := services.ApplyConfig(serviceType, req.Message); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": svc + " config applied and commited"})
	})

	log.Println("Firegate Backend Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
