package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/James-Francis-MT/aob/internal/advent"
)

func TestServer_HandleHome(t *testing.T) {
	// Arrange: Create test templates
	templateDir := t.TempDir()
	indexTemplate := `<!DOCTYPE html>
<html>
<head><title>Advent Calendar</title></head>
<body>
{{range .Days}}
<div>Day {{.Number}}</div>
{{end}}
</body>
</html>`
	err := os.WriteFile(templateDir+"/index.html", []byte(indexTemplate), 0644)
	if err != nil {
		t.Fatalf("Failed to create template: %v", err)
	}

	// Arrange: Create a test calendar
	contentDir := t.TempDir()
	calendar := advent.NewCalendar(time.Now().Year(), contentDir)
	srv := New(calendar, templateDir, "")

	// Act: Make a request to the home page
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	srv.ServeHTTP(rec, req)

	// Assert: Response should be OK
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	// Assert: Response should contain rendered HTML
	body := rec.Body.String()
	if len(body) == 0 {
		t.Error("Expected non-empty response body")
	}
	if !strings.Contains(body, "Day 1") {
		t.Error("Expected response to contain calendar days")
	}
}

func TestServer_HandleDay_Unlocked(t *testing.T) {
	// Arrange: Create test templates
	templateDir := t.TempDir()
	dayTemplate := `<!DOCTYPE html>
<html>
<head><title>Day {{.Number}}</title></head>
<body>
<p>{{.Content}}</p>
</body>
</html>`
	err := os.WriteFile(templateDir+"/day.html", []byte(dayTemplate), 0644)
	if err != nil {
		t.Fatalf("Failed to create template: %v", err)
	}

	// Arrange: Create a calendar with a day in the past (unlocked)
	contentDir := t.TempDir()
	calendar := advent.NewCalendar(2020, contentDir) // Past year, all days unlocked
	srv := New(calendar, templateDir, "")

	// Act: Request an unlocked day
	req := httptest.NewRequest(http.MethodGet, "/day/1", nil)
	rec := httptest.NewRecorder()
	srv.ServeHTTP(rec, req)

	// Assert: Should return 200 OK
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200 for unlocked day, got %d", rec.Code)
	}
}

func TestServer_HandleDay_Locked(t *testing.T) {
	// Arrange: Create test templates
	templateDir := t.TempDir()
	dayTemplate := `<html><body>Day content</body></html>`
	err := os.WriteFile(templateDir+"/day.html", []byte(dayTemplate), 0644)
	if err != nil {
		t.Fatalf("Failed to create template: %v", err)
	}

	// Arrange: Create a calendar with future days (locked)
	contentDir := t.TempDir()
	calendar := advent.NewCalendar(2099, contentDir) // Future year, all days locked
	srv := New(calendar, templateDir, "")

	// Act: Request a locked day
	req := httptest.NewRequest(http.MethodGet, "/day/1", nil)
	rec := httptest.NewRecorder()
	srv.ServeHTTP(rec, req)

	// Assert: Should return 403 Forbidden
	if rec.Code != http.StatusForbidden {
		t.Errorf("Expected status 403 for locked day, got %d", rec.Code)
	}
}
