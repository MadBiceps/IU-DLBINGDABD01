package main

import (
    "context"
    "fmt"
    "math/rand"
    "strconv"
    "time"

    "github.com/segmentio/kafka-go"
)

func main() {
    // Kafka-Topic konfigurieren
    topic := "temperature"

    // Erstelle einen Kafka-Writer
    w := kafka.NewWriter(kafka.WriterConfig{
        Brokers: []string{"kafka:9092"},
        Topic:   topic,
    })
    defer w.Close()

    rand.Seed(time.Now().UnixNano())

    for {
        // Generiere eine zufällige Temperatur
        temp := 20 + rand.Float64()*(500-100) // Zufällige Temperatur zwischen 20°C und 35°C
        tempStr := fmt.Sprintf("%.2f", temp)

        // Erstelle eine Nachricht
        msg := kafka.Message{
            Key:   []byte(strconv.FormatInt(time.Now().Unix(), 10)),
            Value: []byte(tempStr),
        }

        // Veröffentliche die Nachricht auf Kafka
        err := w.WriteMessages(context.Background(), msg)
        if err != nil {
            fmt.Println("Failed to write messages:", err)
            return
        }

        fmt.Println("Published temperature:", tempStr)

        // Warte für die Simulation
        time.Sleep(5 * time.Second)
    }
}