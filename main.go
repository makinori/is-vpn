package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/makinori/is-vpn/services"
)

var (
	//go:embed template.html public/*
	STATIC_CONTENT embed.FS

	SERVICE = getEnv("SERVICE", "")
	PORT, _ = strconv.Atoi(getEnv("PORT", "8080"))
)

func getEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

func httpErrorJson(w http.ResponseWriter, err string, statusCode int) {
	data, _ := json.Marshal(map[string]string{
		"error": err,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}

func main() {
	if !slices.Contains(services.SERVICE_LIST, SERVICE) {
		panic("service not found")
	}

	statusResolveFunc, err := services.GetStatusResolveFunc(SERVICE)
	if err != nil {
		panic(err)
	}

	htmlTemplate, err := STATIC_CONTENT.ReadFile("template.html")
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("template").Parse(string(htmlTemplate))
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api" {
			status, err := statusResolveFunc()
			if err != nil {
				httpErrorJson(w, err.Error(), http.StatusInternalServerError)
				log.Println(err.Error())
				return
			}

			data, err := json.Marshal(status)
			if err != nil {
				httpErrorJson(w, err.Error(), http.StatusInternalServerError)
				log.Println(err.Error())
				return
			}

			w.Header().Add("Content-Type", "application/json")
			w.Write(data)

			return
		}

		if r.URL.Path == "/" {
			if err != nil {
				httpErrorJson(w, err.Error(), http.StatusInternalServerError)
				log.Println(err.Error())
				return
			}

			status, err := statusResolveFunc()
			if err != nil {
				httpErrorJson(w, err.Error(), http.StatusInternalServerError)
				log.Println(err.Error())
				return
			}

			var bytes bytes.Buffer
			tmpl.Execute(&bytes, status)

			w.Write(bytes.Bytes())

			return
		}

		if strings.HasPrefix(r.URL.Path, "/public") {
			http.ServeFileFS(w, r, STATIC_CONTENT, r.URL.Path)
			return
		}

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	})

	log.Printf(
		"starting web server: http://127.0.0.1:%d\n",
		PORT,
	)

	err = http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
	if err != nil {
		log.Fatalf("error starting server: %s", err)
	}
}
