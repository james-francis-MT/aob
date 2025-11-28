package advent

import (
	"fmt"
	"os"
	"time"
)

// Day represents a single day in the advent calendar
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
	filePath := fmt.Sprintf("%s/day%d.txt", contentDir, d.Number)
	content, err := os.ReadFile(filePath)
	if err != nil {
		// File doesn't exist, use default content
		d.Content = "Happy Advent! 🎄"
		return nil
	}
	d.Content = string(content)
	return nil
}

// Calendar represents an advent calendar with 25 days
type Calendar struct {
	Days []Day
}

// NewCalendar creates a new advent calendar for the specified year
func NewCalendar(year int, contentDir string) *Calendar {
	days := make([]Day, 25)
	for i := 0; i < 25; i++ {
		dayNum := i + 1
		date := time.Date(year, time.December, dayNum, 0, 0, 0, 0, time.Local)
		days[i] = Day{
			Number: dayNum,
			Date:   date,
		}
		days[i].CheckUnlocked()
		days[i].LoadContent(contentDir)
	}
	return &Calendar{Days: days}
}
