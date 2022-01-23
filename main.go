package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/docker", printHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logrus.Fatalf("unable to listen the server at the port %s", "8080")
	}
	logrus.Println("server connected on port 8080")

}

func printHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "learn docker")
	if err != nil {
		log.Fatalf("can't print data %v", err)
	}
	w.WriteHeader(http.StatusOK)
}
