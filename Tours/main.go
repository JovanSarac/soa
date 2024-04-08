package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"tours_service/handler"
	"tours_service/repository"
	"tours_service/service"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	gorillaHandlers "github.com/gorilla/handlers"
)

func initMongoDb() *mongo.Client {

	dburi := "mongodb://localhost:27017"

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		fmt.Print(err)
	}

	return client
}
func manageRouter(client *mongo.Client) http.Server {
	FacilityRepository := &repository.FacilityRepository{FacilityClient: client}
	FacilityService := &service.FacilityService{FacilityRepository: FacilityRepository}
	FacilityHandler := &handler.FacilityHandler{FacilityService: FacilityService}

	router := mux.NewRouter().StrictSlash(true)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/facilities", FacilityHandler.CreateFacility)
	postRouter.Use(FacilityHandler.MiddlewareFacilityDeserialization)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/facilities/{id}", FacilityHandler.DeleteFacility)

	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))

	server := http.Server{
		Addr:         ":8080",
		Handler:      cors(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	return server

}

// func initDB() *gorm.DB {
// 	connStr := "user=postgres dbname=explorer password=super sslmode=disable port=5432 host=database"
// 	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	return db
// }

func startServer(FacilityHandler *handler.FacilityHandler, KeypointHandler *handler.KeypointHandler, tourHandler *handler.TourHandler, tourRatingHandler *handler.TourRatingHandler, tourProblemHandler *handler.TourProblemHandler) {
	router := mux.NewRouter().StrictSlash(true)

	// router.HandleFunc("/facilities", FacilityHandler.Create).Methods("POST")
	// router.HandleFunc("/facilities/{id}", FacilityHandler.Delete).Methods("DELETE")

	router.HandleFunc("/keypoints", KeypointHandler.Create).Methods("POST")

	router.HandleFunc("/createTour", tourHandler.CreateTour).Methods("POST")
	router.HandleFunc("/tours", tourHandler.Update).Methods("PUT")
	router.HandleFunc("/tours/{id}", tourHandler.GetTourByID).Methods("GET")

	router.HandleFunc("/createTourRating", tourRatingHandler.CreateTourRating).Methods("POST")

	router.HandleFunc("/getByAuthorId/{authorId}", tourProblemHandler.GetByAuthorId).Methods("GET")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Add("Content-Type", "application/json")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
	println("Server starting")

	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(router)))
}

// func setUpDependencies(database *gorm.DB) (*handler.FacilityHandler, *handler.TourHandler, *handler.KeypointHandler, *handler.TourRatingHandler, *handler.TourProblemHandler) {
// 	FacilityRepository := &repository.FacilityRepository{DatabaseConnection: database}
// 	FacilityService := &service.FacilityService{FacilityRepository: FacilityRepository}
// 	FacilityHandler := &handler.FacilityHandler{FacilityService: FacilityService}

// 	KeypointRepository := &repository.KeypointRepository{DatabaseConnection: database}
// 	KeypointService := &service.KeypointService{KeypointRepository: KeypointRepository}
// 	KeypointHandler := &handler.KeypointHandler{KeypointService: KeypointService}

// 	tourRepository := &repository.TourRepository{DatabaseConnection: database}
// 	tourService := &service.TourService{TourRepository: tourRepository}
// 	tourHandler := &handler.TourHandler{TourService: tourService}

// 	tourRatingRepository := &repository.TourRatingRepository{DatabaseConnection: database}
// 	tourRatingService := &service.TourRatingService{TourRatingRepository: tourRatingRepository}
// 	tourRatingHandler := &handler.TourRatingHandler{TourRatingService: tourRatingService}

// 	tourProblemRepository := &repository.TourProblemRepository{DatabaseConnection: database}
// 	tourProblemService := &service.TourProblemService{TourProblemRepository: tourProblemRepository,
// 		TourService: tourService}
// 	tourProblemHandler := &handler.TourProblemHandler{TourProblemService: tourProblemService}

// 	return FacilityHandler, tourHandler, KeypointHandler, tourRatingHandler, tourProblemHandler
// }

func main() {
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client := initMongoDb()

	err := client.Connect(timeoutContext)
	if err != nil {
		fmt.Print(err)
	}
	logger := log.New(os.Stdout, "[logger-main] ", log.LstdFlags)
	server := manageRouter(client)

	logger.Println("Server listening on port", "8080")
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	logger.Println("Received terminate, graceful shutdown", sig)

	if server.Shutdown(timeoutContext) != nil {
		logger.Fatal("Cannot gracefully shutdown...")
	}
	logger.Println("Server stopped")

	print("ok")
}
