package main

import (
	"net/http"
	"time"
	"fmt"
	"log"

	"github.com/it-02/dormitory/db"
	"github.com/it-02/dormitory/handlers"
)

func RunHttpServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", handlers.PingHandler)

	mux.HandleFunc("/rooms", handlers.GetRoomsHandler)
	mux.HandleFunc("/rooms/add", handlers.AddRoomHandler)
	mux.HandleFunc("/rooms/byPlaceId", handlers.GetRoomByPlaceIdHandler)
	mux.HandleFunc("/rooms/stats/byNumber", handlers.GetRoomStatsByNumberHandler)
	
	mux.HandleFunc("/places", handlers.GetPlacesHandler)
	mux.HandleFunc("/places/free", handlers.GetFreePlacesHandler)
	mux.HandleFunc("/places/free/byRoomId", handlers.GetFreePlacesByRoomIdHandler)
	mux.HandleFunc("/places/byRoomId", handlers.GetPlacesByRoomIdHandler)

	mux.HandleFunc("/contracts", handlers.GetContractsHandler)
	mux.HandleFunc("/contracts/add", handlers.AddContractHandler)

	mux.HandleFunc("/student_tickets", handlers.GetStudentTicketsHandler)
	mux.HandleFunc("/student_tickets/add", handlers.AddStudentTicketHandler)
	mux.HandleFunc("/student_tickets/bySerialNumber", handlers.GetStudentTicketBySerialNumberHandler)

	mux.HandleFunc("/students", handlers.GetStudentsHandler)

	mux.HandleFunc("/ask_admin/register", handlers.AskAdminRegisterHandler)
	mux.HandleFunc("/ask_admin/contract/sign", handlers.AskAdminSignContractHandler)
	mux.HandleFunc("/ask_admin/settle", handlers.AskAdminSettleHandler)
	mux.HandleFunc("/ask_admin/unsettle", handlers.AskAdminUnsettleHandler)
	mux.HandleFunc("/ask_admin/resettle", handlers.AskAdminResettleHandler)
	mux.HandleFunc("/ask_admin/actions", handlers.GetActionsHandler)
	// only for admin
	mux.HandleFunc("/ask_admin/actions/resolve", handlers.ResolveActionHandler)

	mux.HandleFunc("/dormitory/load/stats", handlers.GetDormitoryLoadHandler)

	mux.HandleFunc("/user/register", handlers.RegisterUserHandler)
	mux.HandleFunc("/user/login", handlers.LoginUserHandler)

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

func main() {
	db.SetupDB()
	RunHttpServer()
}
