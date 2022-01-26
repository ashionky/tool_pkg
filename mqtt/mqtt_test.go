/**
 * @Author pibing
 * @create 2021/3/24 10:40 AM
 */

package mqtt

import (
	"encoding/json"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"testing"
	"time"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func TestMqtt(t *testing.T) {
	cfg := &Conf{
		Host:     "broker.emqx.io:1883",
		UserName: "emqx",
		Password: "public",
		ClientId: "test1--11",
	}
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", cfg.Host))
	opts.SetClientID(cfg.ClientId)
	opts.SetUsername(cfg.UserName)
	opts.SetPassword(cfg.Password)
	opts.SetDefaultPublishHandler(messagePubHandler) //可以提前注册好处理函数，也可以在订阅的时候callback
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	sub(client)
	publish(client)

	client.Disconnect(250)
}

func publish(client mqtt.Client) {
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message %d", i)
		token := client.Publish("topic/test", 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}

func sub(client mqtt.Client) {
	topic := "topic/test"
	token := client.Subscribe(topic, 1, nil)

	token.Wait()
	fmt.Printf("Subscribed to topic: %s \n", topic)
}

//===以上是原始的使用测试===


//以下是封装的方法测试
func TestNewClient(t *testing.T) {
	var (
		clientId = "test_clientid"
		//wg       sync.WaitGroup
	)

	cfg := &Conf{
		Host:     "broker.emqx.io:1883",
		UserName: "account_test",
		Password: "123456",
		ClientId: clientId,
	}
	client := NewClient(cfg)
	err := client.Connect()
	if err != nil {
		t.Errorf(err.Error())
	}

	//wg.Add(1)
	go func() {
		topics := []string{"mqtt", "mqtt2"}  //可以订阅多个topic
		err := client.Subscribe(func(c *Client, msg *Message) {
			fmt.Printf("接收到消息: %+v \n", msg)
			//wg.Done()
		}, 1, topics...)

		if err != nil {
			panic(err)
		}
	}()

	go func() {

		count := 0
		for true {
			count++
			if count > 30 {
				return
			}
			msg := &Message{
				ClientID: clientId,
				Type:     "",
				Data:     "Hello",
				Time:     int64(count),
			}
			topic := "mqtt"
			if count%2 == 0 {
				topic = "mqtt2"
			}
			msg.Type = topic
			data, _ := json.Marshal(msg)
			err = client.Publish(topic, 1, false, data)
			if err != nil {
				panic(err)
			}
		}

	}()

	time.Sleep(10 * time.Second)
	//wg.Wait()

}
