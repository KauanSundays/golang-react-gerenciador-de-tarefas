package main

import (
	"fmt"
	"log"
	"net/http"
	"golang-react-gerenciador-de-tarefas/server/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting the server on port 9000...")

	log.Fatal(http.ListenAndServe(":9000", r))
}
