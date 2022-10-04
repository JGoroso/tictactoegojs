package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"tictactoe/board"
)

func homeGame(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "index.html")

}

func gameRounds(w http.ResponseWriter, r *http.Request) {

	playerChoice := r.URL.Query().Get("p")
	indexChoice := r.URL.Query().Get("i")
	restart := r.URL.Query().Get("r")

	restartConvert, _ := strconv.ParseBool(restart)
	indexConverted, _ := strconv.Atoi(indexChoice)

	result := board.GameFunc(playerChoice, indexConverted, restartConvert)

	out, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}

}

func main() {
	fs := http.FileServer((http.Dir("assets")))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	http.HandleFunc("/", homeGame)
	http.HandleFunc("/gameRounds", gameRounds)
	http.ListenAndServe(":8080", nil)
}
