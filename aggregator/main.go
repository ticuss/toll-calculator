package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/tolling/client"
	"github.com/tolling/types"
	"google.golang.org/grpc"
)

func main() {
	httpListenAddr := flag.String("httpAddr", ":3000", "The listen address of the htrtp server")
	grpcListenAddr := flag.String("grpcAddr", ":3001", "The listen address of the htrtp server")
	store := NewMemoryStore()
	svc := NewInvoiceAggregator(store)
	svc = NewLogMiddleware(svc)
	time.Sleep(2 * time.Second)
	c, err := client.NewGRPCClient(*grpcListenAddr)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		log.Fatal(makeGRPCTransport(*grpcListenAddr, svc))
	}()
	c.Aggregate(context.Background(), &types.AggregateRequest{ObuID: 1, Value: 100, Unix: time.Now().Unix()})
	log.Fatal(makeHTTPTransport(*httpListenAddr, svc))
	fmt.Println("Hello")
}

func makeHTTPTransport(listenAddr string, svc Aggregator) error {
	fmt.Println("HTTP transport running on port", listenAddr)
	http.HandleFunc("/aggregate", handleAggregate(svc))
	http.HandleFunc("/invoice", handleGetInvoice(svc))
	return http.ListenAndServe(listenAddr, nil)
}

func makeGRPCTransport(listenAddr string, svc Aggregator) error {
	fmt.Println("GRPC transport running on port", listenAddr)
	// Make a tcp listener
	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	// Make a new grpx native server with (options)
	server := grpc.NewServer([]grpc.ServerOption{}...)
	// Register (our) GRPC server
	types.RegisterAggregatorServer(server, NewAggregatorGRPCServer(svc))
	return server.Serve(ln)
}

func handleGetInvoice(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		obuIDStr := r.URL.Query().Get("obu")
		if obuIDStr == "" {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "obuID is required"})
			return
		}
		obuID, err := strconv.Atoi(obuIDStr)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "obuID must be an integer"})
			return
		}
		inv, err := svc.CalculateInvoice(obuID)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, http.StatusOK, inv)
	}
}

func handleAggregate(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		err := svc.AggregateDistance(distance)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		fmt.Printf("error %s\n", err)
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
