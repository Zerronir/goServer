
package main

import (
	"fmt"
	"net/http"
	"html/template"
	"encoding/json"
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

type User struct {
	Name string
	Age int
	Langs string
	Id int
}

func getUser(w http.ResponseWriter, r *http.Request) {

	method := "GET"
	userName := "Raul"
	age := 23
	langs := "Java, Python and PHP"
	userId := 1

	returnData := User{userName, age, langs, userId}
	if r.Method == method {

		jsonResp, err := json.Marshal(returnData)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResp)

	} else {
		http.Redirect(w, r, "/", 404)
	}

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hola", hola)
	http.HandleFunc("/api/getUser", getUser)
	fmt.Println("El servidor acepta solicitudes en el puerto 8080")
	http.ListenAndServe(":8080",nil)
}