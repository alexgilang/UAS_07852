package controller

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"../model"
)

// Admin Cities Handler
func AdminCitiesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		cityName := r.FormValue("city_name")
		if cityName != "" {
			model.AddCity(cityName)
		}
		http.Redirect(w, r, "/admin/cities", http.StatusSeeOther)
		return
	}

	cities := model.GetAllCities()
	data := model.PageData{
		Title:       "Kelola Kota - Admin",
		Description: "Manajemen Data Kota",
		Data:        cities,
	}

	tmpl, err := template.ParseFiles("templates/admin_cities.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

// Admin Stations Handler
func AdminStationsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		stationName := r.FormValue("station_name")
		address := r.FormValue("address")
		phone := r.FormValue("phone")
		cityIDStr := r.FormValue("city_id")

		if stationName != "" && address != "" && phone != "" && cityIDStr != "" {
			cityID, err := strconv.Atoi(cityIDStr)
			if err == nil {
				model.AddStation(stationName, address, phone, cityID)
			}
		}
		http.Redirect(w, r, "/admin/stations", http.StatusSeeOther)
		return
	}

	stations := model.GetAllStations()
	cities := model.GetAllCities()
	
	data := struct {
		Title       string
		Description string
		Stations    []*model.Station
		Cities      []*model.City
	}{
		Title:       "Kelola Stasiun - Admin",
		Description: "Manajemen Data Stasiun",
		Stations:    stations,
		Cities:      cities,
	}

	tmpl, err := template.ParseFiles("templates/admin_stations.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

// Admin Routes Handler
func AdminRoutesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fromCityIDStr := r.FormValue("from_city_id")
		toCityIDStr := r.FormValue("to_city_id")
		distanceStr := r.FormValue("distance")

		if fromCityIDStr != "" && toCityIDStr != "" && distanceStr != "" {
			fromCityID, err1 := strconv.Atoi(fromCityIDStr)
			toCityID, err2 := strconv.Atoi(toCityIDStr)
			distance, err3 := strconv.ParseFloat(distanceStr, 64)

			if err1 == nil && err2 == nil && err3 == nil {
				model.AddRoute(fromCityID, toCityID, distance)
			}
		}
		http.Redirect(w, r, "/admin/routes", http.StatusSeeOther)
		return
	}

	routes := model.GetAllRoutes()
	cities := model.GetAllCities()

	data := struct {
		Title       string
		Description string
		Routes      []*model.Route
		Cities      []*model.City
	}{
		Title:       "Kelola Rute - Admin",
		Description: "Manajemen Data Rute",
		Routes:      routes,
		Cities:      cities,
	}

	tmpl, err := template.ParseFiles("templates/admin_routes.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

// Admin Schedules Handler
func AdminSchedulesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		routeIDStr := r.FormValue("route_id")
		trainName := r.FormValue("train_name")
		departureTimeStr := r.FormValue("departure_time")
		arrivalTimeStr := r.FormValue("arrival_time")
		priceStr := r.FormValue("price")

		if routeIDStr != "" && trainName != "" && departureTimeStr != "" && arrivalTimeStr != "" && priceStr != "" {
			routeID, err1 := strconv.Atoi(routeIDStr)
			departureTime, err2 := time.Parse("15:04", departureTimeStr)
			arrivalTime, err3 := time.Parse("15:04", arrivalTimeStr)
			price, err4 := strconv.ParseFloat(priceStr, 64)

			if err1 == nil && err2 == nil && err3 == nil && err4 == nil {
				model.AddSchedule(routeID, trainName, departureTime, arrivalTime, price)
			}
		}
		http.Redirect(w, r, "/admin/schedules", http.StatusSeeOther)
		return
	}

	schedules := model.GetAllSchedules()
	routes := model.GetAllRoutes()

	data := struct {
		Title       string
		Description string
		Schedules   []*model.Schedule
		Routes      []*model.Route
	}{
		Title:       "Kelola Jadwal - Admin",
		Description: "Manajemen Data Jadwal",
		Schedules:   schedules,
		Routes:      routes,
	}

	tmpl, err := template.ParseFiles("templates/admin_schedules.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

// Admin Prices Handler
func AdminPricesHandler(w http.ResponseWriter, r *http.Request) {
	schedules := model.GetAllSchedules()

	data := model.PageData{
		Title:       "Kelola Harga - Admin",
		Description: "Daftar Harga Tiket",
		Data:        schedules,
	}

	tmpl, err := template.ParseFiles("templates/admin_prices.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}
