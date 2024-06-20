package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/tolling/types"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "The listen address of the htrtp server")
	store := NewMemoryStore()
	svc := NewInvoiceAggregator(store)
	makeHTTPTransport(*listenAddr, svc)
}

func makeHTTPTransport(listenAddr string, svc Aggregator) {
	fmt.Println("HTTP transport running on port", listenAddr)
	http.HandleFunc("/aggregate", handleAggregate(svc))
	http.ListenAndServe(listenAddr, nil)
}

func handleAggregate(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			panic(err)
		}
	}
}
