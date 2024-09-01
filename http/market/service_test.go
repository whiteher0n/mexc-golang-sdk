package mexchttpmarket

import (
	"context"
	"net/http"
	"testing"

	mexchttp "github.com/bogdankorobka/mexc-golang-sdk/http"

	"github.com/stretchr/testify/assert"
)

func TestService_Ping(t *testing.T) {
	client := mexchttp.NewClient("", "", &http.Client{})

	service := &Service{
		client: client,
	}

	ctx := context.Background()

	response, err := service.Ping(ctx)

	assert.NoError(t, err)

	expectedResponse := "{}"
	assert.Equal(t, expectedResponse, response)
}
