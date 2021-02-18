package MQService

import (
    "fmt"
    "log"

    "github.com/streadway/amqp"
)

var MQURL string

//rabbitMQ结构体
type RabbitMQ struct {
    conn    *amqp.Connection
    channel *amqp.Channel
    //队列名称
    QueueName string
    //交换机名称
    Exchange string
    //bind Key 名称
    Key string
    //连接信息
    Mqurl string
}

func init() {

}

//创建结构体实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
    return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
}

//断开channel 和 connection
func (r *RabbitMQ) Destory() {
    r.channel.Close()
    r.conn.Close()
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
    if err != nil {
        log.Fatalf("%s:%s", message, err)
        panic(fmt.Sprintf("%s:%s", message, err))
    }
}
