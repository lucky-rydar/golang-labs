package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func RunHttpServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", PingHandler)

	mux.HandleFunc("/rooms", GetRoomsHandler)
	mux.HandleFunc("/rooms/add", AddRoomHandler)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	fmt.Println("Server is listening on port 8080")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
