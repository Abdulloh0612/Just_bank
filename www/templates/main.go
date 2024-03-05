package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                 string
	Age                  int
	Money                int
	Avg_grades, Happines float64
	Hobbies              []string
}

func (u *User) getAllinfo() string {
	return fmt.Sprintf("User name id: %s, He is %d and he has money"+
		"equal: %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(NewName string) {
	u.Name = NewName
}

func home_page(page http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8, []string{"Football", "Skate", "Dance"}}
	tmpl, err := template.ParseFiles("www/templates/home_page.html")
	if err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(page, bob)
}

func contacts_page(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "Contacts page!")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
