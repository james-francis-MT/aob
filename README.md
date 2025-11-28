# Advent of Beckie (AOB)

A personal advent calendar website built for Beckie, where doors unlock each day in December to reveal personalized content.

## Features

- 25 interactive doors (one for each day of December)
- Doors automatically unlock on their designated day
- Personalized content for each day
- Responsive design with festive styling
- Built with Go's standard library only

## Project Structure

```
aob/
├── cmd/aob/              # Application entry point
│   └── main.go
├── internal/
│   ├── advent/           # Calendar and day logic
│   │   ├── calendar.go
│   │   └── calendar_test.go
│   └── server/           # HTTP server and handlers
│       ├── server.go
│       └── server_test.go
├── templates/            # HTML templates
│   ├── index.html        # Main calendar view
│   └── day.html          # Individual day view
├── static/               # Static assets
│   ├── css/
│   │   └── style.css
│   └── js/
├── content/              # Day content files
│   ├── day1.txt
│   ├── day2.txt
│   └── ...
├── Makefile
├── go.mod
└── README.md
```

## Getting Started

### Prerequisites

- Go 1.24 or higher

### Running the Application

```bash
# Run directly
make run

# Or build and run
make build
./bin/aob
```

The server will start on `http://localhost:8080`

### Adding Content

Create text files in the `content/` directory named `day1.txt` through `day25.txt`. Each file should contain the message or content you want to display when that door is opened.

Example:
```bash
echo "Happy December 1st! ❤️" > content/day1.txt
```

## Development

```bash
# Run the server
make run

# Run tests
make test

# Run tests with coverage
make test-coverage

# Format code
make fmt

# Clean build artifacts
make clean
```

## Testing

This project was built using **Test-Driven Development (TDD)**. All core functionality has comprehensive test coverage:

- **advent package**: 100% test coverage
  - Day unlocking logic (past, future, and current dates)
  - Calendar creation and initialization
  - Content loading from files with fallback defaults

- **server package**: 75% test coverage
  - Home page rendering
  - Day detail page for unlocked days
  - 403 Forbidden enforcement for locked days

### Running Tests

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run with coverage report
go test -cover ./...

# Run tests for specific package
go test ./internal/advent/
```

### Test Philosophy

Following the **Red-Green-Refactor** cycle:
1. **RED**: Write a failing test first
2. **GREEN**: Write minimal code to make it pass
3. **REFACTOR**: Improve code while keeping tests green

## How It Works

1. The calendar checks the current date and unlocks doors up to today
2. Each door corresponds to a day in December (1-25)
3. Content is loaded from `content/dayN.txt` files
4. Users can click unlocked doors to view the content
5. Locked doors display a lock icon and can't be opened

## Customization

- Edit `templates/*.html` to change the layout
- Modify `static/css/style.css` to adjust styling
- Add images or other media to the `static/` directory
- Update content files in `content/` directory

## Technical Details

- **Module**: `github.com/James-Francis-MT/aob`
- **Dependencies**: None - uses only Go's standard library
- **Go Version**: 1.24.2
- **Architecture**: Clean separation between business logic (`internal/advent`) and HTTP layer (`internal/server`)
- **Testing**: Comprehensive test suite with 100% coverage on core logic

## License

Personal project - built with ❤️ for Beckie
