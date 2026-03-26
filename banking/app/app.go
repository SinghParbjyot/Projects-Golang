package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/singhparbjyot/banking/domain"
	"github.com/singhparbjyot/banking/service"
)

func SanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment varibale not defined....")
	}
}
func Start() {

	dbClient := getDbClient()
	customerRepositroyDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

	//- Creates a instance  multiplexor of requests
	router := mux.NewRouter()
	ch := &CustomerHandlers{service.NewCustomerService(customerRepositroyDb)}
	ah := &AccountHandler{service: service.NewAccountService(accountRepositoryDb)}
	//Defining routes for customers
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers?status", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id}/account/{account_id}", ah.MakeTransaction).Methods(http.MethodPost)

	//Route that return one customer requested by get with an id
	//wiring
	err := godotenv.Load()
	if err != nil {

		log.Println("Advertencia: no se pudo cargar el archivo .env:", err)
	}
	// This function to start the server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	addr := fmt.Sprintf("%s:%s", address, port)

	log.Fatal(http.ListenAndServe(addr, router))

}

func getDbClient() *sqlx.DB {
	//Load the .env variables
	err := godotenv.Load()
	//We define here the .env variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	//dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	// (opcional, para depurar)
	//log.Println("DB_USER:", dbUser)
	//log.Println("DB_HOST:", dbHost)
	//log.Println("DB_PORT:", dbPort)
	//log.Println("DB_NAME:", dbName)
	//Conecting with db postgresql
	//URL
	// DB_ADDR lo puedes seguir leyendo si quieres para logs:
	//dbAddr := os.Getenv("DB_ADDR") // "localhost:5432" solo para imprimir

	// Conect with Postgres (lib/pq)
	dataSource := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)
	// Function open sql(DriverName,URL)
	client, err := sqlx.Open("postgres", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	//Settings for the database
	client.SetConnMaxIdleTime(time.Minute)
	//Maxium of open connections
	client.SetMaxOpenConns(10)
	//Maxium life connections
	client.SetConnMaxLifetime(10)
	return client
}
