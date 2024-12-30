# Certwatch 🔍

Certwatch is a real-time Certificate Transparency (CT) log monitor that tracks and displays SSL/TLS certificate issuance across the internet. It provides live monitoring of newly issued certificates, helping security researchers and system administrators stay informed about certificate activities.


https://github.com/user-attachments/assets/25e23352-bac0-4fa2-8542-968150b049cb


## Features

- Real-time monitoring of Certificate Transparency logs
- Live display of newly issued certificates
- Shows certificate source and timestamp
- Color-coded output for better readability
- Regex pattern matching to filter domains
- Silent mode for pipeline integration
- Lightweight and efficient

## Prerequisites

- Go 1.21 or higher

## Installation

```bash
git clone https://github.com/storbeck/certwatch.git
cd certwatch
go mod download
```

## Usage

### Basic Monitoring

To start monitoring all certificate transparency logs:

```bash
go run main.go
```

### Filtering Domains

To filter domains using a regex pattern:

```bash
go run main.go -E "pattern"
```

Examples:
```bash
# Monitor only staging and test domains
go run main.go -E "test|staging|internal"

# Monitor specific TLDs
go run main.go -E "\.edu$|\.gov$"

# Monitor subdomains
go run main.go -E "^api\.|^dev\."
```

### Silent Mode

Use silent mode (-s) to output only the matching domains, perfect for piping to other tools:

```bash
# Output only matching domains
go run main.go -s -E "\.edu$"

# Pipe to other tools
go run main.go -s -E "\.edu$" 
```

### Output Format

In normal mode, the program will display certificates in the following format:
```
[TIME] ✓ DOMAIN (SOURCE)
```

Where:
- TIME: Timestamp when the certificate was seen
- DOMAIN: The primary domain name on the certificate
- SOURCE: The Certificate Transparency log source

In silent mode (-s), only the domain name is printed:
```
domain.com
```

## Dependencies

- [certstream-go](https://github.com/CaliDog/certstream-go) - Go client for the CertStream protocol
- [color](https://github.com/fatih/color) - Color output formatting

## License

This project is open source and available under the MIT License.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
