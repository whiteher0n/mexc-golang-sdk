package mexcwsmarket

import (
	"encoding/json"
	"fmt"

	mexcws "github.com/bogdankorobka/mexc-golang-sdk/websocket"
)

type Trade struct {
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

func (s *Service) Trade(symbol string, interval string, listener func(*Trade)) error {
	lstnr := func(message string) {
		var deal Trade

		fmt.Println("Message: ", message)

		err := json.Unmarshal([]byte(message), &deal)
		if err != nil {
			fmt.Println("Trade listener unmarshal error:", err)
			return
		}

		listener(&deal)
	}

	req := &mexcws.WsReq{
		Method: "SUBSCRIPTION",
		Params: []string{},
	}

	channel := fmt.Sprintf("spot@public.aggre.deals.v3.api.pb@%s@%s", interval, symbol)

	req.Params = append(req.Params, channel)
	s.client.Subs.Add(channel, lstnr)

	return s.client.Send(req)
}
