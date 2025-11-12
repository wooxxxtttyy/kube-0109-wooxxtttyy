package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

var gwAppName, gwBindPort, gwPriority, gwThreads string

func init() {
	var ok bool
	if gwAppName, ok = os.LookupEnv("GATEWAY_APP_NAME"); !ok {
		gwAppName = "gateway"
	}

	if gwBindPort, ok = os.LookupEnv("GATEWAY_BIND_PORT"); !ok {
		gwBindPort = "8080"
	}

	if gwPriority, ok = os.LookupEnv("GATEWAY_PRIORITY"); !ok {
		gwPriority = "1"
	}

	if gwThreads, ok = os.LookupEnv("GATEWAY_THREADS"); !ok {
		gwThreads = "4"
	}
}

func main() {
	fmt.Println("GATEWAY_APP_NAME:", gwAppName)
	fmt.Println("GATEWAY_BIND_PORT:", gwBindPort)
	fmt.Println("GATEWAY_PRIORITY:", gwPriority)
	fmt.Println("GATEWAY_THREADS:", gwThreads)

	pgCtx, pgCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer pgCancel()

	pgConn, err := pgx.Connect(pgCtx, "postgres://user:pass@localhost:5432/demodb")
	if err != nil {
		fmt.Println("database connect error:", err)
	}

	_ = pgConn

	http.HandleFunc("/", rootHandler)

	fmt.Println("\nserver started")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h2>Hello from GoLang server</h2>"))
}
