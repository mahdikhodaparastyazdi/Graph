package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func main() {
	// Start the server
	go startServer()

	// Wait for user input to exit
	fmt.Println("Press enter to exit.")
	fmt.Scanln()
}
func startServer() {
	// Listen for incoming connections
	listener, err := net.Listen("tcp", "localhost:8888") // Replace with the desired listening address and port
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()

	log.Println("Server started, listening on", listener.Addr())

	var (
		totalSize int64
		count     int
		//mu        sync.Mutex
	)
	var conn net.Conn
	var number = 1
	for number != 0 {
		// Accept incoming connection
		conn, err = listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		number = 0
		// Handle connection in a separate goroutine
		//go handleConnection(conn, &totalSize, &count, &mu)
	}
	var numm = 1
	for {
		buffer := make([]byte, 100000) // Create a buffer to store the received data
		n, err := conn.Read(buffer)
		if err != nil {
			//fmt.Println("Failed to read data:", err)
			return
		}

		payload := buffer[:n]

		fmt.Println("number:", numm, ": ", string(payload)[:10])
		numm++
		// Log the received payload
		//log.Printf("Received payload: %s", payload)
		//fmt.Println(string(payload))
		// Update the total size and count
		payloadSize := int64(len(payload))
		//mu.Lock()
		totalSize += payloadSize
		count++
		//mu.Unlock()

		//log.Printf("Total Size: %d, Count: %d", totalSize, count)
	}

}

func handleConnection(conn net.Conn, totalSize *int64, count *int, mu *sync.Mutex) {
	defer conn.Close()

	//log.Printf("New connection from %s", conn.RemoteAddr())

	// Read incoming data from the connection
	//reader := bufio.NewReader(conn)
	//payload, err := reader.ReadBytes('\n')
	//if err != nil {
	//	log.Printf("Failed to read data from connection: %v", err)
	//	return
	//}

	buffer := make([]byte, 4096) // Create a buffer to store the received data
	n, err := conn.Read(buffer)
	if err != nil {
		//fmt.Println("Failed to read data:", err)
		return
	}

	payload := buffer[:n]
	// Log the received payload
	//log.Printf("Received payload: %s", payload)
	//fmt.Println(string(payload))
	// Update the total size and count
	payloadSize := int64(len(payload))
	mu.Lock()
	*totalSize += payloadSize
	*count++
	mu.Unlock()

	log.Printf("Total Size: %d, Count: %d", *totalSize, *count)
}
