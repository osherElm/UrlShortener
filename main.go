package main

import (
	"UrlShortener/config"
	"UrlShortener/handler"
	"UrlShortener/storage/mysql"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/valyala/fasthttp"
)

func main() {

	var configPath = flag.String("config", "./config/config.json", "path of the config file")
	flag.Parse()

	cfg, err := config.ReadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	db, err := mysql.Initialize(cfg.MySQL.User, cfg.MySQL.Password, cfg.MySQL.DB)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	go func() {
		if err := fasthttp.ListenAndServe(":8080", handler.Initialize(cfg.Options.Prefix, &db)); err != http.ErrServerClosed {
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
