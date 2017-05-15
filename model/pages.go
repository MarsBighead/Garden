package model

import (
	"html/template"
	"log"
	"net/http"
)

func HomeList(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(GetCurrentDir() + "/template/list.htm")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}

// Home  home/index web page
func Home(w http.ResponseWriter, r *http.Request) {
	dumx := Person{
		Name:    "zoro",
		Age:     27,
		Emails:  []string{"dg@gmail.com", "dk@hotmail.com"},
		Company: "Omron",
		Role:    "SE"}

	chxd := Person{
		Name:   "chx",
		Age:    26,
		Emails: []string{"test@gmail.com", "d@hotmail.com"}}

	onlineUser := OnlineUser{User: []*Person{&dumx, &chxd}}

	log.Print("Running http handle modle.Home!")
	t, err := template.ParseFiles(GetCurrentDir() + "/template/home.htm")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, onlineUser)
	if err != nil {
		log.Fatal(err)
	}
}
