package advent

import (
	"fmt"
	"os"
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
	filePath := fmt.Sprintf("%s/day%d.txt", contentDir, d.Number)
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
		days[i].LoadContent(contentDir)
	}
	return &Calendar{Days: days}
}
