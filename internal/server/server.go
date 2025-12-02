package server

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/James-Francis-MT/aob/internal/advent"
)

// Server handles HTTP requests
type Server struct {
	calendar    *advent.Calendar
	templateDir string
	staticDir   string
	templates   *template.Template
	mux         *http.ServeMux
}

// New creates a new server
func New(calendar *advent.Calendar, templateDir string, staticDir string) (*Server, error) {
	s := &Server{
		calendar:    calendar,
		templateDir: templateDir,
		staticDir:   staticDir,
		mux:         http.NewServeMux(),
	}

	// Parse templates
	templates, err := template.ParseGlob(templateDir + "/*.html")
	if err != nil {
		return nil, err
	}
	s.templates = templates

	// Register routes
	s.mux.HandleFunc("/", s.handleHome)
	s.mux.HandleFunc("/day/", s.handleDay)

	// Static file serving
	if staticDir != "" {
		fileServer := http.FileServer(http.Dir(staticDir))
		s.mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	}

	return s, nil
}

// ServeHTTP implements http.Handler
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

// handleHome renders the main calendar page
func (s *Server) handleHome(w http.ResponseWriter, r *http.Request) {
	err := s.templates.ExecuteTemplate(w, "index.html", s.calendar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// handleDay renders a single day's content
func (s *Server) handleDay(w http.ResponseWriter, r *http.Request) {
	// Extract day number from URL path
	path := strings.TrimPrefix(r.URL.Path, "/day/")
	dayNum, err := strconv.Atoi(path)
	if err != nil || dayNum < 1 || dayNum > 12 {
		http.Error(w, "Invalid day number", http.StatusBadRequest)
		return
	}

	// Defensive bounds check before array access
	if dayNum-1 >= len(s.calendar.Days) || dayNum-1 < 0 {
		http.Error(w, "Invalid day number", http.StatusBadRequest)
		return
	}

	// Get the day from the calendar
	day := s.calendar.Days[dayNum-1]

	// Check if the day is unlocked
	if !day.Unlocked {
		http.Error(w, "This day is not yet unlocked", http.StatusForbidden)
		return
	}

	// Render the day template
	err = s.templates.ExecuteTemplate(w, "day.html", day)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
