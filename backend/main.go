// Copyright (C) 2026 Siddharth H
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

package main

import (
	
	// Internal Modules (Coded by me)
	
	"firegate/internal/git"
	"firegate/internal/services"
	"firegate/internal/auth"
	
	// External Modules
	
	"log"
	"net/http"
	"encoding/json"
	"fmt"	
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/alexedwards/scs/v2"
	"time"
)

type ApplyRequest struct {
	Message string `json:"message"`
}

var sm *scs.SessionManager

func main() {
	
	sm = scs.New()
	sm.Lifetime = 24 * time.Hour
	
	r := chi.NewRouter()
	r.Use(sm.LoadAndSave)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Login

	r.Post("/api/login", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON Body", 400)
			return
		}
		
		if err := auth.Verify(req.Username, req.Password); err != nil {
			http.Error(w, "Unauthorized", 401)
			return
		}	
		sm.Put(r.Context(), "authenticated", true)
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]bool{"success": true})
	})
	
	r.Group(func(r chi.Router) {
		
		r.Use(authMiddleware)
		
		// Get Git repo Tree status
		
		r.Get("/api/git/{repo}/status", func(w http.ResponseWriter, r *http.Request) {
			repo := chi.URLParam(r, "repo")
			status, err := git.GetStatus(repo)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			if status == "" {
				status = "working tree clean"
			}
			json.NewEncoder(w).Encode(map[string]string{"status": status})
		})
		
		// Restart a Service
		
		r.Get("/api/service/{service}/restart", func(w http.ResponseWriter, r *http.Request) {
			svc := chi.URLParam(r, "service")
			
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
			
			if  err := services.Restart(serviceType); err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": svc + " restarted with no errors"})
		})		
		// Apply changes
		
		r.Post("/api/service/{service}/apply", func(w http.ResponseWriter, r *http.Request) {
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
	
	})
	
	log.Println("Firegate Backend Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Middleware

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		// Check for Authentication
		isAuthenticated := sm.GetBool(r.Context(), "authenticated")
		
		if !isAuthenticated {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Access Denied. Authentication required."})
			return
		}
		
		next.ServeHTTP(w, r)		
	})	
}
