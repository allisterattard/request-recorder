# request-recorder 🚀🛡️

![Go](https://img.shields.io/badge/Go-1.27%2B-00ADD8?logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green.svg)

`request-recorder` is a lightweight, high-performance HTTP request logging and callback server written in Go. Designed with a modular architecture and an asynchronous producer-consumer concurrency model, it captures, inspects, and persists incoming HTTP traffic (headers, cookies, query parameters, and raw bodies) to disk without blocking network threads or dropping connections under heavy traffic loads.

Developed as a reliable utility for web application security assessments, lab environments, and tracking out-of-band interactions (such as blind SSRF or webhooks).

---

## Project Architecture

The project is structured following clean Go modular design patterns, separating data models, request handling logic, and application entry points:

```text
request-recorder/
├── api/
│   └── api.go         # HTTP request handler, parsing, and queue injection
├── model/
│   └── model.go       # Data structures (Logs struct) and shared buffered channel queue
├── go.mod             # Go module definition
├── main.go            # Entry point, background consumer worker (init), and HTTP server
└── hits.log           # Persistent JSON log output (generated at runtime)
```

## Key Features

- **Clean Modular Design**: Separates core data structures (`model`), HTTP request parsing (`api`), and application lifecycle (`main`) to prevent circular dependencies and maintain clean code readability.
- **Asynchronous Producer-Consumer Architecture**: Incoming requests are parsed, acknowledged instantly, and dropped into a buffered channel (`LogQueue`), isolating the network layer from slow disk I/O.
- **Non-Blocking Overflow Handling**: Implements a Go `select` statement with a `default` fallback. If the 10,000-item queue buffer fills up under extreme traffic spikes (e.g., heavy fuzzing), the server safely writes to standard output rather than freezing or dropping connections.
- **Thread-Safe Sequential Logging**: A dedicated background worker routine drains the queue sequentially, writing structured JSON lines without mutex lock contention or race conditions.
- **Zero External Dependencies**: Compiles down to a single, portable static binary that can be deployed instantly onto a target environment or VPS.

## Security Use Cases

- **Blind SSRF & Callback Logging**: Deploy on an external listener VPS during web assessments to catch out-of-band interactions, webhooks, or SSRF payloads.
- **Payload & Fuzzing Inspection**: Capture raw incoming bodies, malformed headers, and multi-value cookie strings sent by automated scanners or manual testing tools.
- **CPTS Lab & Assessment Utility**: Developed and utilized during HTB Certified Penetration Testing Specialist (CPTS) coursework, lab workflows, and offensive security practice to inspect local HTTP interactions and web app behavior.

## Getting Started

### Prerequisites
- Go 1.27+

### Installation & Compilation

1. Clone the repository and initialize the module:

``` bash
git clone https://github.com/allisterattard/request-recorder.git
cd request-recorder
go mod init request-recorder
```
2. Build the binary:

``` bash
go build -o request-recorder main.go
```

### Running the Server

Start the application (listens on port 8888 by default):

``` bash
./request-recorder
```

## Example Usage & Testing

Send a test request with custom headers and multiple cookies using `curl`:

``` bash
curl -X POST http://127.0.0.1:8888/api/v1/test \
  --cookie "session=xyz123; user=admin" \
  -H "X-Attack-Vector: blind-ssrf" \
  -d '{"target": "internal-service", "port": 8080}'
```

## Sample Log Output (`hits.log`)

``` json
{
  "timestamp": "2026-09-30 14:30:00",
  "method": "POST",
  "proto": "HTTP/1.1",
  "path": "/api/v1/test",
  "query": "",
  "host": "127.0.0.1:8888",
  "remoteAddress": "127.0.0.1:54321",
  "headers": {
    "Accept": ["*/*"],
    "Content-Length": ["43"],
    "Content-Type": ["application/x-www-form-urlencoded"],
    "Cookie": ["session=xyz123; user=admin"],
    "User-Agent": ["curl/8.7.1"],
    "X-Attack-Vector": ["blind-ssrf"]
  },
  "contentLength": 43,
  "cookies": "session=xyz123; user=admin",
  "body": "{\"target\": \"internal-service\", \"port\": 8080}"
}
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.