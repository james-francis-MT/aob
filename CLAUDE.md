# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a personal advent calendar website built for a girlfriend. Users click doors to open them each day in December, revealing personalized content. The application uses only Go's standard library (no external frameworks).

## Build and Run Commands

```bash
# Run the application (starts on :8080)
make run
# or: go run cmd/aob/main.go

# Build binary
make build

# Run tests
make test

# Run with coverage
make test-coverage

# Format code
make fmt

# Clean artifacts
make clean
```

## Architecture

### Core Components

**Entry Point (`cmd/aob/main.go`)**
- Initializes the server
- Starts HTTP listener on port 8080
- Minimal logic - delegates to internal packages

**Server Layer (`internal/server/server.go`)**
- HTTP routing and request handling
- Template rendering for HTML responses
- Routes:
  - `GET /` - Main calendar grid view
  - `GET /day/{N}` - Individual day content (only if unlocked)
  - `GET /static/*` - Static file serving
- Enforces unlocking logic: returns 403 if user tries to access locked days

**Calendar Logic (`internal/advent/calendar.go`)**
- `Calendar` struct: Holds 25 Day objects for December
- `Day` struct: Number, Date, Unlocked status, Content
- Unlocking logic: Compares current time with day's date (using `time.Local`)
- Content loading: Reads from `content/dayN.txt` files, defaults to generic message if file missing

### Data Flow

1. Server initializes Calendar on startup with current year
2. For each day 1-25:
   - Creates Day with December date
   - Checks if `time.Now()` is after/equal to day's date
   - Loads content from `content/dayN.txt` (if exists)
3. Home page displays all 25 doors with lock/unlock state
4. Clicking unlocked door navigates to `/day/N`
5. Day handler validates unlock status before showing content

### Template System

- Uses Go's `html/template` package
- Templates parsed on server startup from `templates/*.html`
- **index.html**: Renders calendar grid, iterates over Days, conditionally shows lock icons
- **day.html**: Displays single day's content with back link

### Static Assets

- CSS in `static/css/style.css` - festive gradient background, door styling, animations
- JavaScript directory exists but currently unused
- Static files served via `http.FileServer`

## Development Patterns

### Adding New Features

**Adding new routes:**
- Add handler method to Server struct in `internal/server/server.go`
- Register route in `New()` function
- Keep handlers focused on HTTP concerns, delegate logic to `internal/advent` or new packages

**Modifying calendar logic:**
- Edit `internal/advent/calendar.go`
- Don't change unlocking logic without considering timezone implications
- Content loading happens once at startup - restart needed after changing content files

**Template changes:**
- Edit templates in `templates/`
- Server must be restarted to pick up template changes
- Use `{{range}}`, `{{if}}` for logic in templates
- Data passed via structs defined in handler functions

### Content Management

- Each day's content stored in `content/day1.txt` through `content/day25.txt`
- Content is plain text, loaded via `os.ReadFile`
- Missing files default to "Happy Advent! 🎄"
- Content is cached in memory at startup (not reloaded per request)
- To support rich content (images, HTML), modify `Day.Content` type and template rendering

### Testing Approach

- Test calendar unlocking logic across different dates
- Mock `time.Now()` for deterministic tests
- Test HTTP handlers with `httptest.ResponseRecorder`
- Test template rendering with sample data

## Important Constraints

- **No external dependencies**: Uses only Go standard library
- **Timezone**: Calendar uses `time.Local` - ensure server timezone is correct
- **Personal project**: Features prioritize simplicity over scalability
- **Content security**: Content loaded from filesystem, no database or user input validation needed
- **Static content**: No admin interface, content managed via text files

## Common Modifications

**Change unlock time to midnight:**
Currently unlocks at midnight UTC. To change timezone, modify date creation in `calendar.go`:
```go
date := time.Date(year, time.December, dayNum, 0, 0, 0, 0, time.Local)
```

**Add images to doors:**
1. Add image path to Day struct
2. Store images in `static/images/`
3. Update template to render images
4. Modify content loading to read metadata

**Enable rich HTML content:**
1. Change `Content` type from `string` to `template.HTML`
2. Update `loadDayContent` to read `.html` files
3. Update `day.html` template to render unescaped HTML

**Add authentication:**
This is a personal gift site, but to restrict access:
1. Add middleware to Server for basic auth
2. Store password hash in environment variable
3. Wrap handlers with auth check

## Module Path

Update the module path in `go.mod` from `github.com/yourusername/aob` to your actual repository path before publishing.
