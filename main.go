package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/kilianp07/Calendar-Cli/calib"
)

func main() {
	http.HandleFunc("/json-calendar", JSONCalendarHandler)
	http.HandleFunc("/html-calendar", HTMLCalendarHandler)

	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}

func JSONCalendarHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	monthStr := r.URL.Query().Get("month")
	yearStr := r.URL.Query().Get("year")

	// Default to the current month and year if not provided
	if monthStr == "" || yearStr == "" {
		now := time.Now()
		monthStr = strconv.Itoa(int(now.Month()))
		yearStr = strconv.Itoa(now.Year())
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil {
		http.Error(w, "Invalid month", http.StatusBadRequest)
		return
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		http.Error(w, "Invalid year", http.StatusBadRequest)
		return
	}

	calendar := calib.NewCalendar(month, year)

	// Convert calendar to JSON
	jsonData, err := calendar.ToJSON()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, jsonData)
}

func HTMLCalendarHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	monthStr := r.URL.Query().Get("month")
	yearStr := r.URL.Query().Get("year")

	// Default to the current month and year if not provided
	if monthStr == "" || yearStr == "" {
		now := time.Now()
		monthStr = strconv.Itoa(int(now.Month()))
		yearStr = strconv.Itoa(now.Year())
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil {
		http.Error(w, "Invalid month", http.StatusBadRequest)
		return
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		http.Error(w, "Invalid year", http.StatusBadRequest)
		return
	}

	calendar := calib.NewCalendar(month, year)

	// Convert calendar to HTML
	html, err := calendar.ToHtml()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}
