package flight

import (
	"github.com/javierlopez987/seminarioGoLang/internal/config"

	"github.com/jmoiron/sqlx"
)

// Flight ...
type Flight struct {
	ID int
	AirlineName string
	FlightNumber string
	DepartureDateTime string
	ArrivalDateTime string
	DepartureAirport string
	ArrivalAirport string
}

// NewFlight ...
func NewFlight(ID int, AirlineName string,
	FlightNumber string,
	DepartureDateTime string,
	ArrivalDateTime string,
	DepartureAirport string,
	ArrivalAirport string) (Flight, error) {
	return Flight{ID, AirlineName, FlightNumber, DepartureDateTime, ArrivalDateTime, DepartureAirport, ArrivalAirport}, nil
}

// Service ...
type Service interface {
	Add(Flight) error
	FindByID(int) *Flight
	FindAll() []*Flight
	Delete(int)
	Update(Flight)
}

type service struct {
	db *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

// Add ...
func (s service) Add(f Flight) error {
	query := `INSERT INTO flights (
		AirlineName, FlightNumber, DepartureDateTime, 
		ArrivalDateTime, DepartureAirport, ArrivalAirport) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := s.db.Exec(query, f.AirlineName, 
		f.FlightNumber, f.DepartureDateTime, f.ArrivalDateTime, f.DepartureAirport, f.ArrivalAirport)
	if err != nil {
		panic(err)
	}

	return err
}

// FindByID ...
func (s service) FindByID(ID int) (*Flight)  {
	var f []*Flight
	if err := s.db.Select(&f, "SELECT * FROM flights WHERE id = ?", ID); err != nil {
		panic(err)
	}

	if len(f) > 0 {
		return f[0]
	} else {
		return nil
	}
}

// FindAll ...
func (s service) FindAll() []*Flight {
	var list []*Flight
	if err := s.db.Select(&list, "SELECT * FROM flights"); err != nil {
		panic(err)
	}
	return list
}

// Delete ...
func (s service) Delete(ID int) {
	query := "DELETE FROM flights WHERE id = ?"
	s.db.Exec(query, ID)
}

// Update ...
func (s service) Update(f Flight) {
	query := `UPDATE flights 
		SET AirlineName = ?, FlightNumber = ?, DepartureDateTime = ?, 
		ArrivalDateTime = ?, DepartureAirport = ?, ArrivalAirport = ? 
		WHERE ID = ?`
	s.db.Exec(query, f.AirlineName, 
		f.FlightNumber, f.DepartureDateTime, f.ArrivalDateTime, f.DepartureAirport, f.ArrivalAirport, f.ID)
}