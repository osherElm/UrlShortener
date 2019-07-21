package main

import (
	"UrlShortener/config"
	"UrlShortener/handler"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {

	var configPath = flag.String("config", "./config/config.json", "path of the config file")
	flag.Parse()

	cfg, err := config.ReadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		Handler: handler.Initialize(cfg.Options.Prefix),
	}
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("%v", err)
		} else {
			log.Println("Server closed!")
		}
	}()
}
