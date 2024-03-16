package main

// GOOS=linux GOARCH=arm64 go build

import (
	"context"
	"fmt"
	clientconfig "mqtt-robo-car-v1/clientconfig"
	robocaractions "mqtt-robo-car-v1/robo-car-actions"
	"os"
	"os/signal"
	"syscall"
	"time"

	autopaho "github.com/eclipse/paho.golang/autopaho"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	mqttConnection, err := autopaho.NewConnection(ctx, clientconfig.GetClientConfig())

	if err != nil {
		panic(err)
	}

	if err = mqttConnection.AwaitConnection(ctx); err != nil {
		panic(err)
	}

	mqttConnection.AddOnPublishReceived(func(pr autopaho.PublishReceived) (bool, error) {
		b, err := controlDirectionMessages(pr, mqttConnection, ctx)
		return b, err
	})

	tickerHandler(ctx)

	fmt.Println("signal caught - exiting")

	<-mqttConnection.Done()
}

func controlDirectionMessages(pr autopaho.PublishReceived, mqttConnection *autopaho.ConnectionManager, ctx context.Context) (bool, error) {
	switch pr.Packet.Topic {
	case clientconfig.GetDirectionTopic():
		robocaractions.SetDirection(pr, mqttConnection, ctx)
	case clientconfig.GetSpeedTopic():
		robocaractions.SetSpeed(pr, mqttConnection, ctx)
	}
	return true, nil
}

func tickerHandler(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	msgCount := 0
	defer ticker.Stop()
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			msgCount++
			continue
		case <-ctx.Done():
		}
		break
	}
}
