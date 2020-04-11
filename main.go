package main

import (
	"log"

	"github.com/marceloagmelo/go-rabbitmq-receive/lib"
	"github.com/marceloagmelo/go-rabbitmq-receive/models"
	"github.com/marceloagmelo/go-rabbitmq-receive/utils"
	"github.com/streadway/amqp"
)

const (
	fila string = "go-rabbitmq"
)

func main() {
	conn := lib.ConectarRabbitMQ()
	defer conn.Close()

	//lib.LerMensagensRabbitMQ(conn)

	processa(conn)

	/*msgs := lib.LerMensagensRabbitMQ(conn)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Processou a mensagem: %s", d.Body)
			strID := utils.BytesToString(d.Body)
			msg := models.AtualizarMensagem(strID)
			log.Println(msg)
		}
	}()

	log.Printf(" [*] Esperando mensagens. Para sair pressione CTRL+C")
	<-forever*/
}

func processa(conn *amqp.Connection) {
	// Abrir o canal
	ch, err := conn.Channel()
	utils.CheckErrFatal(err, "Falha ao abrir o canal no rabbitmq")
	defer ch.Close()

	// Declarara fila
	q, err := ch.QueueDeclare(
		fila,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	utils.CheckErrFatal(err, "Falha ao declarar a fila no rabbitmq")

	// Ler mensagens
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.CheckErrFatal(err, "Falha ao ler as mensagens no rabbitmq")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Processou a mensagem: %s", d.Body)
			strID := utils.BytesToString(d.Body)
			msg := models.AtualizarMensagem(strID)
			log.Println(msg)
		}
	}()

	log.Printf(" [*] Esperando mensagens da fila: %s", fila)
	<-forever

}
