package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

var domainMappings map[string]string

func loadConfig() {
	data, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	err = json.Unmarshal(data, &domainMappings)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}
}

func main() {
	loadConfig()
	log.Println("Loaded config.json file")

	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			// Use the domain mapping from the config
			if newDomain, ok := domainMappings[req.Host]; ok {
				log.Printf("Forwarding request form %s to %s", req.Host, newDomain)
				req.URL.Host = newDomain
				req.URL.Scheme = "http"
				req.Host = newDomain
			} else {
				log.Printf("No mapping found for domain: %s", req.Host)
			}
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	log.Println("Starting DomainForwarder (proxy) server on port 80")
	log.Fatal(http.ListenAndServe(":80", nil))
}
