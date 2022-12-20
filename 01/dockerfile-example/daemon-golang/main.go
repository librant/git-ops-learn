package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", mainIndex)
	mux.HandleFunc("/healthz", healthz)

	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func mainIndex(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}
	w.Header().Add("systemEnv", os.Getenv("GOPATH"))
	cliIP := clientIP(r)
	w.Write([]byte(fmt.Sprintf("cliIP: %s", cliIP)))
	log.Printf("Success! Response code: %d client ip: %s", http.StatusOK, cliIP)
}

func clientIP(r *http.Request) string {
	ip := strings.TrimSpace(strings.Split(r.Header.Get("x-Forwarded-For"), ",")[0])
	if len(ip) != 0 {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if len(ip) != 0 {
		return ip
	}
	ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	if err != nil {
		return ""
	}
	return ip
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fmt.Sprintf("%d", http.StatusOK))
}




