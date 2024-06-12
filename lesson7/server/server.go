package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	conn net.Conn
	nick string
}

var clients = make(map[string]*client)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is running on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Read nickname from client
	nick, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading nickname:", err)
		return
	}
	nick = strings.TrimSpace(nick)

	// Check if nickname is unique
	if _, ok := clients[nick]; ok {
		conn.Write([]byte("Nickname is already taken. Please choose another one.\n"))
		return
	}

	// Add client to the list
	clients[nick] = &client{conn: conn, nick: nick}
	fmt.Printf("New client connected: %s\n", nick)

	for {
		// Read message from client
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("Client disconnected: %s\n", nick)
			delete(clients, nick)
			return
		}
		message = strings.TrimSpace(message)

		// Parse recipient and message
		parts := strings.SplitN(message, " ", 2)
		if len(parts) != 2 {
			conn.Write([]byte("Invalid message format. Please use: <recipient> <message>\n"))
			continue
		}
		recipient, message := parts[0], parts[1]

		// Check if recipient exists
		if recipientClient, ok := clients[recipient]; ok {
			// Send message to recipient
			recipientClient.conn.Write([]byte(fmt.Sprintf("%s: %s\n", nick, message)))
		} else {
			conn.Write([]byte(fmt.Sprintf("Recipient '%s' not found.\n", recipient)))
		}
	}
}
