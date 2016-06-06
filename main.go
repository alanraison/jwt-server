package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SermoDigital/jose/jws"
)

func handle(w http.ResponseWriter, r *http.Request) {
  log.Println("Handling Request")
	j, err := jws.ParseJWTFromRequest(r)

	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "https:///www.google.com", http.StatusFound)
    return
	}

	fmt.Printf("Valid token issued to %s\n", j.Claims().Get("name"))
  fmt.Fprintf(w, "Token validated!\n")
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":9090", nil)
	fmt.Println("Listening on port 9090")
}
