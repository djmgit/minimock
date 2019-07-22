package handlers

import (
	"net/http"
	"fmt"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	serverResponse := "Hellow world"
	fmt.Fprintf(w, serverResponse)
}
