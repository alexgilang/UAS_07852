package model

import (
	"time"
)

// City represents a city node in linked list
type City struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Next *City  `json:"-"`
}

// Station represents a station node in linked list
type Station struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Address string   `json:"address"`
	Phone   string   `json:"phone"`
	CityID  int      `json:"city_id"`
	Next    *Station `json:"-"`
}

// Route represents a route node in linked list
type Route struct {
	ID         int     `json:"id"`
	FromCityID int     `json:"from_city_id"`
	ToCityID   int     `json:"to_city_id"`
	Distance   float64 `json:"distance"`
	Next       *Route  `json:"-"`
}

// Schedule represents a schedule node in linked list
type Schedule struct {
	ID            int       `json:"id"`
	RouteID       int       `json:"route_id"`
	TrainName     string    `json:"train_name"`
	DepartureTime time.Time `json:"departure_time"`
	ArrivalTime   time.Time `json:"arrival_time"`
	Price         float64   `json:"price"`
	Next          *Schedule `json:"-"`
}

// Booking represents a booking node in linked list
type Booking struct {
	ID          int       `json:"id"`
	ScheduleID  int       `json:"schedule_id"`
	UserName    string    `json:"user_name"`
	UserEmail   string    `json:"user_email"`
	Seats       int       `json:"seats"`
	TotalPrice  float64   `json:"total_price"`
	BookingDate time.Time `json:"booking_date"`
	Next        *Booking  `json:"-"`
}

// Response structures for API
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type PageData struct {
	Title       string
	Description string
	Data        interface{}
}
