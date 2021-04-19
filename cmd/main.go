package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	logrus.Info("received [GET] /hello request")
	if _, err := fmt.Fprintf(w, "hello\n"); err != nil {
		panic(err)
	}
}

func headers(w http.ResponseWriter, req *http.Request) {
	logrus.Info("received [GET] /headers request")
	for name, headers := range req.Header {
		for _, h := range headers {
			if _, err := fmt.Fprintf(w, "%v: %v\n", name, h); err != nil {
				panic(err)
			}
		}
	}
}

func main() {
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "80"
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", serverPort), nil); err != nil {
		panic(err)
	}
}
