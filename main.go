package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	for {
		conn, err := net.Dial("tcp", "localhost:3333")
		if err != nil {
			fmt.Errorf("Error connect to server", err)
		}
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Text to send: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Errorf("Error: ", err)
		}
		fmt.Fprint(conn, text+"\n")
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Errorf("Error :", err)
		}
		fmt.Println("Message from server: " + response)
	}

}
