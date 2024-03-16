package robocaractions

import (
	"context"
	"fmt"
	"strconv"

	directions "mqtt-robo-car-v1/robo-car-actions/directions"
	systemconfigreader "mqtt-robo-car-v1/systemconfigreader"

	autopaho "github.com/eclipse/paho.golang/autopaho"
	paho "github.com/eclipse/paho.golang/paho"
)

var config systemconfigreader.RoboCarConfig = systemconfigreader.GetRoboCarConfigParams()

/*
const (
	fowardMessage = "Forward"
	leftMessage   = "Left"
	rightMessage  = "Right"
	stopMessage   = "Stop"

	acknowledgementTopic   = "robo-car/Acknowledgement"
	acknowledgementMessage = "Acknowledgement"
)
*/

var ackMsgPaho = &paho.Publish{
	QoS:     1,
	Topic:   config.AcknowledgementTopic,
	Payload: []byte(config.AcknowledgementMessage),
}

var motA, motB = directions.Motors()

func SetDirection(pr autopaho.PublishReceived, mqttConnection *autopaho.ConnectionManager, ctx context.Context) {
	switch string(pr.Packet.Payload) {
	case config.FowardMessage:
		forward()
		mqttConnection.Publish(ctx, ackMsgPaho)
	case config.LeftMessage:
		left()
		mqttConnection.Publish(ctx, ackMsgPaho)
	case config.RightMessage:
		right()
		mqttConnection.Publish(ctx, ackMsgPaho)
	case config.StopMessage:
		stop()
		mqttConnection.Publish(ctx, ackMsgPaho)
	default:
		fmt.Println("an unknown robo-car direction message was given")
	}
}

func SetSpeed(pr autopaho.PublishReceived, mqttConnection *autopaho.ConnectionManager, ctx context.Context) {

	speed, err := strconv.Atoi(string(pr.Packet.Payload))

	if err != nil {
		fmt.Println("an unknown robo-car set speed message was given")
	}

	if speed < 1 {
		fmt.Println("incorrect speed range")
	} else if speed > 100 {
		fmt.Println("incorrect speed range")
	} else {
		setSpeed(speed)
	}
}

func forward() {
	fmt.Println("robo-car " + config.FowardMessage)
	directions.Foward(motA, motB)
}

func left() {
	fmt.Println("robo-car " + config.LeftMessage)
	directions.Left(motA, motB)
}

func right() {
	fmt.Println("robo-car " + config.RightMessage)
	directions.Right(motA, motB)
}

func stop() {
	fmt.Println("robo-car " + config.StopMessage)
	directions.Stop(motA, motB)
}

func setSpeed(speed int) {
	realSpeed := speed * 10000
	fmt.Println("robo-car setting speed to " + fmt.Sprint(speed) + "%")
	directions.SetSpeed(motA, motB, realSpeed)
}
