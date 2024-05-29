package price_service

import (
	"context"
	"github.com/sergey4qb/ltp/internal/core/domain/valueobject"
	"github.com/sergey4qb/ltp/internal/core/ports"

	"sync"
)

type Service struct {
	client ports.TickerClient
}

func NewService(client ports.TickerClient) *Service {
	return &Service{client: client}
}

func (s *Service) GetPricesForPairs(ctx context.Context, pairs []*valueobject.Pair) ([]valueobject.LTP, error) {
	var wg sync.WaitGroup
	results := make([]valueobject.LTP, 0, len(pairs))
	mu := sync.Mutex{}
	errChan := make(chan error, len(pairs))

	resultsChan := make(chan valueobject.LTP, len(pairs))

	for _, pair := range pairs {
		wg.Add(1)
		go func(pair *valueobject.Pair) {
			defer wg.Done()
			price, err := s.getLastTradedPrice(ctx, pair)
			if err != nil {
				errChan <- err
				return
			}
			resultsChan <- valueobject.LTP{
				Pair:   pair.ToStringDivider(),
				Amount: price,
			}
		}(pair)
	}
	wg.Wait()
	close(resultsChan)
	close(errChan)
	for result := range resultsChan {
		mu.Lock()
		results = append(results, result)
		mu.Unlock()
	}
	if len(errChan) > 0 {
		return nil, <-errChan
	}

	return results, nil
}

func (s *Service) getLastTradedPrice(ctx context.Context, pair *valueobject.Pair) (string, error) {
	ticker, err := s.client.GetTickerInfo(ctx, pair.ToString())
	if err != nil {
		return "", err
	}
	krakenStandard, err := pair.ToStringKraken()
	if err != nil {
		return "", err
	}
	return ticker.Result[krakenStandard].GetLastTradeClosed(), nil
}
