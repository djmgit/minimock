package server

import (
	"minimock/handlers"
	"net/http",
	"strconv"
)

func Server(port int) {
	http.HandleFunc("/", handlers.Handler)
	http.ListenAndServe(":" + strconv.FormatInt(port, 10), nil)
}
