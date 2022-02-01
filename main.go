package main

import (
	"fmt"
	"github.com/athioushranjan/docker-learn/database"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	// add print handler to the list
	http.HandleFunc("/docker", printHandler)
	// load env
	//database.LoadEnv()
	logrus.Println(os.Getenv("DB_HOST"))
	logrus.Println(os.Getenv("DB_PORT"))
	logrus.Println(os.Getenv("DB_NAME"))
	logrus.Println(os.Getenv("DB_USER"))
	logrus.Println(os.Getenv("DB_PASS"))
	// connect database and migrate
	if err := database.ConnectAndMigrate(os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		database.SSLModeDisable); err != nil {
		logrus.Panicf("Failed to initialize and migrate database with error: %+v", err)
	}
	logrus.Print("database connection and migration successful!!")
	// listen and serve to port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logrus.Fatalf("unable to listen the server at the port %s", "8080")
	}
	logrus.Println("server connected on port 8080")

}

func printHandler(w http.ResponseWriter, r *http.Request) {
	// print message
	_, err := fmt.Fprintf(w, "first docker file")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Fatalf("can't print data %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
