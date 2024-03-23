package main

import (
    "context"
    "encoding/json"
    "fmt"
    "math/rand"
    "strconv"
    "time"

    "github.com/segmentio/kafka-go"
)

func main() {
    topic := "temperature"

    w := kafka.NewWriter(kafka.WriterConfig{
        Brokers: []string{"kafka:39092"},
        Topic:   topic,
    })
    defer w.Close()

    rand.Seed(time.Now().UnixNano())

    for {
        value1 := 20 + rand.Float64()*(35-20)
        value2 := rand.Intn(15) + 10
        value3 := rand.Float64() * 100

        msgMap := map[string]interface{}{
            "machine1": value1,
            "machine2": value2,
            "machine3": value3,
        }
        msgBytes, err := json.Marshal(msgMap)
        if err != nil {
            fmt.Println("Failed to marshal JSON:", err)
            return
        }

        msg := kafka.Message{
            Key:   []byte(strconv.FormatInt(time.Now().Unix(), 10)),
            Value: msgBytes,
        }

        err = w.WriteMessages(context.Background(), msg)
        if err != nil {
            fmt.Println("Failed to write messages:", err)
            return
        }

        fmt.Println("Published message:", string(msgBytes))
        time.Sleep(5 * time.Second)
    }
}