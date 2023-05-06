package main

import (
	"net/http"
	"time"
	"fmt"
	"log"

	"github.com/it-02/dormitory/internals/db"
	handlers "github.com/it-02/dormitory/internals/handlers"
	
	ask_admin_repository "github.com/it-02/dormitory/internals/features/ask_admin/repository"
	ask_admin_service "github.com/it-02/dormitory/internals/features/ask_admin/service"

	contract_repository "github.com/it-02/dormitory/internals/features/contract/repository"
	contract_service "github.com/it-02/dormitory/internals/features/contract/service"

	dormitory_load_service "github.com/it-02/dormitory/internals/features/dorm_stats/service"

	place_repository "github.com/it-02/dormitory/internals/features/place/repository"
	place_service "github.com/it-02/dormitory/internals/features/place/service"

	room_repository "github.com/it-02/dormitory/internals/features/room/repository"
	room_service "github.com/it-02/dormitory/internals/features/room/service"

	student_repository "github.com/it-02/dormitory/internals/features/student/repository"
	student_service "github.com/it-02/dormitory/internals/features/student/service"

	student_ticket_repository "github.com/it-02/dormitory/internals/features/student_ticket/repository"
	student_ticket_service "github.com/it-02/dormitory/internals/features/student_ticket/service"

	user_repository "github.com/it-02/dormitory/internals/features/user/repository"
	user_service "github.com/it-02/dormitory/internals/features/user/service"

	"gorm.io/gorm"
)

func RunHttpServer(db *gorm.DB) {
	// init repositories
	room_repository := room_repository.NewRoom(db)
	place_repository := place_repository.NewPlace(db)
	contract_repository := contract_repository.NewContract(db)
	student_ticket_repository := student_ticket_repository.NewStudentTicket(db)
	student_repository := student_repository.NewStudent(db, place_repository)
	user_repository := user_repository.NewUser(db)
	ask_admin_repository := ask_admin_repository.NewAskAdmin(db)

	// init services
	user_service := user_service.NewUserService(user_repository)
	student_service := student_service.NewStudentService(student_repository, student_ticket_repository, room_repository, place_repository, contract_repository, user_service)
	ask_admin_service := ask_admin_service.NewAskAdminService(ask_admin_repository, user_service, student_service)
	place_service := place_service.NewPlaceService(place_repository)
	room_service := room_service.NewRoomService(room_repository, place_repository, student_repository, user_service)
	contract_service := contract_service.NewContractService(contract_repository, user_service)
	student_ticket_service := student_ticket_service.NewStudentTicketService(student_ticket_repository)
	dormitory_load_service := dormitory_load_service.NewDormitoryLoadService(place_repository, room_repository, user_service)

	// init controllers
	ask_admin_controller := handlers.NewAskAdminController(ask_admin_service)
	contract_controller := handlers.NewContractController(contract_service)
	dormitory_load_controller := handlers.NewDormitoryLoadController(dormitory_load_service)
	place_controller := handlers.NewPlaceController(place_service)
	room_controller := handlers.NewRoomController(room_service)
	student_controller := handlers.NewStudentController(student_service)
	student_ticket_controller := handlers.NewStudentTicketController(student_ticket_service)
	user_controller := handlers.NewUserController(user_service)


	// setup http service
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", handlers.PingHandler)

	mux.HandleFunc("/rooms", room_controller.GetRoomsHandler)
	mux.HandleFunc("/rooms/add", room_controller.AddRoomHandler)
	mux.HandleFunc("/rooms/byPlaceId", room_controller.GetRoomByPlaceIdHandler)
	mux.HandleFunc("/rooms/stats/byNumber", room_controller.GetRoomStatsByNumberHandler)
	
	mux.HandleFunc("/places", place_controller.GetPlacesHandler)
	mux.HandleFunc("/places/free", place_controller.GetFreePlacesHandler)
	mux.HandleFunc("/places/free/byRoomId", place_controller.GetFreePlacesByRoomIdHandler)
	mux.HandleFunc("/places/byRoomId", place_controller.GetPlacesByRoomIdHandler)

	mux.HandleFunc("/contracts", contract_controller.GetContractsHandler)
	mux.HandleFunc("/contracts/add", contract_controller.AddContractHandler)

	mux.HandleFunc("/student_tickets", student_ticket_controller.GetStudentTicketsHandler)
	mux.HandleFunc("/student_tickets/add", student_ticket_controller.AddStudentTicketHandler)
	mux.HandleFunc("/student_tickets/bySerialNumber", student_ticket_controller.GetStudentTicketBySerialNumberHandler)

	mux.HandleFunc("/students", student_controller.GetStudentsHandler)

	mux.HandleFunc("/ask_admin/register", ask_admin_controller.AskAdminRegisterHandler)
	mux.HandleFunc("/ask_admin/contract/sign", ask_admin_controller.AskAdminSignContractHandler)
	mux.HandleFunc("/ask_admin/settle", ask_admin_controller.AskAdminSettleHandler)
	mux.HandleFunc("/ask_admin/unsettle", ask_admin_controller.AskAdminUnsettleHandler)
	mux.HandleFunc("/ask_admin/resettle", ask_admin_controller.AskAdminResettleHandler)
	mux.HandleFunc("/ask_admin/actions", ask_admin_controller.GetActionsHandler)
	// // only for admin
	mux.HandleFunc("/ask_admin/actions/resolve", ask_admin_controller.ResolveActionHandler)

	mux.HandleFunc("/dormitory/load/stats", dormitory_load_controller.GetDormitoryLoadHandler)

	mux.HandleFunc("/user/register", user_controller.RegisterUserHandler)
	mux.HandleFunc("/user/login", user_controller.LoginUserHandler)

	port := "80"

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	fmt.Println("Server is listening on port " + port + "...")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db := db.InitDB()
	
	RunHttpServer(db)
}
