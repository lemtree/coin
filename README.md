# coin 数字货币交易平台API
digiccy trading platform API collection, such as Bitcoin, Litecoin. current complete huobi api v3. 

数字货币交易平台API，目前包含了火币网的API V3的大部分接口，包含交易和行情接口，剩余的接口会陆续完善。其他平台的API接口后续补充。
##火币网API V3 

## 安装
```go
go get github.com/lemtree/coin/huobi
```


## 实例

```go
package main

import (
	"fmt"
	"github.com/lemtree/coin/huobi"
)

func main() {
	c := huobi.NewHuobiClient()

	//获取BTC-CNY行情，返回解析好的struct	
	QuoData, err := c.Quotation(huobi.COIN_BTC, huobi.CURRENCY_CNY)
	checkError(err)
	fmt.Println(QuoData)

	//获取BTC-CNY行情，返回的火币api的原始JSON
	jsonBlob, err := c.QuotationJson(huobi.COIN_BTC, huobi.CURRENCY_CNY)
	checkError(err)
	fmt.Println(string(jsonBlob))

	// 获取 BTC-CNY 1分钟K线数据，返回100条
	KlineData1mins, err := c.KlineBtcCny(huobi.PERIOD_1_MIN, 100)
	checkError(err)
	fmt.Println(KlineData1mins)

	// 获取 BTC-CNY 小时K线数据，返回200条
	KlineData1Hour, err := c.KlineBtcCny(huobi.PERIOD_60_MIN, 200)
	checkError(err)
	fmt.Println(KlineData1Hour)

	// 获取我的账户信息。
	// 需要在火币网获取访问秘钥,然后设置你的key
	huobi.SetAccessKey("access-key")
	huobi.SetSecretKey("secret-key")
	AccountInfo, err := c.GetAccountInfo()
	checkError(err)
	fmt.Println(AccountInfo)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
```

## 已经完成的API列表
### 行情相关API
```go
func (hb *HuobiClient) Kline(coinType coinT, currencyType currencyT, period periodT, length int) ([]Kline, error){}
func (hb *HuobiClient) KlineJson(coinType coinT, currencyType currencyT, period periodT, length int) ([]byte, error){}
func (hb *HuobiClient) KlineBtcCny(period periodT, length int) ([]Kline, error){}
func (hb *HuobiClient) KlineBtcUsd(period periodT, length int) ([]Kline, error){}
func (hb *HuobiClient) KlineLtcCny(period periodT, length int) ([]Kline, error){}
func (hb *HuobiClient) Quotation(coinType coinT, currencyType currencyT) (*RealTimeQuotation, error){}
func (hb *HuobiClient) QuotationJson(coinType coinT, currencyType currencyT) ([]byte, error){}
func (hb *HuobiClient) Depth(coinType coinT, currencyType currencyT, length int) (*Depth, error){} 
func (hb *HuobiClient) DepthJson(coinType coinT, currencyType currencyT, length int) ([]byte, error){}
func (hb *HuobiClient) RealTimeTransaction(coinType coinT, currencyType currencyT) (*RealTimeTransactionData, error){}
func (hb *HuobiClient) RealTimeTransactionJson(coinType coinT, currencyType currencyT) ([]byte, error){}
```

### 交易相关API
	注：交易相关API需要取得huobi的交易秘钥才能访问
```go
func (hb *HuobiClient) GetAccountInfo() (*AccountInfo, error){}
func (hb *HuobiClient) GetAccountInfoJson() ([]byte, error){}
func (hb *HuobiClient) GetOrders(coinType int) ([]Order, error){}
func (hb *HuobiClient) GetOrdersJson(coinType int) ([]byte, error){}
func (hb *HuobiClient) OrderInfo(coinType, id int){}
func (hb *HuobiClient) OrderInfoJson(coinType, id int) ([]byte, error){}
func (hb *HuobiClient) Buy(coinType int, price, amount float64) (*Result, error){}
func (hb *HuobiClient) BuyJson(coinType int, price, amount float64) ([]byte, error){}
func (hb *HuobiClient) Sell(coinType int, price, amount float64) (*Result, error){}
func (hb *HuobiClient) SellJson(coinType int, price, amount float64) ([]byte, error){}
func (hb *HuobiClient) BuyMarket(coinType int, amount float64) (*Result, error){}
func (hb *HuobiClient) BuyMarketJson(coinType int, amount float64) ([]byte, error){}
func (hb *HuobiClient) SellMarket(coinType int, amount float64) (*Result, error){}
func (hb *HuobiClient) SellMarketJson(coinType int, amount float64) ([]byte, error){}
func (hb *HuobiClient) CancelOrder(coinType, id int) (*Result, error){}
func (hb *HuobiClient) CancelOrderJson(coinType, id int) ([]byte, error){}
func (hb *HuobiClient) GetNewDealOrders(coinType int) ([]OrderTraded, error){}
func (hb *HuobiClient) GetNewDealOrdersJson(coinType int) ([]byte, error){}
```

## API接口说明
本接口尽量保持和火币的API接口一致<br>
每个对应火币网的API接口函数都有2种，一种是直接返回火币api的原始json，另一种是解析后的struct，更方便操作。<br>
例如 func Quotation 返回struct, func QuotationJson返回api的原始json。其他接口也类似。<br>
基于火币的API返回的原始Json，对于数字和字符串没有严格的区分，有的接口返回的json中价格是字符串，有的接口中是浮点型，如果你想保持和火币网一致，可以使用 funcname+Json的函数返回api的原始json，如果想规范点，则使用返回struct的接口，然后把返回的struct format成json，这样输出的json数字和字符串就统一了，不会出现价格可能是字符串也可能是浮点型的问题了。


## 更多
利用放年假的前几天的空暇时间完成，虽然已经对API借口做了测试，但是仍难免有疏漏之处，如果发现bug，请联系QQ：290924805。<br>
也欢迎讨论各种交易策略。
