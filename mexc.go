package mexc

import (
	mexchttp "github.com/bogdankorobka/mexc-golang-sdk/http"
	mexchttpmarket "github.com/bogdankorobka/mexc-golang-sdk/http/market"
	mexcws "github.com/bogdankorobka/mexc-golang-sdk/websocket"
	mexcwsmarket "github.com/bogdankorobka/mexc-golang-sdk/websocket/market"
)

type Rest struct {
	MarketService *mexchttpmarket.Service
}

func NewRest(mexcHttp *mexchttp.Client) *Rest {
	return &Rest{
		MarketService: mexchttpmarket.New(mexcHttp),
	}
}

type Ws struct {
	MarketService *mexcwsmarket.Service
}

func NewWs(mexcWs *mexcws.MEXCWebSocket) *Ws {
	return &Ws{
		MarketService: mexcwsmarket.New(mexcWs),
	}
}
