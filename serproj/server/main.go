package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "net"
    "os"
    "sync"
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

var mu sync.Mutex

func main() {
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Errore apertura porta:", err)
        return
    }
    defer ln.Close()
    fmt.Println("Server avviato sulla porta 8080")

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Errore accettazione connessione:", err)
            continue
        }

        go gestisciConnessione(conn)
    }
}

func gestisciConnessione(conn net.Conn) {
    defer conn.Close()
    indirizzo := conn.RemoteAddr().String()
    fmt.Println("Nuova connessione da:", indirizzo)

    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        var prodotto Prodotto
        err := json.Unmarshal(scanner.Bytes(), &prodotto)
        if err != nil {
            fmt.Println("Errore parsing JSON da", indirizzo, ":", err)
            continue
        }

        log := fmt.Sprintf("[%s] Macchinario %d (%s): %s\n",
            prodotto.Timestamp, prodotto.IDMacchinario, prodotto.Stato, prodotto.Messaggio)

        fmt.Print(log)
        salvaLog(log)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Errore di lettura da", indirizzo, ":", err)
    }

    fmt.Println("Connessione terminata con", indirizzo)
}

func salvaLog(entry string) {
    mu.Lock()
    defer mu.Unlock()

    f, err := os.OpenFile("log_macchinari.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Errore salvataggio log:", err)
        return
    }
    defer f.Close()

    if _, err := f.WriteString(entry); err != nil {
        fmt.Println("Errore scrittura log:", err)
    }
}
