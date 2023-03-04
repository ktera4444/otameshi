package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world from backend as webserver with hot-reloading")
}

func main()  {
	http.HandleFunc("/", helloWorld)

	fmt.Println("Starting webserver on port:3000...")
	// serve from port:3000 (inside container) -> need to expose this port outside
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}