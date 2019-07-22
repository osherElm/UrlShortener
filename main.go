package main

import (
	"UrlShortener/config"
	"UrlShortener/handler"
	"flag"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	var configPath = flag.String("config", "./config/config.json", "path of the config file")
	flag.Parse()

	cfg, err := config.ReadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := fasthttp.ListenAndServe(":8080", handler.Initialize(cfg.Options.Prefix)); err != http.ErrServerClosed {
			log.Fatalf("%v", err)
		} else {
			log.Println("Server closed!")
		}
	}()

	// Check for a closing signal
	// Graceful shutdown
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	s := <-gracefulStop
	log.Printf("caught sig: %+v", s)
	log.Printf("Gracefully shutting down server...")
}
