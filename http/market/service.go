package mexchttpmarket

import (
	mexchttp "github.com/bogdankorobka/mexc-golang-sdk/http"
	"strconv"
	"time"
)

type Service struct {
	client *mexchttp.Client
}

func New(client *mexchttp.Client) *Service {
	return &Service{client: client}
}

func (s *Service) getTimestamp() string {
	return strconv.FormatInt(time.Now().UnixMilli()-s.client.SyncTimeDeltaMilliSeconds, 10)
}
