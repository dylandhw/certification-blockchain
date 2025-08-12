# Certification Blockchain

A tamper-proof system for recording and verifying event attendance and certifications using blockchain technology. Built with Go and featuring a modern web interface.

![Go Version](https://img.shields.io/badge/Go-1.22.2+-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)
![Status](https://img.shields.io/badge/Status-Development-orange.svg)

## Features

- **Immutable Blockchain**: Cryptographic verification ensures data integrity
- **Web Interface**: Clean, responsive forms for attendance submission
- **Tamper-Proof**: SHA-256 hashing with linked blocks prevents data manipulation
- **Persistent Storage**: JSON-based storage with automatic blockchain persistence
- **QR Code Generation**: Generate QR codes linking to submission forms
- **Real-time Processing**: Immediate blockchain updates with instant verification

## Architecture

### Core Components

```
certification-blockchain/
├── Core Blockchain Logic
│   ├── block.go          # Block structure and validation
│   ├── blockchain.go     # Chain management and operations
│   └── storage.go        # Persistence layer
├── Web Interface
│   ├── form.html         # Attendance submission form
│   ├── lookup.html       # Certification lookup (planned)
│   └── styles.css        # Styling
├── Utilities
│   ├── qrcode.go         # QR code generation
│   └── main.go           # HTTP server and routing
└── Configuration
    ├── go.mod            # Go module dependencies
    └── blockchain.json   # Persistent blockchain data
```

### Data Flow

1. **User Submission** → Web form captures attendance data
2. **Block Creation** → New block with certificate data
3. **Hash Calculation** → SHA-256 hash of block contents
4. **Chain Validation** → Verify block integrity and linkage
5. **Persistent Storage** → Save to JSON file
6. **Verification** → Cryptographic proof of authenticity

## Getting Started

### Prerequisites

- Go 1.22.2 or higher
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd certification-blockchain
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Run the application**
   ```bash
   go run .
   ```

4. **Access the web interface**
   ```
   http://localhost:8080
   ```

### Building from Source

```bash
go build -o certchain .
./certchain
```

## Usage

### Submitting Attendance

1. Navigate to the web form at `http://localhost:8080/form`
2. Enter your name and event name
3. Submit the form
4. Your attendance is automatically recorded in the blockchain

### API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/` | GET | Redirects to attendance form |
| `/form` | GET | Serves the attendance submission form |
| `/submit` | POST | Processes attendance submissions |

### Blockchain Operations

- **Add Certification**: Automatically creates new blocks for each submission
- **Chain Validation**: Built-in integrity checking
- **Persistent Storage**: Automatic saving to `blockchain.json`

## Technical Details

### Block Structure

```go
type Block struct {
    Index     int         // Block position in chain
    Timestamp string      // Creation timestamp (RFC3339)
    Data      Certificate // Certificate information
    PrevHash  string      // Hash of previous block
    Hash      string      // Current block hash (SHA-256)
}
```

### Certificate Structure

```go
type Certificate struct {
    MemberID   string    // Unique member identifier
    Name       string    // Attendee name
    EventName  string    // Event name
    DateIssued time.Time // Issue timestamp
}
```

### Cryptographic Security

- **Hash Algorithm**: SHA-256
- **Block Linking**: Each block references the previous block's hash
- **Integrity Verification**: Automatic validation of chain integrity
- **Tamper Detection**: Any modification breaks the cryptographic chain

## Development

### Project Structure

- **`block.go`**: Core block operations and validation
- **`blockchain.go`**: Chain management and certification handling
- **`storage.go`**: JSON-based persistence layer
- **`qrcode.go`**: QR code generation utilities
- **`main.go`**: HTTP server and routing logic
- **`web/`**: Frontend HTML and CSS files

### Adding New Features

1. **Extend Certificate Structure**: Modify the `Certificate` struct in `block.go`
2. **Add New Endpoints**: Extend routing in `main.go`
3. **Enhance Validation**: Add custom validation logic in `block.go`
4. **Improve Storage**: Extend persistence methods in `storage.go`

### Testing

```bash
# Run tests (when implemented)
go test ./...

# Run with coverage
go test -cover ./...
```

## Roadmap

### Planned Features

- [ ] **Certification Lookup**: Search and verify existing certifications
- [ ] **Admin Panel**: Approval workflow for certifications
- [ ] **Batch Operations**: Bulk certification processing
- [ ] **API Authentication**: Secure access to blockchain operations
- [ ] **Mobile App**: Native mobile interface
- [ ] **Blockchain Explorer**: Visual chain inspection tool
- [ ] **Smart Contracts**: Advanced validation rules
- [ ] **Multi-chain Support**: Interoperability with other blockchains

### Current Status

- Core blockchain implementation
- Web interface for submissions
- Persistent storage
- Cryptographic verification
- QR code generation (basic implementation)
- Lookup functionality (planned)

## Contributing

We welcome contributions! Here's how you can help:

1. **Fork the repository**
2. **Create a feature branch**: `git checkout -b feature/amazing-feature`
3. **Commit your changes**: `git commit -m 'Add amazing feature'`
4. **Push to the branch**: `git push origin feature/amazing-feature`
5. **Open a Pull Request**

### Development Guidelines

- Follow Go coding standards
- Add tests for new functionality
- Update documentation for API changes
- Ensure blockchain integrity is maintained

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with [Go](https://golang.org/) programming language
- QR code generation powered by [go-qrcode](https://github.com/skip2/go-qrcode)
- SHA-256 cryptographic hashing for security
- Modern web standards for the user interface

## Support

- **Issues**: Report bugs and feature requests on GitHub
- **Discussions**: Join community discussions
- **Documentation**: Check this README and inline code comments

---

**Built with Go and blockchain technology**

*Secure • Immutable • Verifiable*
