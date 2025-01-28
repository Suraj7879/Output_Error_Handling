# Output Error Handling Project

A project to learn graceful error handling with a simulated unstable weather API server.

## Project Overview

This project simulates interaction with an unreliable weather API server that returns various types of errors. The goals are:
- Handle different error scenarios gracefully
- Provide clear, user-friendly error messages
- Demonstrate proper error recovery strategies

## Features

- Simulated server with multiple failure scenarios
- Client-side error handling for:
  - Server busy responses (with retry-after headers)
  - Malformed responses
  - Connection drops
  - Missing error headers
- Clear user feedback system
- Retry logic implementation

## Server Behavior Probabilities

| Scenario                      | Probability | Details |
|-------------------------------|-------------|---------|
| Successful response           | 50%         | Returns valid weather data |
| Server too busy               | 20%         | - 10%: Retry-After (seconds)<br>- 10%: Retry-After (timestamp) |
| Bugged response               | 10%         | Missing Retry-After header when required |
| Connection drop               | 20%         | No response received |

## Getting Started

### Prerequisites
- Go 1.20+ installed
- Basic understanding of HTTP protocols

### Installation
1. Clone the repository:
```bash
git clone https://github.com/yourusername/output_error_handling.git
cd output_error_handling
```

### Run the server
```bash
go run server/main.go
```