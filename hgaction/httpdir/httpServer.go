package main

import (
	"fmt"
	"hgaction/routes"
	"log"
	"net/http"
)

func main() {
	fmt.Println("request start")
	routes.SetRouters()
	err := http.ListenAndServe("localhost:8082", nil)
	if err != nil {
		log.Fatal("listen error", err.Error())
	}
}
