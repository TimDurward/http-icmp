package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

// Host represents the request headers that HTTPIcmp func uses
type Host struct {
	// Hostname: Host that will be resolved to IPv4/IPv6 for http request
	Hostname string `json:"hostname"`

	// Count: Count is the amount of times ICMP is invoked
	Count int `json:"count"`
}

// HTTPIcmp decodes Request headers that are described by Host struct
// It also encodes a Response from ICMP func
// icmp accepts Ping data types
func HTTPIcmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var host Host
	json.NewDecoder(r.Body).Decode(&host)

	ping := icmp(host.Hostname, host.Count)
	json.NewEncoder(w).Encode(ping)

}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/ping", HTTPIcmp).Methods("POST")

	var term = make(chan os.Signal)
	signal.Notify(term, syscall.SIGTERM)
	signal.Notify(term, syscall.SIGINT)

	//Go routine to shutdown server gracefully
	go func() {
		sig := <-term
		fmt.Printf("\n[ %+v ] Shutting down...\n", sig)
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}()

	log.Fatal(http.ListenAndServe(":8000", router))
}
