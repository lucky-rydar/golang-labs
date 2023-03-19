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
	mux.HandleFunc("/rooms/byPlaceId", GetRoomByPlaceIdHandler)
	
	mux.HandleFunc("/places", GetPlacesHandler)
	mux.HandleFunc("/places/free", GetFreePlacesHandler)
	mux.HandleFunc("/places/free/byRoomId", GetFreePlacesByRoomIdHandler)
	mux.HandleFunc("/places/byRoomId", GetPlacesByRoomIdHandler)

	mux.HandleFunc("/contracts", GetContractsHandler)
	mux.HandleFunc("/contracts/add", AddContractHandler)

	mux.HandleFunc("/student_tickets", GetStudentTicketsHandler)
	mux.HandleFunc("/student_tickets/add", AddStudentTicketHandler)
	mux.HandleFunc("/student_tickets/bySerialNumber", GetStudentTicketBySerialNumberHandler)

	mux.HandleFunc("/student/register", RegisterStudentHandler)
	mux.HandleFunc("/student/contract/sign", SignContractHandler)
	mux.HandleFunc("/student/settle", SettleHandler)
	mux.HandleFunc("/student/unsettle", UnsettleHandler)
	mux.HandleFunc("/student/resettle", ResettleHandler)
	mux.HandleFunc("/students", GetStudentsHandler)

	mux.HandleFunc("/dormitory/load/stats", GetDormitoryLoadHandler)

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
