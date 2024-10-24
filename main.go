package main

import (
	routes "CRUD/Routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello 🥹")
	r := routes.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("listening at port:4000")
}
