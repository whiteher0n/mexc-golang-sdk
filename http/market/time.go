package mexchttpmarket

import (
	"context"
	"encoding/json"
	"fmt"
)

func (c *Service) Time(ctx context.Context) (*TimeResponse, error) {
	endpoint := "/api/v3/time"

	res, err := c.client.SendRequest(ctx, "GET", endpoint, nil)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil, err
	}

	var timeResponse TimeResponse
	err = json.Unmarshal(res, &timeResponse)
	if err != nil {
		return nil, err
	}

	return &timeResponse, nil
}

type TimeResponse struct {
	ServerTime int64 `json:"serverTime"`
}
