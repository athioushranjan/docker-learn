package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

func main() {
	// add print handler to the list
	http.HandleFunc("/docker", printHandler)
	// listen and serve to port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logrus.Fatalf("unable to listen the server at the port %s", "8080")
	}
	logrus.Println("server connected on port 8080")

}

func printHandler(w http.ResponseWriter, r *http.Request) {
	// print message
	_, err := fmt.Fprintf(w, "learn docker")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("can't print data %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
