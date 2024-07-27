package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main() {
	// Open the file for reading
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read and divide the file into parts
	var parts [][]byte
	for scanner.Scan() {
		line := scanner.Bytes()
		part := append([]byte{}, line...) // Create a new slice for each part
		part = append(part, '\n')         // Append newline to preserve original lines
		parts = append(parts, part)
	}

	// Calculate the number of slaves and parts per slave
	numSlaves := len(slaveAddresses)
	partsPerSlave := len(parts) / numSlaves
	remainingParts := len(parts) % numSlaves

	// Connect to the slave machines and send parts of the file
	for i, slave := range slaveAddresses {
		conn, err := net.Dial("tcp", slave)
		if err != nil {
			log.Printf("Failed to connect to slave %s: %v", slave, err)
			continue
		}

		// Determine the range of parts to send to this slave
		start := i * partsPerSlave
		end := (i + 1) * partsPerSlave
		if i < remainingParts {
			end += 1 // Distribute remaining parts evenly
		}

		// Send parts of the file to the slave
		for j := start; j < end; j++ {
			_, err := conn.Write(parts[j])
			if err != nil {
				log.Printf("Failed to send data to slave %s: %v", slave, err)
				break
			}
		}
		conn.Close()
	}
}

var slaveAddresses = []string{"10.177.240.48:5000", "10.177.240.42:5000" /* "slave3_ip:port"*/}
