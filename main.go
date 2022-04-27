package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"mathapi/base"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
)

func main() {
	var (
		basePath = flag.String("service.base.path", "mathapi", "Name of microservice")
		version  = flag.String("service.version", "v1", "Version of microservice (Default v1)")
		httpPort = flag.Int("http.port", 8080, "This is the port at which http requests are accepted (Default :8080)")
	)

	flag.Parse()

	errs := make(chan error)
	var s base.Service
	{
		s = base.NewService()
	}
	h := base.MakeHTTPHandler(s, *version, *basePath)

	httpServer := http.Server{
		Addr:    ":" + strconv.Itoa(*httpPort),
		Handler: handlers.RecoveryHandler()(h),
	}
	go func() {
		log.Printf("Application HTTP server is now starting on port %d", *httpPort)
		errs <- httpServer.ListenAndServe()
	}()
	error := <-errs
	errHTTPServer := httpServer.Shutdown(context.Background())
	log.Fatal(fmt.Sprintf("Error occured,shutting down %s, HTTP server error %s", error.Error(), errHTTPServer.Error()))

}
