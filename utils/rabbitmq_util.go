package utils

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

var conn *amqp.Connection

var AmqpUrl = "amqp://user:pwd@host:port/vhost"

func init() {
	var err error
	conn, err = amqp.Dial(AmqpUrl)
	CheckErr(err)

	logrus.Warn("初始化连接：", AmqpUrl)
}

//发送消息到
func SendMsg(exchange, queue string, body []byte) {
	ch, err := conn.Channel()
	CheckErr(err)
	defer ch.Close()
	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	CheckErr(err)

	err = ch.Publish(
		exchange, // exchange
		q.Name,   // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	CheckErr(err)

	logrus.Infof("发送消息：%s", body)
}

//该方法会造成阻塞，协程调用
func ReceiveMsg(consumer, queue string, f func([]byte)) {
	ch, err := conn.Channel()
	CheckErr(err)
	defer ch.Close()
	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	CheckErr(err)
	msgs, err := ch.Consume(
		q.Name,   // queue
		consumer, // consumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	for d := range msgs {
		f(d.Body)
	}
}
