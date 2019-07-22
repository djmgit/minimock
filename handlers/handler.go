package handlers

import (
	"net/http"
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	serverResponse := "Hello world Hello world Hello world"
	fmt.Fprintf(w, serverResponse)
}

func HandlerSleep(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	randomWait := rand.Intn(60)
	time.Sleep(time.Duration(randomWait) * time.Second)
	serverResponse := "Hello world Hello world Hello worldHello world Hello world Hello world Hello world Hello world Hello world Hello world Hello world Hello world Hello world Hello world Hello world Hello world Hello world Hello world" + strconv.FormatInt(int64(randomWait), 10)

	fmt.Fprintf(w, serverResponse)
}
