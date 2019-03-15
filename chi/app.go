/**
 * 轻量级的http api框架
 * https://github.com/go-chi/chi
 */
package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	log.Println("server has run on 3000")
	http.ListenAndServe(":3000", r)
}
