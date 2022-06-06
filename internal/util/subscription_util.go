package util

import (
	"ether-subscription-app/internal/config"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func GetWebsocketSubscription(mainConfig config.MainConfig, bloxrouteConfig config.BloxrouteConfig) *websocket.Conn {
	var websocketSubscriptionUrl string

	if bloxrouteConfig.EnterpriseSubscriptionEnabled {
		websocketSubscriptionUrl = bloxrouteConfig.EnterpriseSubscriptionUrl
		log.Println(fmt.Sprintf("Using Enterprise Subscription Plan URL of %s", websocketSubscriptionUrl))
	} else {
		websocketSubscriptionUrl = bloxrouteConfig.ProfessionalSubscriptionUrl
		log.Println(fmt.Sprintf("Using Professional Subscription Plan URL of %s", websocketSubscriptionUrl))
	}

	wsSubscriber := getWebsocket(websocketSubscriptionUrl, bloxrouteConfig.AuthorizationKey)
	subscriptionRequest := getSubscriptionRequest(mainConfig)

	err := wsSubscriber.WriteMessage(websocket.TextMessage, []byte(subscriptionRequest))
	if err != nil {
		panic(err)
	}

	return wsSubscriber
}

func getWebsocket(websocketSubscriptionUrl string, authorizationKey string) *websocket.Conn {
	dialer := websocket.DefaultDialer
	wsSubscriber, _, err := dialer.Dial(websocketSubscriptionUrl, http.Header{"Authorization": {authorizationKey}})
	if err != nil {
		panic(err)
	}
	return wsSubscriber
}

func getSubscriptionRequest(mainConfig config.MainConfig) string {

	toAddressFilter := mainConfig.SubscriptionFilters.ToAddress
	fromAddressFilter := mainConfig.SubscriptionFilters.FromAddress

	if len(toAddressFilter) == 0 && len(fromAddressFilter) == 0 {
		panic("Must include a to or from address filter")
	}

	var filtersStringQuery string
	if len(toAddressFilter) > 0 {
		filtersStringQuery = fmt.Sprintf("to = %s", toAddressFilter)
	}

	if len(fromAddressFilter) > 0 {
		fromString := fmt.Sprintf("from = %s", fromAddressFilter)
		if len(filtersStringQuery) > 0 {
			filtersStringQuery = fmt.Sprintf("%s and %s", filtersStringQuery, fromString)
		} else {
			filtersStringQuery = fromString
		}
	}

	subscriptionDetails := fmt.Sprintf(`{
													"jsonrpc": "2.0",
													"id": 1,
													"method": "subscribe",
													"params": [
														"newTxs",
														{
															"include": [
																"tx_hash",
																"tx_contents.from",
																"tx_contents.gas",
																"tx_contents.gas_price",
																"tx_contents.value",
																"tx_contents.input",
																"tx_contents.max_fee_per_gas",
																"tx_contents.max_priority_fee_per_gas",
																"tx_contents.nonce",
																"tx_contents.r",
																"tx_contents.s",
																"tx_contents.v",
																"tx_contents.to"
															],
														"filters": "%s"
														}
													]
												}`, filtersStringQuery)

	return subscriptionDetails
}
