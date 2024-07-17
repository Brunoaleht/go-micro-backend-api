package main

import (
	"context"
	"database/sql"
	httpHandler "go-backend-api/internal/events/infra/http"
	"go-backend-api/internal/events/infra/repository"
	"go-backend-api/internal/events/infra/service"
	"go-backend-api/internal/events/usecase"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Events API
// @version 1.0
// @description This is a server for managing events. Imersão Full Cycle
// @host localhost:8080
// @BasePath /
func main() {

	// Openning a connection to the database
	db, err := sql.Open("mysql", "test_user:test_password@tcp(golang-mysql:3306)/test_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Starting Repository
	eventRepo, err := repository.NewMysqlEventRepository(db)
	if err != nil {
		log.Fatal(err)
	}

	//Urls base path for the Partners API
	partnersAPIBasePath := map[int]string{
		1: "http://host.docker.internal:8000/partner1",
		2: "http://host.docker.internal:8000/partner2",
	}

	// Starting the use case
	listEventsUseCase := usecase.NewListEventsUseCase(eventRepo)
	getEventUseCase := usecase.NewGetEventUseCase(eventRepo)
	createEventUseCase := usecase.NewCreateEventUseCase(eventRepo)
	partnerFactory := service.NewPartnerFactory(partnersAPIBasePath)
	buyTicketsUseCase := usecase.NewBuyTicketsUseCase(eventRepo, partnerFactory)
	createSpotsUseCase := usecase.NewCreateSpotsUseCase(eventRepo)
	listSpotsUseCase := usecase.NewListSpotsUseCase(eventRepo)

 
	// Starting the handler HTTP
	eventsHandler := httpHandler.NewEventsHandler(
		listEventsUseCase,
		getEventUseCase,
		createEventUseCase,
		buyTicketsUseCase,
		createSpotsUseCase,
		listSpotsUseCase,
	)
	router := http.NewServeMux()
	router.HandleFunc("/events", eventsHandler.ListEvents)
	router.HandleFunc("/events/{eventId}", eventsHandler.GetEvent)
	router.HandleFunc("/events/{eventId}/spots", eventsHandler.ListSpots)
	router.HandleFunc("POST /events", eventsHandler.CreateEvent)
	router.HandleFunc("POST /events/buy-tickets", eventsHandler.BuyTickets)
	router.HandleFunc("POST /events/{eventId}/spots", eventsHandler.CreateSpots)

	// Starting the server
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Canal para escutar sinais do sistema operacional
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		// Recebido sinal de interrupção, iniciando o graceful shutdown
		log.Println("Recebido sinal de interrupção, iniciando o graceful shutdown...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			// Erro ao desligar o servidor
			log.Printf("Erro ao desligar o servidor: %v\n", err)
		}

		close(idleConnsClosed)
	}()

	//init server HTTP
	log.Println("Server started at :8080")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Erro ao iniciar o servidor: %v\n", err)
	}

	<-idleConnsClosed
	log.Println("Servidor desligado com sucesso")

}