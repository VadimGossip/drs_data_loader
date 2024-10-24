package model

const (
	RAObjectKey     string = "RA"
	RBObjectKey     string = "RB"
	RTSObjectKey    string = "RTS"
	RVObjectKey     string = "RV"
	CURRTSObjectKey string = "CURRTS"
)

type GwgrRmsgKey struct {
	GwgrId int64
	RmsgId int64
}

type ARmsgKey struct {
	GwgrId    int64
	Direction uint8
	BRmsgId   int64
	Code      uint64
}

type ARmsgShortKey struct {
	GwgrId  int64
	BRmsgId int64
	Code    uint64
}

type BRmsgKey struct {
	GwgrId    int64
	Direction uint8
	Code      uint64
}

type BRmsgShortKey struct {
	GwgrId int64
	Code   uint64
}

type IdHistItem struct {
	Id     int64
	DBegin int64
	DEnd   int64
}

type RateKey struct {
	GwgrId    int64
	Direction uint8
	ARmsgId   int64
	BRmsgId   int64
}

type RmsRateHistItem struct {
	RmsrId int64
	RmsvId int64
	DBegin int64
	DEnd   int64
}

type Rate struct {
	Price      float64
	CurrencyId int64
}

type CurrencyRateHist struct {
	CurrencyRate float64
	DBegin       int64
	DEnd         int64
}

type RateBase struct {
	RmsrId    int64
	PriceBase float64
}
