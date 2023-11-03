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

func X(
	w http.ResponseWriter,
	starting_point string,
	funcMap template.FuncMap,
	layouts []string,
	components []string,
	pages []string,
	data interface{}) error {

	var filesString []string
	filesString = append(filesString, layouts...)
	filesString = append(filesString, components...)
	filesString = append(filesString, pages...)

	tmpl, err := template.New("").Funcs(funcMap).ParseFiles(filesString[:]...)
	Check("Template Parsing Error: %v", err)

	err = tmpl.ExecuteTemplate(w, starting_point, data)
	Check("Template Execution Error: %v", err)

	return nil
}
