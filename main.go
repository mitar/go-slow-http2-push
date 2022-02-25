package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/small-sub", func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("sub request %s: start\n", req.URL.String())
		now := time.Now()
		w.Header().Set("Content-Type", "application/octet-stream")
		f, err := os.Open("small")
		if err != nil {
			log.Fatalf("%v", err)
		}
		defer f.Close()
		http.ServeContent(w, req, "", time.Time{}, f)
		fmt.Printf("sub request %s: %.3f\n", req.URL.String(), time.Since(now).Seconds()*1000)
	})
	http.HandleFunc("/large-sub", func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("sub request %s: start\n", req.URL.String())
		now := time.Now()
		w.Header().Set("Content-Type", "application/octet-stream")
		f, err := os.Open("large")
		if err != nil {
			log.Fatalf("%v", err)
		}
		defer f.Close()
		http.ServeContent(w, req, "", time.Time{}, f)
		fmt.Printf("sub request %s: %.3f\n", req.URL.String(), time.Since(now).Seconds()*1000)
	})
	http.HandleFunc("/small", func(w http.ResponseWriter, req *http.Request) {
		now := time.Now()
		fmt.Printf("main %s: start\n", req.URL.String())
		pusher, ok := w.(http.Pusher)
		if ok {
			for i := 0; i < 100; i++ {
				if err := pusher.Push(fmt.Sprintf("/small-sub?%d", i), nil); err != nil {
					log.Printf("Failed to push: %v", err)
				}
			}
		}
		w.Header().Set("Content-Type", "text/plain")
		http.ServeContent(w, req, "", time.Time{}, strings.NewReader("main"))
		fmt.Printf("main %s: %.3f\n", req.URL.String(), time.Since(now).Seconds()*1000)
	})
	http.HandleFunc("/large", func(w http.ResponseWriter, req *http.Request) {
		now := time.Now()
		fmt.Printf("main %s: start\n", req.URL.String())
		pusher, ok := w.(http.Pusher)
		if ok {
			for i := 0; i < 100; i++ {
				if err := pusher.Push(fmt.Sprintf("/large-sub?%d", i), nil); err != nil {
					log.Printf("Failed to push: %v", err)
				}
			}
		}
		w.Header().Set("Content-Type", "text/plain")
		http.ServeContent(w, req, "", time.Time{}, strings.NewReader("main"))
		fmt.Printf("main %s: %.3f\n", req.URL.String(), time.Since(now).Seconds()*1000)
	})
	log.Fatal(http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil))
}
