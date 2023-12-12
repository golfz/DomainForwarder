# DomainForwarder

This Go-based reverse proxy service maps and forwards incoming HTTP requests from specified domains to target domains,
while maintaining the original request attributes such as query strings, paths, headers, and bodies.

## Features

- Maps full domains to other domains as specified in the JSON configuration file.
- Preserves the original properties of HTTP requests.
- Easy to configure and manage.

## Getting Started

### Prerequisites

- Go (version 1.20 or later)
- JSON file for domain mappings configuration.

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/golfz/DomainForwarder.git
   ```

2. Navigate to the project directory:

   ```bash
   cd DomainForwarder
   ```

### Configuration

Create a JSON file for domain mappings. The format should map source domains to target domains.

Example `config.json`:

```json
{
    "api.example1.com": "api.newdomain.com",
    "ing.example2.com": "img.newdomain.com"
}
```

example file: [config.json](config.json)

### Running the Service

Run the service using the Go command:

```bash
go run main.go
```

The service will start and listen on port 8080 by default.

### Usage
After starting the service, it will intercept incoming HTTP requests and forward them based on the domain mappings
defined in the `config.json` file.

## Building for Docker

```shell
env GOOS=linux GOARCH=amd64 go build -o DomainForwarder
```

### Contributing
Contributions are welcome! Please feel free to open issues or submit pull requests.

### License
This project is licensed under the MIT License.

### Contact
- [@golfz](https://twitter.com/golfz)

### Project Link: 
https://github.com/golfz/DomainForwarder