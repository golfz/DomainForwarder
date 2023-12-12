# DomainForwarder

This Go-based reverse proxy service maps and forwards incoming HTTP requests from specified domains to target domains, while maintaining the original request attributes such as query strings, paths, headers, and bodies.

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
   git clone https://github.com/yourusername/reverse-proxy-service.git

### Configuration

The configuration file is a JSON file that contains the mappings of domains to other domains. The file is structured as follows:

Example `config.json`:

```json
{
  "api.example.com": "api.newdomain.com",
  "img.example.com": "img.newdomain.com"
}
```