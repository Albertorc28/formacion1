package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Peticion struct {
	ID      int
	Palabra string
	Fecha   time.Time
}

func main() {
	var puerto = 8080
	fmt.Println("Iniciando servidor...", puerto)
	http.HandleFunc("/", iniciohandle)
	http.HandleFunc("/js/", jshandle)
	http.HandleFunc("/ejercicio", ejercicio)
	http.ListenAndServe(":"+strconv.Itoa(puerto), nil)
	/*js := http.FileServer(http.Dir("js/"))
	http.Handle("/js/", http.StripPrefix("/js/", js))*/
}

func iniciohandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recibiendo peticion desde " + r.URL.EscapedPath())
	if r.URL.EscapedPath() != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "pages/index.html")

}

func jshandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recibiendo petición desde " + r.URL.EscapedPath())
	/*if r.URL.EscapedPath() != "/js/" {
		http.NotFound(w, r)
		return
	}*/

	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, "js/base.js")
}

func ejercicio(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recibiendo peticion desde " + r.URL.EscapedPath())

	if r.URL.Path != "/ejercicio" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	defer r.Body.Close()
	var bytes, e = ioutil.ReadAll(r.Body)
	if e == nil {
		var peticion Peticion
		json.Unmarshal(bytes, &peticion)
		if peticion.Palabra != "" {
			peticion.Palabra = strings.ToUpper(peticion.Palabra)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, e)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		respuesta, _ := json.Marshal(peticion)
		fmt.Fprint(w, string(respuesta))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, e)
	}
}
