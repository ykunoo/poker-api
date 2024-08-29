package poker

import "fmt"

func EvaluateHands(hands []string) ([]HandResult, []HandError) {
	var internalResults []InternalResult
	var errors []HandError
	bestRank := 0

	for i, hand := range hands {

		requestId := fmt.Sprintf("01-000002-%02d", i+1)

		cards, err := ParseHand(hand)
		if err != nil {
			errors = append(errors, HandError{
				RequestId:     requestId,
				Hand:          hand,
				ErrorMessages: []string{err.Error()},
			})
			continue
		}

		yaku, rank := EvaluateHand(cards)
		if rank > bestRank {
			bestRank = rank
		}

		internalResults = append(internalResults, InternalResult{
			HandResult: HandResult{
				RequestId: requestId,
				Hand:      hand,
				Yaku:      yaku,
				Strongest: false,
			},
			Rank: rank,
		})
	}

	// 最強のランクを持つ手札に対してStrongestフラグを設定
	for i := range internalResults {
		if internalResults[i].Rank == bestRank {
			internalResults[i].Strongest = true
		}
	}

	// Rankフィールドを削除した最終的な結果を作成
	var results []HandResult
	for _, result := range internalResults {
		results = append(results, HandResult{
			RequestId: result.RequestId,
			Hand:      result.Hand,
			Yaku:      result.Yaku,
			Strongest: result.Strongest,
		})
	}

	return results, errors
}
