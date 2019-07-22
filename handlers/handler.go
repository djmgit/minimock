package handlers

import (
	"net/http"
	"fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
	serverResponse := "Hellow world"
	fmt.Fprintf(w, serverResponse)
}
