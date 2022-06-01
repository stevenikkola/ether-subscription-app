package util

import (
	"ether-subscription-app/internal/config"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

func GetWebsocketSubscription(mainConfig config.MainConfig, bloxrouteConfig config.BloxrouteConfig) *websocket.Conn {
	wsSubscriber := getWebsocket(bloxrouteConfig)
	subscriptionDetails := getSubscriptionDetails(mainConfig)

	err := wsSubscriber.WriteMessage(websocket.TextMessage, []byte(subscriptionDetails))
	if err != nil {
		panic(err)
	}

	return wsSubscriber
}

func getWebsocket(bloxrouteConfig config.BloxrouteConfig) *websocket.Conn {
	dialer := websocket.DefaultDialer
	wsSubscriber, _, err := dialer.Dial(bloxrouteConfig.WebsocketsCloudApiBaseUri, http.Header{"Authorization": {bloxrouteConfig.AuthorizationHeader}})
	if err != nil {
		panic(err)
	}
	return wsSubscriber
}

func getSubscriptionDetails(mainConfig config.MainConfig) string {

	toAddressFilter := mainConfig.SubscriptionFilters.ToAddress
	fromAddressFilter := mainConfig.SubscriptionFilters.FromAddress

	if len(toAddressFilter) == 0 && len(fromAddressFilter) == 0 {
		panic("Must include at least one address filter (to and/or from)")
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
