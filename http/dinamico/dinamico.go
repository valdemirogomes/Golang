package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func hora_certa(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "<h1>Hora atual %s<h1>", s)
}
func main() {
	http.HandleFunc("/hora/", hora_certa)
	log.Println("Servidor rodando")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
