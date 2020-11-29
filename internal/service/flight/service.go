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
	return nil
}

// FindByID ...
func (s service) FindByID(ID int) (f *Flight)  {
	return nil
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
func (s service) Delete(ID int)  {
	
}

// Update ...
func (s service) Update(f Flight) {
	
}