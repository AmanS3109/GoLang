package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

func main() {
	flag.Parse()
	http.HandleFunc("/", handler)
	http.HandleFunc("/event", event)
	http.ListenAndServe(":8080", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

//simple event stream

func event(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	token := []string{"Hello", "Shruti", "How", "Are", "You"}

	for _, t := range token {
		content := fmt.Sprintf("data: %s\n\n", string(t))
		w.Write([]byte(content))
		w.(http.Flusher).Flush()

		//intentional delay
		time.Sleep(time.Millisecond * 420)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
