package mexcwsmarket

import (
	"encoding/json"
	"fmt"

	mexcws "github.com/bogdankorobka/mexc-golang-sdk/websocket"
)

type OrderBook struct {
	Channel string `json:"c"`
	Data    struct {
		Bids []struct {
			Price  string `json:"p"`
			Volume string `json:"v"`
		} `json:"bids"`
		Asks []struct {
			Price  string `json:"p"`
			Volume string `json:"v"`
		} `json:"asks"`
		Event     string `json:"e"`
		RequestID string `json:"r"`
	} `json:"d"`
	Symbol    string `json:"s"`
	Timestamp int64  `json:"t"`
}

func (s *Service) OrderBook(symbols []string, level string, listener func(*OrderBook)) error {
	lstnr := func(message string) {
		var book OrderBook

		err := json.Unmarshal([]byte(message), &book)
		if err != nil {
			fmt.Println("OrderBook listener unmarshal error:", err)
			return
		}

		listener(&book)
	}

	req := &mexcws.WsReq{
		Method: "SUBSCRIPTION",
		Params: []string{},
	}

	for _, symbol := range symbols {
		channel := fmt.Sprintf("spot@public.limit.depth.v3.api@%s@%s", symbol, level)

		req.Params = append(req.Params, channel)
		s.client.Subs.Add(channel, lstnr)
	}

	return s.client.Send(req)
}

func (s *Service) OrderBookV3(symbols []string, level string, listener func(*OrderBook)) error {
	fmt.Println("Connecting Orderbook... ")
	lstnr := func(message string) {
		var book OrderBook

		err := json.Unmarshal([]byte(message), &book)
		if err != nil {
			fmt.Println("OrderBook listener unmarshal error:", err)
			return
		}

		listener(&book)
	}

	req := &mexcws.WsReq{
		Method: "SUBSCRIPTION",
		Params: []string{},
	}

	for _, symbol := range symbols {
		channel := fmt.Sprintf("spot@public.aggre.depth.v3.api.pb@%sms@%s", level, symbol)
		fmt.Println(channel)

		req.Params = append(req.Params, channel)
		s.client.Subs.Add(channel, lstnr)
	}

	return s.client.Send(req)
}
