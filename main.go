package main // Package where the code belongs to

import (
	// List of packages to be imported
	"net/http"
)

func main()  {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello World"))
	})

	// Nil to GO that we want to use the default server multiplexer
	// or MUX to handle the request
	http.ListenAndServe(":8000", nil)
}