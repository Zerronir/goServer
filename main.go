package main

import (
	"fmt"
	"net/http"
	"html/template"
)
// Respondemos con archivos usando template
func index(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola mundo")
	template, err :=template.ParseFiles("./template/index.html")

	if err != nil {
		fmt.Fprintf(w, "Página no encontrada")
	} else {
		template.Execute(w,nil)
	}

}

func hola(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./template/hola.html")

	if err != nil {
		fmt.Fprintf(w, "Error 404")
	} else {
		template.Execute(w, nil)
	}

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hola", hola)
	fmt.Println("El servidor acepta solicitudes en el puerto 8080")
	http.ListenAndServe(":8080",nil)
}
