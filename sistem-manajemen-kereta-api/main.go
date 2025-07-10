package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"./controller"
	"./model"
)

func main() {
	// Initialize data
	model.InitializeData()

	// Static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("templates/"))))

	// Routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/admin", adminHandler)
	http.HandleFunc("/user", userHandler)

	// Admin routes
	http.HandleFunc("/admin/cities", controller.AdminCitiesHandler)
	http.HandleFunc("/admin/stations", controller.AdminStationsHandler)
	http.HandleFunc("/admin/routes", controller.AdminRoutesHandler)
	http.HandleFunc("/admin/schedules", controller.AdminSchedulesHandler)
	http.HandleFunc("/admin/prices", controller.AdminPricesHandler)

	// User routes
	http.HandleFunc("/user/stations", controller.UserStationsHandler)
	http.HandleFunc("/user/routes", controller.UserRoutesHandler)
	http.HandleFunc("/user/booking", controller.UserBookingHandler)

	fmt.Println("🚄 Server started at http://localhost:8080")
	fmt.Println("📊 Admin Panel: http://localhost:8080/admin")
	fmt.Println("👤 User Panel: http://localhost:8080/user")
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/admin.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/user.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
