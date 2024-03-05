package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web_site/index.html", "web_site/header.html", "web_site/footer.html")
	if err != nil {
		fmt.Println("Ошибка загрузки шаблонов:", err)
		return
	}

	err = t.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		fmt.Println("Ошибка выполнения шаблона:", err)
		return
	}
}

func handFunc() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handFunc()
}
