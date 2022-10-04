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
	// Pasamos como parametros a la funcion encargada de realizar nuestro template nuestro index.html
	renderTemplate(w, "index.html")
}

func gameRounds(w http.ResponseWriter, r *http.Request) {
	// Obtenemos la informacion del frontend
	playerChoice := r.URL.Query().Get("p")
	indexChoice := r.URL.Query().Get("i")
	restart := r.URL.Query().Get("r")

	// Retornamos los tipos de datos necesarios en caso de que no necesitemos strings.
	restartConvert, _ := strconv.ParseBool(restart)
	indexConverted, _ := strconv.Atoi(indexChoice)

	// Guardamos en una variable y le pasamos los datos correspondientes a nuestra funcion
	// GameFunc encargada de centralizar la logica del tateti.
	result := board.GameFunc(playerChoice, indexConverted, restartConvert)

	// Convertimos en json lo que retornamos desde nuestra funcion GameFunc
	// podemos ver el json retornado en localhost:8080/gameRounds
	out, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}

func renderTemplate(w http.ResponseWriter, page string) {
	// Esta funcion la realizamos para realizar templates, debemos parsear nuestros
	// archivos html para que sean correctamente interpretdos.
	// al pasar nuestras pages por parametros nos simplifica si tenemos varias paginas para parsear.
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
	// Consume el directorio assets para poder integrar css VER
	fs := http.FileServer((http.Dir("assets")))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	// Establecemos controladores que toman request y responses como argumentos para interpretar html en estos casos.
	http.HandleFunc("/", homeGame)
	http.HandleFunc("/gameRounds", gameRounds)

	// Instanciamos nuestro servidor local en el puerto 8080
	http.ListenAndServe(":8080", nil)
}
