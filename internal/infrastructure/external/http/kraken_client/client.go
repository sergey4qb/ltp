package kraken_client

import (
	"context"
	"encoding/json"
	"github.com/sergey4qb/ltp/internal/dto"
	"net/http"
	"net/url"
)

type PublicClient struct {
	client http.Client
}

func NewPublicClient() *PublicClient {
	return &PublicClient{
		client: http.Client{},
	}
}

func (c *PublicClient) GetTickerInfo(ctx context.Context, pair string) (*dto.TickerResponse, error) {
	u, err := url.Parse(BaseUrl + PathTicker)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("pair", pair)
	u.RawQuery = params.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errUnexpectedCode(resp.StatusCode)
	}

	var ticker dto.TickerResponse
	if err := json.NewDecoder(resp.Body).Decode(&ticker); err != nil {
		return nil, err
	}
	if len(ticker.Error) > 0 {
		return nil, errThirdPartyError(ticker.Error[0])
	}
	return &ticker, nil
}
