# kpi-lab1

## Overview

This is a simple HTTP server written in Go that provides a single `/time` endpoint. When accessed, the server responds with the current server time in [RFC3339] format as a JSON object.

## Prerequisites

* Go 1.18 or later installed on your system
* `GOPATH` and `GOROOT` properly configured (optional if using Go modules)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Ruslan-Vladyslav/kpi-lab1
   cd kpi-lab1
   ```

2. Download dependencies:

   ```bash
   go mod tidy
   ```

## Usage

1. Build the server:

   ```bash
   go build -o time-server ./main.go
   ```

2. Run the server:

   ```bash
   ./time-server
   ```

   You should see the following log message:

   ```
   2025/05/20 10:00:00 Server is running on http://localhost:8795
   ```

3. Access the `/time` endpoint in your browser or via `curl`:

   ```bash
   curl http://localhost:8795/time
   ```

   Expected JSON response:

   ```json
   {
     "time": "2025-05-20T10:00:00+03:00"
   }
   ```

## API Endpoint

### GET `/time`

Retrieves the current server time.

**Response**

* Status: `200 OK`
* Content-Type: `application/json`
* Body:

  ```json
  {
    "time": "<current server time in RFC3339 format>"
  }
  ```

## Logging

The server logs each request with the timestamp and remote client address:

```
2025/05/20 10:00:00 GET /time -> 2025-05-20T10:00:00+03:00
2025/05/20 10:00:00 Request from 127.0.0.1:54321 to /time
```

## Configuration

* The server listens on port **8795** by default. To change the port, modify the `ListenAndServe` call in `main.go`.
