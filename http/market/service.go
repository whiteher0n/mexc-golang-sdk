package mexchttpmarket

import mexchttp "github.com/bogdankorobka/mexc-golang-sdk/http"

type Service struct {
	client *mexchttp.Client
}

func New(client *mexchttp.Client) *Service {
	return &Service{client: client}
}
