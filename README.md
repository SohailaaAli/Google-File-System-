# Google File System (GFS) Implementation in Go

## Overview

This project implements a simplified version of the Google File System (GFS) in the Go programming language. GFS is a scalable distributed file system developed by Google to handle large data processing workloads across many machines. This implementation is designed for educational purposes, demonstrating the key principles and architecture of GFS.

## Features

- **Scalable Architecture**: Supports distributed storage across multiple nodes.
- **Fault Tolerance**: Designed to handle node failures gracefully.
- **High Throughput**: Optimized for large data processing workloads.
- **Metadata Management**: Efficient management of file metadata and chunk locations.

## Requirements

- Go 1.15+ (https://golang.org/)
- A Linux or macOS environment (recommended for running multiple nodes locally)

## Installation

1. **Clone the repository:**
    ```sh
    git clone https://github.com/your-username/gfs-go.git
    cd gfs-go
    ```

2. **Install dependencies:**
    ```sh
    go mod tidy
    ```

## Usage

### Running the GFS Master Server

1. **Start the master server:**
    ```sh
    go run cmd/master/main.go
    ```

### Running GFS Chunk Servers

1. **Start chunk servers:**
    ```sh
    go run cmd/chunkserver/main.go -id=1
    go run cmd/chunkserver/main.go -id=2
    go run cmd/chunkserver/main.go -id=3
    ```
    Repeat for as many chunk servers as needed, each with a unique `-id`.

### Running the GFS Client

1. **Start the client:**
    ```sh
    go run cmd/client/main.go
    ```

2. **Interact with the file system:**
    Use the client to perform file operations such as creating, reading, and writing files.

## Project Details

For a detailed explanation of the project, including design decisions and implementation details, please refer to the [project details ](./Project_Details.txt).

## Contributing

Contributions are welcome! Please open an issue or submit a pull request with your improvements.

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.

## Contact

For any questions or suggestions, please contact [sohailaali555gmail.com].

