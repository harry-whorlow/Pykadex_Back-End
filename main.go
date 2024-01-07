package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	m := http.NewServeMux()

	const addr = ":8000"

	m.HandleFunc("/", handleSearch)

	srv := http.Server{
		Handler: m,
		Addr: addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout: 30 * time.Second,
	}

	fmt.Println("server started on port", addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	const page = `
		<html>
			<body>
				<p>hello from go server</p>
			</body>
		</html>
	`

	w.WriteHeader(200)
	w.Write([]byte(page))
}
