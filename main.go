package main // Package where the code belongs to

import (
	// List of packages to be imported
	"net/http"
	"io/ioutil"
)

func main()  {
	// the second parameter asterisk means that the parameter will actually
	// be a pointer to and http.Request object
	http.Handle("/", new(MyHandler))

	// Nil to tell GO that we want to use the default server multiplexer
	// or MUX to handle the request
	http.ListenAndServe(":8000", nil)
}


type MyHandler struct {
	http.Handler
}

func (this *MyHandler) ServeHTTP( w http.ResponseWriter, req *http.Request){
	path := "public/"+req.URL.Path
	data, err := ioutil.ReadFile(string(path))

	if(err == nil){
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}

}