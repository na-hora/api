package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type AppointmentHandlerInterface interface {
	SseUpdates(w http.ResponseWriter, r *http.Request)
}

type AppointmentHandler struct {
}

func GetAppointmentHandler() AppointmentHandlerInterface {
	return &AppointmentHandler{}
}

func (ah *AppointmentHandler) SseUpdates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	ctx := r.Context()

	go func() {
		<-ctx.Done()
		fmt.Println("Client disconnected, stopping SSE")
	}()

	type Appointment struct {
		Name  string
		Start time.Time
		End   time.Time
	}

	appointmentsTest := []Appointment{
		{
			Name:  "Meeting with John",
			Start: time.Now().Add(1 * time.Hour),
			End:   time.Now().Add(2 * time.Hour),
		},
		{
			Name:  "Dentist Appointment",
			Start: time.Now().Add(3 * time.Hour),
			End:   time.Now().Add(4 * time.Hour),
		},
		{
			Name:  "Project Deadline",
			Start: time.Now().Add(5 * time.Hour),
			End:   time.Now().Add(6 * time.Hour),
		},
	}

	for _, appointment := range appointmentsTest {
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)

		if err := enc.Encode(appointment); err != nil {
			http.Error(w, "Failed to encode appointment", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "data: %v\n\n", buf.String())
		fmt.Printf("data: %v\n", buf.String())

		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}

		time.Sleep(5 * time.Second)
	}
}
