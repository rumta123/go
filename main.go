package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Title struct {
	Name string
	Desc string
}

func (u Title) getAllInfo() string {
	return fmt.Sprintf(" %s.", u.Name, u.Desc)
}
func (u *Title) setNewName(newName string, newDesc string) {
	u.Name = newName
	u.Desc = newDesc
}
func index(w http.ResponseWriter, r *http.Request) {
	bob := Title{"Страница", "Описание для главной"}
	bob.setNewName("Страница главная", "Описание для главной")
	// fmt.Fprint(w, bob.getAllInfo())
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", bob)

}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	bob := Title{"Страница", "Описание для контактов"}
	bob.setNewName("Контакты", "Описание Контакты")

	// fmt.Fprint(w, bob.getAllInfo())
	t, err := template.ParseFiles("templates/contacts.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "contacts_page", bob)

}

func form_page(w http.ResponseWriter, r *http.Request) {
	bob := Title{"Форма отправки", "Эта форма обработана с помощью golang"}
	// bob.setNewName("Контакты", "Описание Контакт")

	// fmt.Fprint(w, bob.getAllInfo())
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "form_page", bob)

}

func handleFunc() {
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/contacts/", contacts_page)
	http.HandleFunc("/create/", form_page)
	http.ListenAndServe((":8080"), nil)

}
func main() {
	handleFunc()

}
