package model

import (
	"time"
)

// Global linked lists
var (
	CityHead     *City
	StationHead  *Station
	RouteHead    *Route
	ScheduleHead *Schedule
	BookingHead  *Booking
)

// City functions
func AddCity(name string) *City {
	newCity := &City{
		ID:   getNextCityID(),
		Name: name,
		Next: CityHead,
	}
	CityHead = newCity
	return newCity
}

func GetAllCities() []*City {
	var cities []*City
	current := CityHead
	for current != nil {
		cities = append(cities, current)
		current = current.Next
	}
	return cities
}

func GetCityByID(id int) *City {
	current := CityHead
	for current != nil {
		if current.ID == id {
			return current
		}
		current = current.Next
	}
	return nil
}

func DeleteCity(id int) bool {
	if CityHead == nil {
		return false
	}

	if CityHead.ID == id {
		CityHead = CityHead.Next
		return true
	}

	current := CityHead
	for current.Next != nil {
		if current.Next.ID == id {
			current.Next = current.Next.Next
			return true
		}
		current = current.Next
	}
	return false
}

// Station functions
func AddStation(name, address, phone string, cityID int) *Station {
	newStation := &Station{
		ID:      getNextStationID(),
		Name:    name,
		Address: address,
		Phone:   phone,
		CityID:  cityID,
		Next:    StationHead,
	}
	StationHead = newStation
	return newStation
}

func GetAllStations() []*Station {
	var stations []*Station
	current := StationHead
	for current != nil {
		stations = append(stations, current)
		current = current.Next
	}
	return stations
}

func GetStationByID(id int) *Station {
	current := StationHead
	for current != nil {
		if current.ID == id {
			return current
		}
		current = current.Next
	}
	return nil
}

func DeleteStation(id int) bool {
	if StationHead == nil {
		return false
	}

	if StationHead.ID == id {
		StationHead = StationHead.Next
		return true
	}

	current := StationHead
	for current.Next != nil {
		if current.Next.ID == id {
			current.Next = current.Next.Next
			return true
		}
		current = current.Next
	}
	return false
}

// Route functions
func AddRoute(fromCityID, toCityID int, distance float64) *Route {
	newRoute := &Route{
		ID:         getNextRouteID(),
		FromCityID: fromCityID,
		ToCityID:   toCityID,
		Distance:   distance,
		Next:       RouteHead,
	}
	RouteHead = newRoute
	return newRoute
}

func GetAllRoutes() []*Route {
	var routes []*Route
	current := RouteHead
	for current != nil {
		routes = append(routes, current)
		current = current.Next
	}
	return routes
}

func GetRouteByID(id int) *Route {
	current := RouteHead
	for current != nil {
		if current.ID == id {
			return current
		}
		current = current.Next
	}
	return nil
}

func DeleteRoute(id int) bool {
	if RouteHead == nil {
		return false
	}

	if RouteHead.ID == id {
		RouteHead = RouteHead.Next
		return true
	}

	current := RouteHead
	for current.Next != nil {
		if current.Next.ID == id {
			current.Next = current.Next.Next
			return true
		}
		current = current.Next
	}
	return false
}

// Schedule functions
func AddSchedule(routeID int, trainName string, departureTime, arrivalTime time.Time, price float64) *Schedule {
	newSchedule := &Schedule{
		ID:            getNextScheduleID(),
		RouteID:       routeID,
		TrainName:     trainName,
		DepartureTime: departureTime,
		ArrivalTime:   arrivalTime,
		Price:         price,
		Next:          ScheduleHead,
	}
	ScheduleHead = newSchedule
	return newSchedule
}

func GetAllSchedules() []*Schedule {
	var schedules []*Schedule
	current := ScheduleHead
	for current != nil {
		schedules = append(schedules, current)
		current = current.Next
	}
	return schedules
}

func GetScheduleByID(id int) *Schedule {
	current := ScheduleHead
	for current != nil {
		if current.ID == id {
			return current
		}
		current = current.Next
	}
	return nil
}

func DeleteSchedule(id int) bool {
	if ScheduleHead == nil {
		return false
	}

	if ScheduleHead.ID == id {
		ScheduleHead = ScheduleHead.Next
		return true
	}

	current := ScheduleHead
	for current.Next != nil {
		if current.Next.ID == id {
			current.Next = current.Next.Next
			return true
		}
		current = current.Next
	}
	return false
}

// Booking functions
func AddBooking(scheduleID int, userName, userEmail string, seats int, totalPrice float64) *Booking {
	newBooking := &Booking{
		ID:          getNextBookingID(),
		ScheduleID:  scheduleID,
		UserName:    userName,
		UserEmail:   userEmail,
		Seats:       seats,
		TotalPrice:  totalPrice,
		BookingDate: time.Now(),
		Next:        BookingHead,
	}
	BookingHead = newBooking
	return newBooking
}

func GetAllBookings() []*Booking {
	var bookings []*Booking
	current := BookingHead
	for current != nil {
		bookings = append(bookings, current)
		current = current.Next
	}
	return bookings
}

func GetBookingByID(id int) *Booking {
	current := BookingHead
	for current != nil {
		if current.ID == id {
			return current
		}
		current = current.Next
	}
	return nil
}

// Helper functions to get next IDs
func getNextCityID() int {
	maxID := 0
	current := CityHead
	for current != nil {
		if current.ID > maxID {
			maxID = current.ID
		}
		current = current.Next
	}
	return maxID + 1
}

func getNextStationID() int {
	maxID := 0
	current := StationHead
	for current != nil {
		if current.ID > maxID {
			maxID = current.ID
		}
		current = current.Next
	}
	return maxID + 1
}

func getNextRouteID() int {
	maxID := 0
	current := RouteHead
	for current != nil {
		if current.ID > maxID {
			maxID = current.ID
		}
		current = current.Next
	}
	return maxID + 1
}

func getNextScheduleID() int {
	maxID := 0
	current := ScheduleHead
	for current != nil {
		if current.ID > maxID {
			maxID = current.ID
		}
		current = current.Next
	}
	return maxID + 1
}

func getNextBookingID() int {
	maxID := 0
	current := BookingHead
	for current != nil {
		if current.ID > maxID {
			maxID = current.ID
		}
		current = current.Next
	}
	return maxID + 1
}

// Initialize sample data
func InitializeData() {
	// Add sample cities
	AddCity("Jakarta")
	AddCity("Bandung")
	AddCity("Surabaya")
	AddCity("Yogyakarta")
	AddCity("Semarang")

	// Add sample stations
	AddStation("Stasiun Gambir", "Jl. Gambir No.1, Gambir, Jakarta Pusat", "021-3848608", 1)
	AddStation("Stasiun Pasar Senen", "Jl. Pasar Senen, Senen, Jakarta Pusat", "021-4214682", 1)
	AddStation("Stasiun Bandung", "Jl. Stasiun Timur No.1, Bandung", "022-4230631", 2)
	AddStation("Stasiun Gubeng", "Jl. Gubeng Stasiun Utara, Surabaya", "031-5031088", 3)
	AddStation("Stasiun Tugu Yogyakarta", "Jl. Pringgodiningrat No.1, Yogyakarta", "0274-512404", 4)
	AddStation("Stasiun Semarang Tawang", "Jl. Taman Tawang No.1, Semarang", "024-3541862", 5)

	// Add sample routes
	AddRoute(1, 2, 150.0) // Jakarta - Bandung
	AddRoute(1, 3, 800.0) // Jakarta - Surabaya
	AddRoute(1, 4, 560.0) // Jakarta - Yogyakarta
	AddRoute(2, 4, 350.0) // Bandung - Yogyakarta
	AddRoute(4, 3, 320.0) // Yogyakarta - Surabaya

	// Add sample schedules
	departureTime1, _ := time.Parse("15:04", "08:00")
	arrivalTime1, _ := time.Parse("15:04", "11:00")
	AddSchedule(1, "Argo Parahyangan", departureTime1, arrivalTime1, 100000)

	departureTime2, _ := time.Parse("15:04", "19:00")
	arrivalTime2, _ := time.Parse("15:04", "07:00")
	AddSchedule(2, "Argo Bromo Anggrek", departureTime2, arrivalTime2, 350000)

	departureTime3, _ := time.Parse("15:04", "07:30")
	arrivalTime3, _ := time.Parse("15:04", "15:30")
	AddSchedule(3, "Taksaka", departureTime3, arrivalTime3, 200000)

	departureTime4, _ := time.Parse("15:04", "15:30")
	arrivalTime4, _ := time.Parse("15:04", "18:30")
	AddSchedule(1, "Parahyangan", departureTime4, arrivalTime4, 100000)
}
