package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Event struct {
	ID     int
	UserID int
	Date   time.Time
}

func (e Event) JSONify() []byte {
	dataJSON := fmt.Sprintf(
		"{\"id\": %d, \"user_id\": %d, \"date\":\"%s\"}",
		e.ID,
		e.UserID,
		e.Date.Format("2006-01-02"),
	)
	return []byte(dataJSON)
}

var db []*Event
var eID = 0

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("INFO: %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func withLogging(handler http.HandlerFunc) http.HandlerFunc {
	return LoggerMiddleware(handler).ServeHTTP
}

func main() {
	// POST
	http.HandleFunc("/create_event", withLogging(createEventHandler))
	http.HandleFunc("/update_event", withLogging(updateEventHandler))
	http.HandleFunc("/delete_event", withLogging(deleteEventHandler))
	// GET
	http.HandleFunc("/events_for_day", withLogging(dayEventHandler))
	http.HandleFunc("/events_for_week", withLogging(weekEventHandler))
	http.HandleFunc("/events_for_month", withLogging(monthEventHandler))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ CREATE EVENT
func createEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		data := "{\"error\": \"405 Method Not Allowed\"}"
		w.WriteHeader(405)
		w.Write([]byte(data))
		return
	}

	err := r.ParseForm()
	if err != nil {
		data := "{\"error\": \"500 Internal Server Error\"}"
		w.WriteHeader(500)
		w.Write([]byte(data))
		return
	}

	userID := r.FormValue("user_id")
	date := r.FormValue("date")

	event, err := validateCreateValues(userID, date)
	if err != nil {
		data := "{\"error\": \"400 Invalid Form Data\"}"
		w.WriteHeader(400)
		w.Write([]byte(data))
		return
	}
	createEventService(event)

	data := fmt.Sprintf(
		"{\"result\": %s}",
		string(event.JSONify()),
	)
	w.WriteHeader(201)
	w.Write([]byte(data))
}

func validateCreateValues(userID, date string) (*Event, error) {
	eventDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}
	eventUserID, err := strconv.Atoi(userID)
	if err != nil {
		return nil, err
	}
	if eventUserID < 0 {
		return nil, errors.New("Wrong User ID")
	}
	defer func() { eID++ }()
	return &Event{eID, eventUserID, eventDate}, nil
}

func createEventService(e *Event) {
	db = append(db, e)
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ UPDATE EVENT
func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		data := "{\"error\": \"405 Method Not Allowed\"}"
		w.WriteHeader(405)
		w.Write([]byte(data))
		return
	}

	err := r.ParseForm()
	if err != nil {
		data := "{\"error\": \"500 Internal Server Error\"}"
		w.WriteHeader(500)
		w.Write([]byte(data))
		return
	}

	id := r.FormValue("id")
	userID := r.FormValue("user_id")
	date := r.FormValue("date")

	event, err := validateUpdateValues(id, userID, date)
	if err != nil {
		data := "{\"error\": \"400 Invalid Form Data\"}"
		w.WriteHeader(400)
		w.Write([]byte(data))
		return
	}

	updateEventService(event)

	data := fmt.Sprintf(
		"{\"result\": %s}",
		string(event.JSONify()),
	)
	w.WriteHeader(200)
	w.Write([]byte(data))
}

func validateUpdateValues(id, userID, date string) (*Event, error) {
	eventDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}
	eventUserID, err := strconv.Atoi(userID)
	if err != nil {
		return nil, err
	}
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	if intID < 0 || intID >= eID {
		return nil, errors.New("Wrong Event ID")
	}
	if eventUserID < 0 {
		return nil, errors.New("Wrong User ID")
	}

	return &Event{intID, eventUserID, eventDate}, nil
}

func updateEventService(e *Event) {
	for i := range db {
		if db[i].ID == e.ID {
			db[i].UserID = e.UserID
			db[i].Date = e.Date
			break
		}
	}
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ DELETE EVENT
func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		data := "{\"error\": \"405 Method Not Allowed\"}"
		w.WriteHeader(405)
		w.Write([]byte(data))
		return
	}

	err := r.ParseForm()
	if err != nil {
		data := "{\"error\": \"500 Internal Server Error\"}"
		w.WriteHeader(500)
		w.Write([]byte(data))
		return
	}

	id := r.FormValue("id")

	deleteID, err := validateDeleteValues(id)
	if err != nil {
		data := "{\"error\": \"400 Invalid Form Data\"}"
		w.WriteHeader(400)
		w.Write([]byte(data))
		return
	}

	deleteEventService(deleteID)

	data := "{\"result\": \"200 Deleted\"}"
	w.WriteHeader(200)
	w.Write([]byte(data))
}

func validateDeleteValues(id string) (int, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	if intID < 0 || intID >= eID {
		return 0, errors.New("Wrong User ID")
	}
	return intID, nil
}

func deleteEventService(id int) {
	del := -1
	for i := range db {
		if db[i].ID == id {
			del = i
			break
		}
	}
	if del != -1 {
		db = append(db[:del], db[del+1:]...)
	}
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ EVENTS FOR DAY
func dayEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		data := "{\"error\": \"405 Method Not Allowed\"}"
		w.WriteHeader(405)
		w.Write([]byte(data))
		return
	}

	query := r.URL.Query()
	userID := query.Get("user_id")
	date := query.Get("date")

	event, err := validateQueryValues(userID, date)
	if err != nil {
		data := "{\"error\": \"400 Invalid Form Data\"}"
		w.WriteHeader(400)
		w.Write([]byte(data))
		return
	}

	events := dayEventService(event)
	data := fmt.Sprintf("{\"result\": [%s]}", events)

	w.WriteHeader(200)
	w.Write([]byte(data))
}

func dayEventService(eq *Event) string {
	res := []string{}

	for _, edb := range db {
		if eq.UserID == edb.UserID && eq.Date == edb.Date {
			res = append(res, string(edb.JSONify()))
		}
	}

	return strings.Join(res, ", ")
}

func validateQueryValues(userID, date string) (*Event, error) {
	eventDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}
	eventUserID, err := strconv.Atoi(userID)
	if err != nil {
		return nil, err
	}
	if eventUserID < 0 {
		return nil, errors.New("Wrong User ID")
	}
	return &Event{-1, eventUserID, eventDate}, nil
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ EVENTS FOR WEEK
func weekEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		data := "{\"error\": \"405 Method Not Allowed\"}"
		w.WriteHeader(405)
		w.Write([]byte(data))
		return
	}

	query := r.URL.Query()
	userID := query.Get("user_id")
	date := query.Get("date")

	event, err := validateQueryValues(userID, date)
	if err != nil {
		data := "{\"error\": \"400 Invalid Form Data\"}"
		w.WriteHeader(400)
		w.Write([]byte(data))
		return
	}

	events := weekEventService(event)
	data := fmt.Sprintf("{\"result\": [%s]}", events)

	w.WriteHeader(200)
	w.Write([]byte(data))
}

func weekEventService(eq *Event) string {
	res := []string{}

	for _, edb := range db {
		dif := eq.Date.Sub(edb.Date) / 24.0
		if eq.UserID == edb.UserID && dif <= 7 && dif >= 0 {
			res = append(res, string(edb.JSONify()))
		}
	}

	return strings.Join(res, ", ")
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ EVENTS FOR MONTH
func monthEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		data := "{\"error\": \"405 Method Not Allowed\"}"
		w.WriteHeader(405)
		w.Write([]byte(data))
		return
	}

	query := r.URL.Query()
	userID := query.Get("user_id")
	date := query.Get("date")

	event, err := validateQueryValues(userID, date)
	if err != nil {
		data := "{\"error\": \"400 Invalid Form Data\"}"
		w.WriteHeader(400)
		w.Write([]byte(data))
		return
	}

	events := monthEventService(event)
	data := fmt.Sprintf("{\"result\": [%s]}", events)

	w.WriteHeader(200)
	w.Write([]byte(data))
}

func monthEventService(eq *Event) string {
	res := []string{}

	for _, edb := range db {
		dif := eq.Date.Sub(edb.Date) / 24.0
		if eq.UserID == edb.UserID && dif <= 30 && dif >= 0 {
			res = append(res, string(edb.JSONify()))
		}
	}

	return strings.Join(res, ", ")
}
