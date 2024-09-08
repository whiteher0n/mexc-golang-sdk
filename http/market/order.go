package mexchttpmarket

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

func (s *Service) CreateOrder(ctx context.Context, req *CreateOrderRequest) (*OrderResponse, error) {
	endpoint := "/api/v3/order"

	params := make(map[string]string)

	params["symbol"] = req.Symbol
	params["side"] = string(req.Side)
	params["type"] = string(req.Type)
	params["timestamp"] = strconv.FormatInt(req.Timestamp, 10)

	if req.Quantity != nil {
		params["quantity"] = *req.Quantity
	}
	if req.QuoteOrderQty != nil {
		params["quoteOrderQty"] = *req.QuoteOrderQty
	}
	if req.Price != nil {
		params["price"] = *req.Price
	}
	if req.NewClientOrderId != nil {
		params["newClientOrderId"] = *req.NewClientOrderId
	}
	if req.RecvWindow != nil {
		params["recvWindow"] = fmt.Sprintf("%d", *req.RecvWindow)
	}

	res, err := s.client.SendRequest(ctx, "POST", endpoint, params)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil, err
	}

	var orderResponse OrderResponse
	err = json.Unmarshal(res, &orderResponse)
	if err != nil {
		return nil, err
	}

	return &orderResponse, nil
}

type CreateOrderRequest struct {
	Symbol           string  `json:"symbol"`
	Side             Side    `json:"side"`
	Type             Type    `json:"type"`
	Quantity         *string `json:"quantity,omitempty"`
	QuoteOrderQty    *string `json:"quoteOrderQty,omitempty"`
	Price            *string `json:"price,omitempty"`
	NewClientOrderId *string `json:"newClientOrderId,omitempty"`
	RecvWindow       *int64  `json:"recvWindow,omitempty"`
	Timestamp        int64   `json:"timestamp"`
}

type Side string
type Type string

const (
	SideBuy  Side = "BUY"
	SideSell Side = "SELL"
)

const (
	TypeLimit  Type = "LIMIT"
	TypeMarket Type = "MARKET"
)

type OrderResponse struct {
	Symbol       string `json:"symbol"`
	OrderId      string `json:"orderId"`
	OrderListId  int    `json:"orderListId"`
	Price        string `json:"price"`
	OrigQty      string `json:"origQty"`
	Type         string `json:"type"`
	Side         string `json:"side"`
	TransactTime int64  `json:"transactTime"`
}
