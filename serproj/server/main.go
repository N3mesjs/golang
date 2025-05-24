package main

import (
    "bufio"
    "fmt"
    "net"
)

func main() {
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        panic(err)
    }
    fmt.Println("Server avviato sulla porta 8080")

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Errore connessione:", err)
            continue
        }

        go gestisciConnessione(conn)
    }
}

func gestisciConnessione(conn net.Conn) {
    defer conn.Close()
    indirizzo := conn.RemoteAddr().String()
    fmt.Println("Macchinario connesso:", indirizzo)

    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        messaggio := scanner.Text()
        fmt.Printf("Ricevuto da %s: %s\n", indirizzo, messaggio)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Errore durante la lettura:", err)
    }

    fmt.Println("Connessione chiusa da", indirizzo)
}
