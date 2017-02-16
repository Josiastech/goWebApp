package main // Package where the code belongs to

import (
	// List of packages to be imported
	"net/http"
	"strings"
	"os"
	"bufio"
)

func main()  {
	// the second parameter asterisk means that the parameter will actually
	// be a pointer to and http.Request object
	http.Handle("/", new(MyHandler))

	// Nil to tell GO that we want to use the default server multiplexer
	// or MUX to handle the request
	//http.ListenAndServe(":8000", nil)
}


type MyHandler struct {
	http.Handler
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path
	f, err := os.Open(path)

	if err == nil {
		bufferedReader := bufio.NewReader(f)
		var contentType string
		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else {
			contentType = "text/plain"
		}



		w.Header().Add("Content Type", contentType)
		bufferedReader.WriteTo(w)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}
}