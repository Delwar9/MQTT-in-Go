package main

import (
	"fmt"

	"time"

	"github.com/delwar/mqtt/consts"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
    opts := mqtt.NewClientOptions()
    opts.AddBroker(consts.Broker)

    client := mqtt.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        panic(fmt.Sprintf("Error connecting to MQTT broker: %s", token.Error()))
    }

    for i := 1; i <= 20; i++ {
        message := fmt.Sprintf("Publishing message %d", i)
        token := client.Publish(consts.Topic, 0, false, message)
        if token.Wait() && token.Error() != nil {
            fmt.Printf("Error publishing message %d: %s\n", i, token.Error())
        } else {
            fmt.Println("Published:", message)
        }

        time.Sleep(1 * time.Second)
    }

    client.Disconnect(250)
}
