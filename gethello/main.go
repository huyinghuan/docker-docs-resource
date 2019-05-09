package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("SERVER_PORT")
	target := os.Getenv("SERVER_TARGET")
	if port == "" {
		port = "8080"
	}
	if target == "" {
		log.Panicln("No target server")
	}
	client := http.Client{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		req, _ := http.NewRequest("GET", target, nil)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Fprintf(w, fmt.Sprintf("GET Errror %v", err))
			return
		}
		defer resp.Body.Close()
		b, _ := ioutil.ReadAll(resp.Body)
		fmt.Fprintf(w, string(b))
	})
	http.ListenAndServe(":"+port, nil)
}
