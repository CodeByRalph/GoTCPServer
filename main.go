package main

import (
	"fmt"
	"net"
	"os"
)

// Server Information
const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func CreateServer(){
	fmt.Println("Server is running...")
	// Listen function 
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	
	//Check if err
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		os.Exit(1)
	}
	// exits program when function is finished
	defer server.Close()

	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for client...")

	// For each client that connects check for error if no error print client connected
	for {
		// Accept listens for and returns the connections
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Connection Error: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("Client connected!")
		go processClient(connection)
	}

}

// Function processes each client and takes in a connection of type net.Conn
func processClient(connection net.Conn){
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error Reading: ", err.Error())
	}
	fmt.Println("Received: ", string(buffer[:mLen]))
	_, err = connection.Write([]byte("Thanks! Got your message:" + string(buffer[:mLen])))
	connection.Close()
}

func main(){
	CreateServer()
}

