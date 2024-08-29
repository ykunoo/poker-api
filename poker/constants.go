package poker

const (
	// エラーメッセージ
	ErrInvalidCardFormat = "カードのフォーマットが不正です"
	ErrInvalidSuit       = "不正なスートが含まれています"
	ErrInvalidValue      = "不正なカードの値が含まれています"
	ErrInvalidHandSize   = "手札は5枚入力してください"

	// カードの値の範囲
	MinCardValue = 1
	MaxCardValue = 13

	// スート
	Heart   = "h"
	Spade   = "s"
	Club    = "c"
	Diamond = "d"

	// 役のランク
	HighCardRank           = 1
	OnePairRank            = 2
	TwoPairRank            = 3
	ThreeOfAKindRank       = 4
	StraightRank           = 5
	FlushRank              = 6
	FullHouseRank          = 7
	FourOfAKindRank        = 8
	StraightFlushRank      = 9
	RoyalStraightFlushRank = 10

	// 役の名前
	YakuHighCard           = "ハイカード"
	YakuOnePair            = "ワンペア"
	YakuTwoPair            = "ツーペア"
	YakuThreeOfAKind       = "スリーカード"
	YakuStraight           = "ストレート"
	YakuFlush              = "フラッシュ"
	YakuFullHouse          = "フルハウス"
	YakuFourOfAKind        = "フォーカード"
	YakuStraightFlush      = "ストレートフラッシュ"
	YakuRoyalStraightFlush = "ロイヤルストレートフラッシュ"
)
