package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	generator "password/generator"
)

type pass struct {
	Result string
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./template/index.html")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tmpl)
	var data pass
	data.Result = generator.Generator()

	err = tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		fmt.Println(err)
	}
}
