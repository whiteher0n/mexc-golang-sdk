package mexc

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	mexchttp "github.com/bogdankorobka/mexc-golang-sdk/http"
	mexcws "github.com/bogdankorobka/mexc-golang-sdk/websocket"
	mexcwsmarket "github.com/bogdankorobka/mexc-golang-sdk/websocket/market"
)

func TestHttp(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	cl := mexchttp.NewClient("", "", &http.Client{})

	rClient, _ := NewRest(ctx, cl)

	res, _ := rClient.MarketService.Ping(ctx)

	fmt.Println(res)
	cancel()
}

func TestOrderbookWs(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	wc := mexcws.NewMEXCWebSocket(func(err error) {
		fmt.Println("Error: ", err)
	})

	wc.Connect(ctx)

	ws := NewWs(wc)

	ws.MarketService.OrderBook(
		[]string{
			"BTCUSDT",
			"ETHUSDT",
		},
		"5",
		func(book *mexcwsmarket.OrderBook) {
			fmt.Println("Symbol: ", book.Symbol)
			fmt.Println("ASKS: ", book.Data.Asks)
			fmt.Println("BIDS: ", book.Data.Bids)
			fmt.Println("-----------")
		},
	)

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
	fmt.Println("END")
}

func TestOrderbookV2Ws(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	wc := mexcws.NewMEXCWebSocket(func(err error) {
		fmt.Println("Error: ", err)
	})

	wc.Connect(ctx)

	ws := NewWs(wc)

	ws.MarketService.OrderBookV2(
		[]string{
			"BTCUSDT",
			"ETHUSDT",
		},
		"5",
		func(book *mexcwsmarket.OrderBook) {
			fmt.Println("Symbol: ", book.Symbol)
			fmt.Println("ASKS: ", book.Data.Asks)
			fmt.Println("BIDS: ", book.Data.Bids)
			fmt.Println("-----------")
		},
	)

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
	fmt.Println("END")
}

func TestWsTrade(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	wc := mexcws.NewMEXCWebSocket(func(err error) {
		fmt.Println("Error: ", err)
	})

	wc.Connect(ctx)

	ws := NewWs(wc)

	ws.MarketService.Trade(
		"BTCUSDT",
		"10ms",
		func(deal *mexcwsmarket.Trade) {
			fmt.Println("Symbol: ", deal.Symbol)
			fmt.Println("Data: ", deal.Data)
			fmt.Println("-----------")
		},
	)

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
	fmt.Println("END")
}
