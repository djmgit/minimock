package server

import (
	"minimock/handlers"
	"net/http"
)

func Server(port string) {
	http.HandleFunc("/", handlers.Handler)
	http.HandleFunc("/sleep", handlers.HandlerSleep)
	http.ListenAndServe(":" + port, nil)
}
