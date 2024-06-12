package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Read nickname from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your nickname: ")
	nick, _ := reader.ReadString('\n')
	nick = strings.TrimSpace(nick)

	// Send nickname to server
	fmt.Fprintf(conn, "%s\n", nick)

	// Start goroutine to read messages from server
	go func() {
		for {
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println("Error reading message from server:", err)
				return
			}
			fmt.Print(message)
		}
	}()

	// Read messages from user and send to server
	for {
		fmt.Print("Enter recipient and message (e.g., 'John Hello'): ")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		if message == "quit" {
			break
		}

		fmt.Fprintf(conn, "%s\n", message)
	}
}
