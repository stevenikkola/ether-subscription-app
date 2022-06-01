package main

import (
	"encoding/json"
	"ether-subscription-app/internal/config"
	"ether-subscription-app/internal/message"
	"ether-subscription-app/internal/util"
	"log"
)

func main() {
	// format our logger as date | time (resolved to microseconds)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	// load all the configs we'll need
	bloXrouteConfig := config.GetBloxrouteConfig()
	mainConfig := config.GetMainConfig()

	// log the configs out
	configPP, _ := json.MarshalIndent(mainConfig, "", "\t")
	log.Println(string(configPP))

	// create our websocket subscription
	wsSubscriber := util.GetWebsocketSubscription(mainConfig, bloXrouteConfig)

	log.Println("Waiting for messages")

	for {
		_, nextNotification, err := wsSubscriber.ReadMessage()
		if err != nil {
			log.Println(err)
		}

		// for each websocket subscription message, convert the response into a transaction object
		txContents := message.ParseMessage(nextNotification)

		// if we were unable to create a transaction object from the message, we cannot move forward
		// this is usually only the case where we get a response when first setting up our websocket connection
		if len(txContents.Input) == 0 {
			continue
		}

		// create a pretty printed version of our txContents for logging
		txPP, _ := json.MarshalIndent(txContents, "", "\t")
		log.Println(string(txPP))
	}
}
