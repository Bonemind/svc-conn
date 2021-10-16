package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
	"strings"
)

type ServiceResult struct {
	Name   string
	Result bool
}

func main() {
	http.HandleFunc("/", ServiceStateHandler)
	fmt.Println("Listening on port 7887")
	http.ListenAndServe(":7887", nil)
}

func ServiceStateHandler(w http.ResponseWriter, r *http.Request) {
	domainString := strings.TrimSpace(os.Getenv("SERVICE_DOMAINS"))

	var domains []string
	if domainString != "" {
		domains = strings.Split(domainString, ",")
	}

	var results []ServiceResult

	for _, d := range domains {
		_, err := http.Get(d)
		res := ServiceResult{Name: d, Result: true}
		if err != nil {
			res.Result = false
		}
		results = append(results, res)
	}

	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, results); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
