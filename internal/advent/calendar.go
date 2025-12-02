package advent

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Day represents a single day in the 12-day countdown to Christmas
type Day struct {
	Number   int
	Date     time.Time
	Unlocked bool
	Content  string
}

// CheckUnlocked updates the Unlocked field based on whether the day's date has arrived
func (d *Day) CheckUnlocked() {
	d.Unlocked = time.Now().After(d.Date) || time.Now().Equal(d.Date)
}

// LoadContent loads the content for this day from the specified directory
func (d *Day) LoadContent(contentDir string) error {
	// Validate day number to prevent path traversal and invalid access
	if d.Number < 1 || d.Number > 12 {
		return fmt.Errorf("invalid day number: %d (must be between 1 and 12)", d.Number)
	}

	// Use filepath.Join for safe path construction
	filePath := filepath.Join(contentDir, fmt.Sprintf("day%d.txt", d.Number))

	// Validate that the resulting path is within contentDir
	absContentDir, err := filepath.Abs(contentDir)
	if err != nil {
		return fmt.Errorf("failed to resolve content directory: %w", err)
	}
	absFilePath, err := filepath.Abs(filePath)
	if err != nil {
		return fmt.Errorf("failed to resolve file path: %w", err)
	}

	// Check that the file path is within the content directory
	relPath, err := filepath.Rel(absContentDir, absFilePath)
	if err != nil || len(relPath) > 0 && relPath[0] == '.' && relPath[1] == '.' {
		return fmt.Errorf("invalid file path: attempted directory traversal")
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		// File doesn't exist, use default content
		d.Content = "Merry Christmas! 🎄"
		return nil
	}
	d.Content = string(content)
	return nil
}

// Calendar represents a 12-day countdown calendar to Christmas
type Calendar struct {
	Days []Day
}

// NewCalendar creates a new 12-day countdown calendar for the specified year
// Day 1 starts on December 13th, Day 12 ends on December 24th (Christmas Eve)
func NewCalendar(year int, contentDir string) *Calendar {
	days := make([]Day, 12)
	for i := 0; i < 12; i++ {
		dayNum := i + 1
		// Day 1 = December 13, Day 2 = December 14, ..., Day 12 = December 24
		decemberDay := i + 13
		date := time.Date(year, time.December, decemberDay, 0, 0, 0, 0, time.Local)
		days[i] = Day{
			Number: dayNum,
			Date:   date,
		}
		days[i].CheckUnlocked()
		// LoadContent should never error here since dayNum is always 1-12
		// Error is only returned for validation failures with invalid day numbers
		_ = days[i].LoadContent(contentDir)
	}
	return &Calendar{Days: days}
}
