package clientconfig

import (
	"context"
	"fmt"
	systemconfigreader "mqtt-robo-car-v1/systemconfigreader"
	"net/url"

	autopaho "github.com/eclipse/paho.golang/autopaho"
	paho "github.com/eclipse/paho.golang/paho"
)

var config systemconfigreader.ClientConfig = systemconfigreader.GetClientConfigParams()

/*
const (
	ip             = "192.168.50.134"
	port           = "61616"
	clientID       = "robo-car"
	DirectionTopic = "robo-car/Direction"
	SpeedTopic     = "robo-car/Speed"
	Username       = "robo-car"
	Password       = "raspberrypi4_hack"
)
*/

func GetDirectionTopic() string {
	return config.DirectionTopic
}

func GetSpeedTopic() string {
	return config.SpeedTopic
}

func GetClientConfig() autopaho.ClientConfig {

	clientconfig := autopaho.ClientConfig{
		ServerUrls:                    []*url.URL{mqttUrl()},
		KeepAlive:                     20,
		CleanStartOnInitialConnection: false,
		SessionExpiryInterval:         60,
		OnConnectionUp:                func(cm *autopaho.ConnectionManager, connAck *paho.Connack) { onConnectionUp(cm, connAck) },
		OnConnectError:                func(err error) { onConnectError(err) },
		ClientConfig:                  pahoBaseClientConfig(),
		ConnectUsername:               config.Username,
		ConnectPassword:               []byte(config.Password),
	}
	if config.AnonymousConnection {
		clientconfig.ResetUsernamePassword()
	}
	return clientconfig
}

func pahoBaseClientConfig() paho.ClientConfig {
	var clientConfig = paho.ClientConfig{
		ClientID:           config.ClientID,
		OnClientError:      func(err error) { onClientError(err) },
		OnServerDisconnect: func(d *paho.Disconnect) { onServerDisconnect(d) },
	}
	return clientConfig
}

func mqttUrl() *url.URL {
	var mqttbrokerurl string = "mqtt://" + config.Ip + ":" + config.Port
	url, err := url.Parse(mqttbrokerurl)
	if err != nil {
		panic(err)
	}
	return url
}

var mqttSubscriptions = &paho.Subscribe{
	Subscriptions: []paho.SubscribeOptions{
		{Topic: config.DirectionTopic, QoS: 1},
		{Topic: config.SpeedTopic, QoS: 1},
	},
}

func onConnectionUp(cm *autopaho.ConnectionManager, connAck *paho.Connack) {
	fmt.Println("mqtt connection up")
	if _, err := cm.Subscribe(context.Background(), mqttSubscriptions); err != nil {
		fmt.Printf("failed to subscribe (%s). This is likely to mean no messages will be received.", err)
	}
	fmt.Println("mqtt subscription made")
}

func onConnectError(err error) {
	fmt.Printf("error whilst attempting connection: %s\n", err)
}

func onClientError(err error) {
	fmt.Printf("client error: %s\n", err)
}

func onServerDisconnect(d *paho.Disconnect) {
	if d.Properties != nil {
		fmt.Printf("server requested disconnect: %s\n", d.Properties.ReasonString)
	} else {
		fmt.Printf("server requested disconnect; reason code: %d\n", d.ReasonCode)
	}
}
