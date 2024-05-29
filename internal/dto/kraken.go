package dto

type Pair struct {
	Ask [3]string `json:"a"`
	Bid [3]string `json:"b"`
	// LastTradeClosed [price, lot volume]
	LastTradeClosed            [2]string `json:"c"`
	Volume                     [2]string `json:"v"`
	VolumeWeightedAveragePrice [2]string `json:"p"`
	NumberOfTrades             [2]int    `json:"t"`
	Low                        [2]string `json:"l"`
	High                       [2]string `json:"h"`
	TodayOpenPrice             string    `json:"o"`
}

func (p Pair) GetLastTradeClosed() string {
	return p.LastTradeClosed[0]
}

type TickerResponse struct {
	Error  []string        `json:"error"`
	Result map[string]Pair `json:"result"`
}
