package main

import (
	"fmt"
	"flag"
	"os"

	"github.com/javierlopez987/seminarioGoLang/internal/config"
	"github.com/javierlopez987/seminarioGoLang/internal/database"
	"github.com/javierlopez987/seminarioGoLang/internal/service/flight"

	"github.com/gin-gonic/gin"

)

func main()  {
	
	// Lectura de configuracion
	cfg := readConfig()

	// Instanciacion de db
	db, err := database.NewDatabase(cfg)
	defer db.Close()
	
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Instanciacion un servicio y le inyecta la configuracion y la base de datos
	service, _ := flight.New(db, cfg)
	httpService := flight.NewHTTPTransport(service)

	r:= gin.Default()
	httpService.Register(r)
	r.Run()
}

func readConfig() *config.Config {
	configFile := flag.String("config", "./config.yaml", "este es el servicio de configuracion")
	flag.Parse()
	
	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return cfg
}