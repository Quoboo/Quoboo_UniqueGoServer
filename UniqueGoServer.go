package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Anfrage erhalten!")
	fmt.Fprint(w, "Hallo, das ist meine sehr simple Go-Anwendung!")
}

func main() {
	http.HandleFunc("/meinprojekt", handler)
	fmt.Println("Server gestartet. Besuche http://localhost:8080/projekt im Browser.")
	http.ListenAndServe(":8080", nil)
}
