package main

import (
    "fmt"
    "net"
    "encoding/gob"
)

type P struct {
    M, N int64
}

func handleConnection(conn net.Conn) {
    dec := gob.NewDecoder(conn)
    p := &P{}
    dec.Decode(p)
    fmt.Println("Received : %+v", p);
}

func main() {
    fmt.Println("start")
    ln, err := net.Listen("tcp", ":830")
    if err != nil {
        fmt.Println("Couldn't listen on 830: %s", err)
    }

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("some error: %s", err)
            continue
        }
        go handleConnection(conn)
    }
}

