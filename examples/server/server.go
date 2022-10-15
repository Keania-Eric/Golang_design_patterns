package main

// An example using the decorator pattern to write an Http server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type MyServer struct{}

func (m *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Decorator!")
}

type LoggerServer struct {
	Handler   http.Handler
	LogWriter io.Writer
}

func (s *LoggerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(s.LogWriter, "Request URI: %s \n", r.RequestURI)
	fmt.Fprintf(s.LogWriter, "HOST: %s \n", r.Host)
	fmt.Fprintf(s.LogWriter, "Content Length: %d \n", r.ContentLength)
	fmt.Fprintf(s.LogWriter, "Method: %s \n", r.Method)
	fmt.Fprintf(s.LogWriter, "------------------------- \n")

	s.Handler.ServeHTTP(w, r)
}

type BasicAuthMiddleware struct {
	Handler  http.Handler
	User     string
	Password string
}

func (s *BasicAuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()

	if ok {
		if user == s.User && pass == s.Password {
			s.Handler.ServeHTTP(w, r)
		} else {
			fmt.Fprintf(w, "User or password incorrect\n")
		}
	} else {
		fmt.Fprintln(w, "Error trying to retrieve data from Basic auth")
	}
}

func main() {

	fmt.Println("Pick a choice of server from the following:")
	fmt.Println("1. - Plain server")
	fmt.Println("2. - Server with logging")
	fmt.Println("3. - Server with logging and authentication")

	var selection int
	var server http.Handler
	fmt.Fscanf(os.Stdin, "%d", &selection)

	switch selection {
	case 1:
		server = new(MyServer)
	case 2:
		server = &LoggerServer{LogWriter: os.Stdout, Handler: new(MyServer)}
	case 3:
		var user, password string

		fmt.Println("Enter user and password seperated by space")
		fmt.Fscanf(os.Stdin, "%s %s", &user, &password)

		server = &LoggerServer{
			Handler:   &BasicAuthMiddleware{Handler: new(MyServer), User: user, Password: password},
			LogWriter: os.Stdout,
		}
	default:
		server = new(MyServer)
	}

	http.Handle("/", server)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
