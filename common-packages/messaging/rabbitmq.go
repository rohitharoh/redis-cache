package messaging

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"log"
	_ "time"
)

func GetRabbitMQConnection() (db *amqp.Connection, err error) {

	username := viper.GetString("messaging.rabbitmq.username")
	password := viper.GetString("messaging.rabbitmq.password")
	host := viper.GetString("messaging.rabbitmq.host")
	amqpPort := viper.GetString("messaging.rabbitmq.amqpPort")
	rabbitMQURL := "amqp://" + username + ":" + password + "@" + host + ":" + amqpPort + "/"
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to RabbitMQ", err)
	}
	return conn, err

}

func GetRabbitMqMQQTClient() (client MQTT.Client) {

	/*var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
		log.Printf("TOPIC: %s\n", msg.Topic())
		log.Printf("MSG: %s\n", msg.Payload())
	}*/

	username := viper.GetString("messaging.rabbitmq.username")
	password := viper.GetString("messaging.rabbitmq.password")
	host := viper.GetString("messaging.rabbitmq.host")
	mqqtPort := viper.GetString("messaging.rabbitmq.mqqtPort")
	rabbitMqMQQTURL := "tcp://" + host + ":" + mqqtPort
	log.Print(rabbitMqMQQTURL)
	opts := MQTT.NewClientOptions().AddBroker(rabbitMqMQQTURL)
	opts.SetClientID("go-simple")
	//opts.SetDefaultPublishHandler(f)
	opts.Username = username
	opts.Password = password

	cli := MQTT.NewClient(opts)

	if token := cli.Connect(); token.Wait() && token.Error() != nil {
		log.Print(token.Error())
	}

	//time.Sleep(3 * time.Second)

	if token := cli.Subscribe("go-mqtt/sample", 0, onMessageReceived); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}

	return cli

}

func onMessageReceived(client MQTT.Client, message MQTT.Message) {
	fmt.Printf("Received working message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
}
