package advent

import (
	"os"
	"testing"
	"time"
)

func TestDay_IsUnlocked_PastDate(t *testing.T) {
	// Arrange: Create a day in the past (yesterday)
	yesterday := time.Now().AddDate(0, 0, -1)
	day := Day{
		Number:   1,
		Date:     yesterday,
		Unlocked: false,
	}

	// Act: Check if the day is unlocked
	day.CheckUnlocked()

	// Assert: Day should be unlocked since it's in the past
	if !day.Unlocked {
		t.Errorf("Expected day from %v to be unlocked, but it was locked", yesterday)
	}
}

func TestDay_IsUnlocked_FutureDate(t *testing.T) {
	// Arrange: Create a day in the future (tomorrow)
	tomorrow := time.Now().AddDate(0, 0, 1)
	day := Day{
		Number:   1,
		Date:     tomorrow,
		Unlocked: false,
	}

	// Act: Check if the day is unlocked
	day.CheckUnlocked()

	// Assert: Day should remain locked since it's in the future
	if day.Unlocked {
		t.Errorf("Expected day from %v to be locked, but it was unlocked", tomorrow)
	}
}

func TestDay_IsUnlocked_Today(t *testing.T) {
	// Arrange: Create a day for today (1 hour ago to ensure it's definitely passed)
	oneHourAgo := time.Now().Add(-1 * time.Hour)
	day := Day{
		Number:   1,
		Date:     oneHourAgo,
		Unlocked: false,
	}

	// Act: Check if the day is unlocked
	day.CheckUnlocked()

	// Assert: Day should be unlocked since the time has passed
	if !day.Unlocked {
		t.Errorf("Expected day from %v to be unlocked, but it was locked", oneHourAgo)
	}
}

func TestCalendar_New(t *testing.T) {
	// Arrange: Specify the year and create test content directory
	year := 2024
	contentDir := t.TempDir()

	// Create a test content file for day 1
	testContent := "Special content for day 1"
	err := os.WriteFile(contentDir+"/day1.txt", []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Act: Create a new calendar
	calendar := NewCalendar(year, contentDir)

	// Assert: Calendar should have 25 days
	if len(calendar.Days) != 25 {
		t.Errorf("Expected calendar to have 25 days, got %d", len(calendar.Days))
	}

	// Assert: Each day should have the correct date
	for i, day := range calendar.Days {
		expectedDayNum := i + 1
		expectedDate := time.Date(year, time.December, expectedDayNum, 0, 0, 0, 0, time.Local)

		if day.Number != expectedDayNum {
			t.Errorf("Day %d: Expected Number to be %d, got %d", i, expectedDayNum, day.Number)
		}

		if !day.Date.Equal(expectedDate) {
			t.Errorf("Day %d: Expected Date to be %v, got %v", i, expectedDate, day.Date)
		}
	}

	// Assert: Day 1 should have loaded content
	if calendar.Days[0].Content != testContent {
		t.Errorf("Expected day 1 content '%s', got '%s'", testContent, calendar.Days[0].Content)
	}

	// Assert: Day 2 should have default content (no file created)
	expectedDefault := "Happy Advent! 🎄"
	if calendar.Days[1].Content != expectedDefault {
		t.Errorf("Expected day 2 content '%s', got '%s'", expectedDefault, calendar.Days[1].Content)
	}
}

func TestDay_LoadContent_FileExists(t *testing.T) {
	// Arrange: Create a temporary content file
	contentDir := t.TempDir()
	testContent := "Test content for day 1"
	filePath := contentDir + "/day1.txt"
	err := os.WriteFile(filePath, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	day := Day{Number: 1}

	// Act: Load content
	err = day.LoadContent(contentDir)

	// Assert: Content should be loaded
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if day.Content != testContent {
		t.Errorf("Expected content '%s', got '%s'", testContent, day.Content)
	}
}

func TestDay_LoadContent_FileNotExists(t *testing.T) {
	// Arrange: Use a temporary empty directory
	contentDir := t.TempDir()
	day := Day{Number: 1}

	// Act: Load content from non-existent file
	err := day.LoadContent(contentDir)

	// Assert: Should use default content
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	expectedDefault := "Happy Advent! 🎄"
	if day.Content != expectedDefault {
		t.Errorf("Expected default content '%s', got '%s'", expectedDefault, day.Content)
	}
}
