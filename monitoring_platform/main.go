package main

import (
	"fmt"

	mqqt "github.com/eclipse/paho.mqtt.golang"
)

func messageHandler(client mqqt.Client, msg mqqt.Message) {

	println(string(msg.Payload()))
}

func main() {
	fmt.Println("Water Monitoring System Started!")

	opts := mqqt.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883")
	println("Added broker on the host address and port")

	client := mqqt.NewClient(opts)

	token := client.Connect()
	token.Wait()
	println("Connected to the broker")

	client.Subscribe("test", 0, messageHandler)
	token.Wait()
	println("Subscriber created")
	println("Listening.....")

	select {}
}
