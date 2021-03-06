package api

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

// Serve ... start go web search server
func Serve(port int) error {
	mux := createHTTPMux()

	l, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		return err
	}

	log.Println("Server is running on http://localhost:8080")
	return http.Serve(l, mux)
}

func createHTTPMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
	mux.HandleFunc("/search", resultsPageHandler())
	mux.HandleFunc("/", homePageHandler())

	return mux
}
