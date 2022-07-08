package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

type Game struct {
	Players []Player
}

type Player struct {
	Name   string
	Hurt   bool
	Health int
	Weapon string
}

type PageData struct {
	Title string
	Game  Game
}

func game(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Testing htmlGame",
		Game: Game{
			[]Player{
				{Name: "John", Hurt: false, Health: 100, Weapon: " "},
				{Name: "Bill", Hurt: true, Health: 90, Weapon: " "},
			},
		},
	}
	tmpl.Execute(w, data)
}

func main() {
	fmt.Println("Init Html game")

	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))

	fs := http.FileServer(http.Dir("/.static"))
	mux.Handle("/static", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/game", game)

	log.Fatal(http.ListenAndServe(":9090", mux))
}
