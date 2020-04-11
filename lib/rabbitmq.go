package lib

import (
	"fmt"
	"os"

	"github.com/marceloagmelo/go-rabbitmq-receive/utils"
	"github.com/streadway/amqp"
)

const (
	fila string = "go-rabbitmq"
)

//ConectarRabbitMQ no rabbitmq
func ConectarRabbitMQ() (conn *amqp.Connection) {

	// Conectar com o rabbitmq
	var connectionString = fmt.Sprintf("amqp://%s:%s@%s:%s%s", os.Getenv("RABBITMQ_USER"), os.Getenv("RABBITMQ_PASS"), os.Getenv("RABBITMQ_HOSTNAME"), os.Getenv("RABBITMQ_PORT"), os.Getenv("RABBITMQ_VHOST"))
	conn, err := amqp.Dial(connectionString)
	utils.CheckErrFatal(err, "Falha ao conectar com o rabbitmq")

	return conn
}

//LerMensagensRabbitMQ no rabbitmq
func LerMensagensRabbitMQ(conn *amqp.Connection) (msgs <-chan amqp.Delivery) {
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
	msgs, err = ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.CheckErrFatal(err, "Falha ao ler as mensagens no rabbitmq")

	return msgs
	/*forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Processou a mensagem: %s", d.Body)
		}
	}()

	log.Printf(" [*] Esperando mensagens. Para sair pressione CTRL+C")
	<-forever*/
}
