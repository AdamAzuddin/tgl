# toGoList

A simple and efficient CLI todo list application built with Go.

## Features

- ✅ Add new tasks
- 📋 List all tasks
- ✔️ Mark tasks as complete
- 🗑️ Delete tasks
- 💾 Persistent storage

## Installation

```bash
# Clone the repository
git clone https://github.com/AdamAzuddin/ToGoList.git
cd ToGoList

# Build the application
go build -o togolist main.go

# Or run directly
go run main.go
```

## Usage

```bash
# Add a new task
./togolist add "Buy groceries"

# List all tasks
./togolist list

# Complete a task
./togolist complete 1

# Delete a task
./togolist delete 1

# Show help
./togolist help
```

## Project Structure

```
ToGoList/
├── cmd/                 # CLI commands
├── internal/            # Private application code
│   ├── todo/           # Todo business logic
│   └── storage/        # Data persistence
├── pkg/                # Public packages
├── main.go             # Application entry point
├── go.mod              # Go module file
└── README.md           # Project documentation
```

## Development

```bash
# Run tests
go test ./...

# Format code
go fmt ./...

# Run linter
golangci-lint run
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is open source and available under the [MIT License](LICENSE).