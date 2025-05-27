package main

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "net"
    "time"
)

type StatoMacchinario string

const (
    StatoIdle    StatoMacchinario = "idle"
    StatoWorking StatoMacchinario = "working"
    StatoError   StatoMacchinario = "error"
)

type Prodotto struct {
    IDMacchinario int              `json:"id_macchinario"`
    Timestamp     string           `json:"timestamp"`
    Stato         StatoMacchinario `json:"stato"`
    ProdottoID    int              `json:"prodotto_id"`
    Messaggio     string           `json:"messaggio"`
}

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Errore di connessione:", err)
        return
    }
    defer conn.Close()
    fmt.Println("Macchinario connesso al server")

    rand.Seed(time.Now().UnixNano())
    idMacchinario := rand.Intn(1000)
    prodottoCounter := 1

    for {
        stato := StatoWorking
        if rand.Float32() < 0.05 {
            stato = StatoError
        }

        prodotto := Prodotto{
            IDMacchinario: idMacchinario,
            Timestamp:     time.Now().Format(time.RFC3339),
            Stato:         stato,
            ProdottoID:    prodottoCounter,
            Messaggio:     fmt.Sprintf("Prodotto #%d completato", prodottoCounter),
        }

        data, err := json.Marshal(prodotto)
        if err != nil {
            fmt.Println("Errore serializzazione:", err)
            continue
        }

        _, err = conn.Write(append(data, '\n'))
        if err != nil {
            fmt.Println("Errore invio:", err)
            return
        }

        fmt.Println("Inviato:", string(data))
        prodottoCounter++
        time.Sleep(5 * time.Second)
    }
}
