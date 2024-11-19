package main

import (
    "fmt"
    "net"
    "log"
)

func main() {
    // Listen on TCP port 8080
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatal(err)
    }
    defer listener.Close()
    fmt.Println("Server is listening on port 8080...")

    for {
        // Accept new connections
        conn, err := listener.Accept()
        if err != nil {
            log.Println(err)
            continue
        }

        // Handle each connection in a new goroutine
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    fmt.Println("New connection established:", conn.RemoteAddr())
    conn.Write([]byte("Hello from TCP server!\n"))
}

