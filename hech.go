package hech

import (
	"html/template"
	"log"
	"net/http"
)

func Check(detail string, err error) {
	if err != nil {
		log.Println(detail, " ", err)
	}
}

func Min(
	w http.ResponseWriter,
	pagePath string,
	data interface{}) error {

	tmpl, err := template.ParseFiles(pagePath)
	Check("Template Parsing Error: %v", err)

	err = tmpl.Execute(w, data)
	Check("Template Execution Error: %v", err)

	return nil
}
