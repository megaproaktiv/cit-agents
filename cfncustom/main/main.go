package main

import (
	"log"
	"net/http"
	"time"

	"github.com/megaproaktiv/citagents/cfncustom"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)


func main() {

	apiServer := &http.Server{
		Addr:         ":443",
		Handler:      cfncustom.ApiRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	controllerServer := &http.Server{
		Addr:         ":8081",
		Handler:      cfncustom.ControllerRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return apiServer.ListenAndServe()
	})

	g.Go(func() error {
		return controllerServer.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
