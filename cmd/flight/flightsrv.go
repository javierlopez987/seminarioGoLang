package main

import (
	"fmt"
	"flag"
	"os"
	"time"

	"github.com/javierlopez987/seminarioGoLang/internal/config"
	"github.com/javierlopez987/seminarioGoLang/internal/database"
	"github.com/javierlopez987/seminarioGoLang/internal/service/flight"

	"github.com/jmoiron/sqlx"
)

func main()  {
	
	// Lectura de configuracion
	cfg := readConfig()

	// Instanciacion de db
	db, err := database.NewDatabase(cfg)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err := createSchema(db); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Instanciacion un servicio y le inyecta la configuracion y la base de datos
	service, _ := flight.New(db, cfg)

	for _, m := range service.FindAll() {
		fmt.Println(m)
	}
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

// Solo en entornos de desarrollo
func createSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS flights (
		id integer primary key autoincrement,
		airlinename text,
		flightnumber text,
		departuredatetime text,
		arrivaldatetime text,
		departureairport text,
		arrivalairport text
		);`

	// ejecuta una query en el servidor
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	// MustExec tira un panic on error
	insertMessage := `INSERT INTO flights (
		AirlineName, FlightNumber, DepartureDateTime, 
		ArrivalDateTime, DepartureAirport, ArrivalAirport) VALUES (?, ?, ?, ?, ?, ?)`
	name := fmt.Sprintf("Airline %v", time.Now().Nanosecond())
	number := fmt.Sprintf("Flight %v", time.Now().Nanosecond())
	departure := fmt.Sprintf("Departure Date %v", time.Now().Nanosecond())
	arrival := fmt.Sprintf("Arrival Date %v", time.Now().Nanosecond())
	db.MustExec(insertMessage, name, number, departure, arrival, "EZE", "MAD")
	return nil
}