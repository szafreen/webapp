package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Initials struct {
	Name string
	Time string
}

func main() {
	initials := Initials{"Michal", time.Now().String()}

	templates := template.Must(template.ParseFiles("templates/template.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		if name != "" {
			initials.Name = name
		}

		if err := templates.ExecuteTemplate(w, "template.html", initials); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	fmt.Println("listening")
	fmt.Println(http.ListenAndServe(":8080", nil))


}



