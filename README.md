# coin 数字货币交易平台API
digiccy trading platform API collection, such as Bitcoin, Litecoin. current complete huobi api v3. 

数字货币交易平台API，目前当前包含了火币网的API V3的大部分接口，包含交易和行情接口，剩余的接口会陆续完善。其他平台的API接口后续补充。
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

	//获取BTC-CNY行情，已经解析后的json
	QuoData, err := c.Quotation(huobi.COIN_BTC, huobi.CURRENCY_CNY)
	checkError(err)
	fmt.Println(QuoData)

	//获取BTC-CNY行情，API返回的JSON
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
## API接口说明
本接口尽量保持和火币的API接口一致<br>
每个对应火币网的API接口函数都有2种，一种是直接返回火币api的原始json，另一种是解析后的struct，更方便操作。<br>
例如 func Quotation 返回struct, func QuotationJson返回api的原始json。其他接口也类似。
基于火币的API返回的原始Json，对于数字和字符串没有严格的区分，价格和金额有的接口是数字有的是字符串，如果你想保持和火币网一致，可以直接返回api的原始json，如果想更规范点，可以使用解析后的struct，在format成json，这样数字和字符串就统一了。


## 更多
利用放假前几天时间空暇完成，虽然已经对API借口做了测试，但是仍难免有疏漏之处，如果发现，请联系QQ：290924805
