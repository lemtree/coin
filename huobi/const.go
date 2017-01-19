package huobi

const (
	HOST       = "https://api.huobi.com/"
	API_V3_URI = HOST + "apiv3/"
)

type periodT int
type coinT int
type currencyT int

// K线周期
const (
	PERIOD_1_MIN   periodT = 1
	PERIOD_5_MIN           = 5
	PERIOD_15_MIN          = 15
	PERIOD_30_MIN          = 30
	PERIOD_60_MIN          = 60
	PERIOD_1_DAY           = 100
	PERIOD_1_WEEK          = 200
	PERIOD_1_MONTH         = 300
	PERIOD_1_YEAR          = 400
)

// 数字货币
const (
	COIN_BTC coinT = 1 + iota
	COIN_LTC
)

//法币
const (
	CURRENCY_CNY currencyT = 1 + iota
	CURRENCY_USD
)

// 根据coin获得coin name，获取url时用
var coinNames map[coinT]string = map[coinT]string{COIN_BTC: "btc", COIN_LTC: "ltc"}

// 根据currency获得currency name，获取url时用
var marketNames map[currencyT]string = map[currencyT]string{CURRENCY_CNY: "staticmarket", CURRENCY_USD: "usdmarket"}

// 账号下的资产信息
type AccountInfo struct {
	Total               float64 `json:"total,string"`
	NetAsset            float64 `json:"net_asset,string"`
	AvailableCnyDisplay float64 `json:"available_cny_display,string"`
	AvailableBtcDisplay float64 `json:"available_btc_display,string"`
	AvailableLtcDisplay float64 `json:"available_ltc_display,string"`
	FrozenCnyDisplay    float64 `json:"frozen_cny_display,string"`
	FrozenBtcDisplay    float64 `json:"frozen_btc_display,string"`
	FrozenLtcDisplay    float64 `json:"frozen_ltc_display,string"`
	LoanCnyDisplay      float64 `json:"loan_cny_display,string"`
	LoanBtcDisplay      float64 `json:"loan_btc_display,string"`
	LoanLtcDisplay      float64 `json:"loan_ltc_display,string"`
}

// 批量返回的订单的单条信息
type Order struct {
	Id              int     `json:"id"`
	Type            int     `json:type`
	OrderPrice      float64 `json:"order_price,string"`
	OrderAmount     float64 `json:"order_amount,string"`
	ProcessedAmount float64 `json:"processed_amount,string"`
	OrderTime       int     `json:"order_time"`
}

// 返回某一条订单的详细信息
type OrderInfo struct {
	Id              int     `json:"id"`
	Type            int     `json:type`
	OrderPrice      float64 `json:"order_price,string"`
	OrderAmount     float64 `json:"order_amount,string"`
	ProcessedPrice  float64 `json:"processed_price,string"`
	ProcessedAmount float64 `json:"processed_amount,string"`
	Vot             float64 `json:"vot,string"`
	Fee             float64 `json:"fee,string"`
	Total           float64 `json:"total,string"`
	Status          int     `json:"status"`
}

// 已经成交的订单的信息
type OrderTraded struct {
	Order
	LastProcessedTime int `json:"last_processed_time"`
	Status            int `json:"status"`
}

// 通用返回信息
type Result struct {
	Result string `json:"result"`
	Id     int    `json:"id"`
}

// k线图
type Kline struct {
	DateTime int
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volue    float64
}

// 返回的实时行情
type RealTimeQuotation struct {
	DateTime int `json:"time,string"`
	Ticker   RealTimeQuotationTicker
}

// 返回的实时行情 ticker
type RealTimeQuotationTicker struct {
	High   float64
	Low    float64
	Symbol string
	Last   float64
	Volue  float64 `json:"vol"`
	Buy    float64
	Sell   float64
	Open   float64
}

type Depth struct {
	Asks   [][]float64
	Bids   [][]float64
	Symbol string
}

// 实时行情
type RealTimeTransactionData struct {
	Amount  float64
	Level   float64
	Buys    []RealTimeTransactionDataSubBuy
	Phigh   float64 `json:"p_high"`
	Plast   float64 `json:"p_last"`
	Plow    float64 `json:"p_low"`
	Pnew    float64 `json:"p_new"`
	Popen   float64 `json:"p_open"`
	Sells   []RealTimeTransactionDataSubBuy
	TopBuy  []RealTimeTransactionDataSubTopBuy `json:"top_buy"`
	TopSell []RealTimeTransactionDataSubTopBuy `json:"top_sell"`
	Total   float64
	Trades  []RealTimeTransactionDataSubTrade
	Symbol  string
}

//实时行情下的trade结构
type RealTimeTransactionDataSubTrade struct {
	Amount float64
	Price  float64
	Time   string
	EnType string `json:"en_type"`
	Type   string
}

//实时行情下的buy或sell结构
type RealTimeTransactionDataSubBuy struct {
	Amount float64
	Level  int
	Price  float64
}

//实时行情下的top_buy或top_sell结构
type RealTimeTransactionDataSubTopBuy struct {
	RealTimeTransactionDataSubBuy
	Accu float64
}
