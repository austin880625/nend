package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/austin880625/nend/cmd_args"
)

var args cmd_args.Args

func handler(port int) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		client := &http.Client{}

		r.URL.Scheme = "http"
		r.URL.Host = fmt.Sprintf("localhost:%d", port)
		fmt.Println(r.URL.String())
		proxyReq, err := http.NewRequest(r.Method, r.URL.String(), r.Body)
		if err != nil {
			http.Error(w, "Failed to create request", http.StatusInternalServerError)
			return
		}
		proxyReq.Header = r.Header.Clone()
		resp, err := client.Do(proxyReq)
		if err != nil {
			http.Error(w, "Failed to forward request", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		for key, value := range resp.Header {
			w.Header()[key] = value
		}

		w.WriteHeader(resp.StatusCode)

		_, err = io.Copy(w, resp.Body)
		if err != nil {
			log.Println("Failed to copy response body:", err)
		}
	}
}

func Run(cmdArgs cmd_args.Args) {
	args = cmdArgs
	fmt.Printf("backend port: %d, frontend port: %d\n", args.BPort, args.FPort)
	apiHandler := handler(args.BPort)
	defaultHandler := handler(args.FPort)
	http.HandleFunc("/api/v1/", apiHandler)

	// Handle all other requests by setting a default handler
	http.HandleFunc("/", defaultHandler)

	// Start the server on port 8080
	fmt.Printf("Server is running on port %d...\n", args.Port)
	err := http.ListenAndServe(":"+strconv.Itoa(args.Port), nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
