package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

func Handler(conn net.Conn) {
	defer conn.Close()

	request, err := http.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		fmt.Println("Error reading request", err.Error())
		return
	}

	var path = request.URL.Path
	if path == "/" {

		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
		return
	}

	if strings.Contains(path, "/echo/") {
		message := strings.Split(path, "/")[2]
		conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(message), message)))
		return
	}

	conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
}

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	c, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	Handler(c)

}
