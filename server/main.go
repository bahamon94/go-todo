package main

import (
	"./router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Route()
	fmt.Println("Iniciando server in port 9090...")
	log.Fatal(http.ListenAndServe(":9090", r))
}
