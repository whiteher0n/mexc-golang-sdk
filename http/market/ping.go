package mexchttpmarket

import (
	"context"
	"fmt"
)

func (c *Service) Ping(ctx context.Context) (string, error) {
	endpoint := "/api/v3/ping"

	res, err := c.client.SendRequest(ctx, "GET", endpoint, nil)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return "", err
	}

	return string(res), nil
}
