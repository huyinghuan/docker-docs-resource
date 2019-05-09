package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

func getIPs() (list []string) {

	ifaces, _ := net.Interfaces()
	// handle err
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			localIP := ip.String()
			if strings.Contains(localIP, "::") {
				continue
			}
			if localIP == "127.0.0.1" {
				continue
			}
			list = append(list, localIP)
			// process IP address
		}
	}
	return
}
func main() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, fmt.Sprintf("Hello Docker, I'm %s", strings.Join(getIPs(), ",")))
	})
	http.ListenAndServe(":"+port, nil)
}
