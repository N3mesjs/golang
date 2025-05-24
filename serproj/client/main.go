package main

import (
    "fmt"
    "math/rand"
    "net"
    "time"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        panic(err)
    }
    defer conn.Close()
    fmt.Println("Macchinario connesso al server")

    rand.Seed(time.Now().UnixNano())
    idMacchinario := rand.Intn(1000)
    i := 1

    for {
        prodotto := fmt.Sprintf("Macchinario %d: prodotto #%d completato", idMacchinario, i)
        _, err := fmt.Fprintln(conn, prodotto)
        if err != nil {
            fmt.Println("Errore invio messaggio:", err)
            break
        }

        fmt.Println("Inviato:", prodotto)
        time.Sleep(10 * time.Second) // Simula tempo produzione
        i++
    }
}
