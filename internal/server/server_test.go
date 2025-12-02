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
	srv, err := New(calendar, templateDir, "")
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

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
	srv, err := New(calendar, templateDir, "")
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

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
	srv, err := New(calendar, templateDir, "")
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Act: Request a locked day
	req := httptest.NewRequest(http.MethodGet, "/day/1", nil)
	rec := httptest.NewRecorder()
	srv.ServeHTTP(rec, req)

	// Assert: Should return 403 Forbidden
	if rec.Code != http.StatusForbidden {
		t.Errorf("Expected status 403 for locked day, got %d", rec.Code)
	}
}

func TestServer_HandleDay_InvalidDay_Zero(t *testing.T) {
	// Arrange: Create test templates
	templateDir := t.TempDir()
	dayTemplate := `<html><body>Day content</body></html>`
	err := os.WriteFile(templateDir+"/day.html", []byte(dayTemplate), 0644)
	if err != nil {
		t.Fatalf("Failed to create template: %v", err)
	}

	contentDir := t.TempDir()
	calendar := advent.NewCalendar(2020, contentDir)
	srv, err := New(calendar, templateDir, "")
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Act: Request day 0
	req := httptest.NewRequest(http.MethodGet, "/day/0", nil)
	rec := httptest.NewRecorder()
	srv.ServeHTTP(rec, req)

	// Assert: Should return 400 Bad Request
	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 for day 0, got %d", rec.Code)
	}
}

func TestServer_HandleDay_InvalidDay_Thirteen(t *testing.T) {
	// Arrange: Create test templates
	templateDir := t.TempDir()
	dayTemplate := `<html><body>Day content</body></html>`
	err := os.WriteFile(templateDir+"/day.html", []byte(dayTemplate), 0644)
	if err != nil {
		t.Fatalf("Failed to create template: %v", err)
	}

	contentDir := t.TempDir()
	calendar := advent.NewCalendar(2020, contentDir)
	srv, err := New(calendar, templateDir, "")
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Act: Request day 13 (out of bounds - calendar only has 12 days)
	req := httptest.NewRequest(http.MethodGet, "/day/13", nil)
	rec := httptest.NewRecorder()
	srv.ServeHTTP(rec, req)

	// Assert: Should return 400 Bad Request
	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 for day 13, got %d", rec.Code)
	}
}

func TestServer_HandleDay_InvalidDay_Negative(t *testing.T) {
	// Arrange: Create test templates
	templateDir := t.TempDir()
	dayTemplate := `<html><body>Day content</body></html>`
	err := os.WriteFile(templateDir+"/day.html", []byte(dayTemplate), 0644)
	if err != nil {
		t.Fatalf("Failed to create template: %v", err)
	}

	contentDir := t.TempDir()
	calendar := advent.NewCalendar(2020, contentDir)
	srv, err := New(calendar, templateDir, "")
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Act: Request day -1
	req := httptest.NewRequest(http.MethodGet, "/day/-1", nil)
	rec := httptest.NewRecorder()
	srv.ServeHTTP(rec, req)

	// Assert: Should return 400 Bad Request
	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 for day -1, got %d", rec.Code)
	}
}

func TestServer_HandleDay_InvalidDay_TooLarge(t *testing.T) {
	// Arrange: Create test templates
	templateDir := t.TempDir()
	dayTemplate := `<html><body>Day content</body></html>`
	err := os.WriteFile(templateDir+"/day.html", []byte(dayTemplate), 0644)
	if err != nil {
		t.Fatalf("Failed to create template: %v", err)
	}

	contentDir := t.TempDir()
	calendar := advent.NewCalendar(2020, contentDir)
	srv, err := New(calendar, templateDir, "")
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Act: Request day 100
	req := httptest.NewRequest(http.MethodGet, "/day/100", nil)
	rec := httptest.NewRecorder()
	srv.ServeHTTP(rec, req)

	// Assert: Should return 400 Bad Request
	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 for day 100, got %d", rec.Code)
	}
}

func TestServer_New_InvalidTemplate(t *testing.T) {
	// Arrange: Create a template with invalid syntax
	templateDir := t.TempDir()
	invalidTemplate := `<!DOCTYPE html>
<html>
<head><title>Test</title></head>
<body>
{{range .Days}}
  <div>{{.InvalidField}</div>
{{end
</body>
</html>`
	err := os.WriteFile(templateDir+"/index.html", []byte(invalidTemplate), 0644)
	if err != nil {
		t.Fatalf("Failed to create template: %v", err)
	}

	contentDir := t.TempDir()
	calendar := advent.NewCalendar(2020, contentDir)

	// Act: Try to create server with invalid template
	srv, err := New(calendar, templateDir, "")

	// Assert: Should return an error, not panic
	if err == nil {
		t.Error("Expected error for invalid template, got nil")
	}
	if srv != nil {
		t.Error("Expected nil server when template parsing fails")
	}
}
