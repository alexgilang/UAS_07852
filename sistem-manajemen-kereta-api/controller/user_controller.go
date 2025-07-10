package controller

import (
	"html/template"
	"net/http"
	"strconv"

	"../model"
)

// User Stations Handler
func UserStationsHandler(w http.ResponseWriter, r *http.Request) {
	stations := model.GetAllStations()

	data := model.PageData{
		Title:       "Daftar Stasiun - User",
		Description: "Informasi Lengkap Stasiun Kereta API",
		Data:        stations,
	}

	tmpl, err := template.ParseFiles("templates/user_stations.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

// User Routes Handler
func UserRoutesHandler(w http.ResponseWriter, r *http.Request) {
	schedules := model.GetAllSchedules()

	data := model.PageData{
		Title:       "Rute & Jadwal - User",
		Description: "Informasi jadwal dan rute perjalanan kereta api",
		Data:        schedules,
	}

	tmpl, err := template.ParseFiles("templates/user_routes.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

// User Booking Handler
func UserBookingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		scheduleIDStr := r.FormValue("schedule_id")
		userName := r.FormValue("user_name")
		userEmail := r.FormValue("user_email")
		seatsStr := r.FormValue("seats")

		if scheduleIDStr != "" && userName != "" && userEmail != "" && seatsStr != "" {
			scheduleID, err1 := strconv.Atoi(scheduleIDStr)
			seats, err2 := strconv.Atoi(seatsStr)

			if err1 == nil && err2 == nil {
				// Find schedule to get price
				schedules := model.GetAllSchedules()
				for _, schedule := range schedules {
					if schedule.ID == scheduleID {
						totalPrice := schedule.Price * float64(seats)
						model.AddBooking(scheduleID, userName, userEmail, seats, totalPrice)
						break
					}
				}
			}
		}
		http.Redirect(w, r, "/user/booking?success=1", http.StatusSeeOther)
		return
	}

	schedules := model.GetAllSchedules()
	success := r.URL.Query().Get("success")

	data := struct {
		Title       string
		Description string
		Schedules   []*model.Schedule
		Success     bool
	}{
		Title:       "Pesan Tiket - User",
		Description: "Pesan tiket kereta api dengan mudah dan cepat",
		Schedules:   schedules,
		Success:     success == "1",
	}

	tmpl, err := template.ParseFiles("templates/user_booking.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}
