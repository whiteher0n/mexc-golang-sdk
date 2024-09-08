package mexc

import (
	"context"
	mexchttp "github.com/bogdankorobka/mexc-golang-sdk/http"
	mexchttpmarket "github.com/bogdankorobka/mexc-golang-sdk/http/market"
	mexcws "github.com/bogdankorobka/mexc-golang-sdk/websocket"
	mexcwsmarket "github.com/bogdankorobka/mexc-golang-sdk/websocket/market"
)

type Rest struct {
	MarketService *mexchttpmarket.Service
}

func NewRest(ctx context.Context, mexcHttp *mexchttp.Client) (*Rest, error) {
	marketService, err := mexchttpmarket.New(ctx, mexcHttp)
	if err != nil {
		return nil, err
	}

	return &Rest{
		MarketService: marketService,
	}, nil
}

type Ws struct {
	*mexcws.MEXCWebSocket
	MarketService *mexcwsmarket.Service
}

func NewWs(mexcWs *mexcws.MEXCWebSocket) *Ws {
	return &Ws{
		MEXCWebSocket: mexcWs,
		MarketService: mexcwsmarket.New(mexcWs),
	}
}
