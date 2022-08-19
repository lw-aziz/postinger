package mqtt

import (
	"postinger/util/logwrapper"

	"encoding/json"
	"fmt"
	"postinger/config"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Connection - Connection structure
type Connection struct {
	Conn mqtt.Client
}

// Client - MQTT Connection
var Client *Connection

// NewConnection - new connection of MQTT
func NewConnection(mqttConfig config.MqttConfig) error {
	if mqttConfig.URL == "" {
		return fmt.Errorf("CONFIGURATION IS MISSING FOR MQTT")
	}

	Client = &Connection{
		Conn: nil,
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(mqttConfig.URL)
	opts.SetUsername(mqttConfig.User)
	opts.SetPassword(mqttConfig.Password)
	opts.SetClientID("MQTT-DMS-" + strconv.Itoa(int(time.Now().UnixNano())))

	Client.Conn = mqtt.NewClient(opts)

	token := Client.Conn.Connect()
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}

	logwrapper.Logger.Infoln("Connected to MQTT_URL : ", mqttConfig.URL)

	return nil
}

// NotifyRealTimeRequest - for send to listner
type NotifyRealTimeRequest struct {
	Listner string
	Message interface{}
	Retain  bool
	Qos     int
}

// NotifyRealTime - send data in MQTT
func NotifyRealTime(request NotifyRealTimeRequest) {
	if request.Listner != "" {
		payload, _ := json.Marshal(request.Message)
		token := Client.Conn.Publish(request.Listner, byte(request.Qos), request.Retain, payload)
		if token.Wait() && token.Error() != nil {
			logwrapper.Logger.Errorln("MQTT Publish Error :", token.Error())
		}
	}
}

// NotifyRealTimeMultiRequest - for send to multiple listner
type NotifyRealTimeMultiRequest struct {
	Listner []string
	Message interface{}
	Retain  bool
	Qos     int
}

// NotifyRealTimeMulti - send data in MQTT to multiple listner
func NotifyRealTimeMulti(request NotifyRealTimeMultiRequest) {
	for _, listner := range request.Listner {
		notifyRealTime := NotifyRealTimeRequest{
			Listner: listner,
			Message: request.Message,
			Retain:  request.Retain,
			Qos:     request.Qos,
		}
		go NotifyRealTime(notifyRealTime)
	}
}
