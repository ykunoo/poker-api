package poker

type HandResult struct {
	RequestId string `json:"requestId"`
	Hand      string `json:"hand"`
	Yaku      string `json:"yaku"`
	Strongest bool   `json:"strongest"`
}

type InternalResult struct {
	HandResult
	Rank int
}

type HandError struct {
	RequestId     string   `json:"requestId"`
	Hand          string   `json:"hand"`
	ErrorMessages []string `json:"errorMessages"`
}

type Card struct {
	Suit  string // スート (h,s,c,d)
	Value string // 値 (1-13)
}
