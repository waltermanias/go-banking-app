package app

import (
	"banking-app/domain"
	"banking-app/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable not defined")
	}
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dataSource := fmt.Sprintf("%s@tcp(localhost:%s)/%s", dbUser, dbPort, dbName)

	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}

func Start() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sanityCheck()

	router := mux.NewRouter()

	dbClient := getDbClient()

	// Create Repositories
	customerRepository := domain.NewCustomerRepositoryDb(dbClient)
	accountRepository := domain.NewAccountRepositoryDb(dbClient)

	ch := CustomerHandlers{service.NewCustomerService(customerRepository)}
	ah := AccountHandlers{service.NewAccountService(accountRepository)}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getById).Methods(http.MethodGet).Subrouter()
	router.HandleFunc("/customers/{customer_id:[0-9]+}/accounts", ah.save).Methods(http.MethodPost)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}
