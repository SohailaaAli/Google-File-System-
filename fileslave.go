package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// Create a TCP server to listen for incoming connections
	listener, err := net.Listen("tcp", "0.0.0.0:5000")
	if err != nil {
		log.Fatal("Error starting TCP server:", err)
	}
	defer listener.Close()

	log.Println("Slave server started. Waiting for connections...")

	for {
		// Accept incoming connection
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		// Handle incoming file part
		go handleFilePart(conn)
	}
}

func handleFilePart(conn net.Conn) {
	defer conn.Close()

	// Create a buffer to hold received data
	buffer := make([]byte, 1024) // Adjust the buffer size as needed

	// Open a new file to write the received data
	file, err := os.Create("received_part.txt")
	if err != nil {
		log.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for {
		// Read from the connection into the buffer
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break // Connection closed by the client
			}
			log.Println("Error reading from connection:", err)
			return
		}

		// Write the received data from the buffer to the file
		_, err = file.Write(buffer[:n])
		if err != nil {
			log.Println("Error writing to file:", err)
			return
		}
	}

	log.Println("File part received and saved.")
}
