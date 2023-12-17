# MQTT-in-Go
Real-time IoT Applications with MQTT and Golang, the popular MQTT library is paho.mqtt.golang
## What is MQTT?
MQTT is a lightweight messaging protocol based on the publish/subscribe model that is widely used in IoT and other applications where low bandwidth and low power consumption are important.

## What is Topic?
MQTT topics are a form of addressing that allows MQTT clients to share information. MQTT topics are structured in a hierarchy similar to a file system, with levels separated by forward slashes(/)

Sample : test/demo , room/temperature, room/humidity

## Access Details
Broker : The hostname or IP address of the MQTT broker
TCP Port : The port number used for MQTT communication over TCP
Topic: A form of addressing that allows MQTT clients to share information.

## Install paho.mqtt.golang
```bash
go get github.com/eclipse/paho.mqtt.golang
```
## Creating an MQTT Publisher in Golang
```go
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
```
## Creating an MQTT Subscriber in Golang

```go
package main

import (
	"fmt"

	"os"
	"os/signal"
	"syscall"

	"github.com/delwar/mqtt/consts"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", message.Payload(), message.Topic())
}

func main() {
	// mqtt topic
	topic := consts.Topic

	opts := mqtt.NewClientOptions()
	opts.AddBroker(consts.Broker)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("Error connecting to MQTT broker: %s", token.Error()))
	}

	if token := client.Subscribe(topic, 0, onMessageReceived); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("Error subscribing to topic: %s", token.Error()))
	}

	defer func() {
		client.Unsubscribe(topic)
		client.Disconnect(250)
	}()
	
	fmt.Println("Subscribed to topic:", topic)

	// Wait for a signal to exit the program gracefully
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	client.Unsubscribe(topic)
	client.Disconnect(250)
}

```
 
## Run the sample
Run Subscriber & Publisher in two different terminal
```bash
go run subscriber.go
go run publisher.go
```

Thanks for reading.
