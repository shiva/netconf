package main

import (
    "fmt"
    "log"
    "net"
    "encoding/gob"
)

type P struct {
    M, N int64
}

func main() {
    fmt.Println("Start client")
    conn, err := net.Dial("tcp", "localhost:830")
    if err != nil {
        log.Fatal("Connection error", err)
    }

    encoder := gob.NewEncoder(conn)
    p := &P{1, 2}
    encoder.Encode(p)
    conn.Close()
    fmt.Println("done");
}

