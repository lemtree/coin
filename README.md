# coin 数字货币交易平台API
digiccy trading platform API collection, such as Bitcoin, Litecoin. current complete huobi api v3. 

数字货币交易平台API，目前包含了火币网的API V3的大部分接口，包含交易和行情接口，剩余的接口会陆续完善。其他平台的API接口后续补充。
##火币网API V3 

## 安装
```go
go get -u github.com/lemtree/coin/huobi
```


## 实例
#### 实例一：获取当前市场行情
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
	fmt.Printf("当前BTC-CNY报价(golang struct)：%+v \r\n", QuoData)

	//获取BTC-CNY行情，返回火币api的原始JSON
	jsonBlob, err := c.QuotationJson(huobi.COIN_BTC, huobi.CURRENCY_CNY)
	checkError(err)
	fmt.Print("当前BTC-CNY报价(火币API原始json)：", string(jsonBlob))

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

```
对应火币API接口，我们可以获取火币api原始json信息或者解析后的struct结构信息，调用其中一种即可。

#### 实例二：获取我的账户信息
```go
package main

import (
	"fmt"
	"github.com/lemtree/coin/huobi"
)

func main() {
	c := huobi.NewHuobiClient()
	
	// 获取我的账户信息,需要在火币网获取访问秘钥,然后设置你的API key
	c.SetApiKey("your-access-key", "your-secret-key")

	accountInfo, err := c.GetAccountInfo()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("我的资产：%+v", accountInfo)
	}
}
```
默认是人民币市场，包括个人信息，买卖等交易。若要操作美元市场，//设置人民币市场，增加 c.SetMarket("cny") 即可。


## API接口其他补充说明
本接口尽量保持和火币的API接口一致<br>
每个对应火币网的API接口函数都有2种，一种是直接返回火币api的原始json，另一种是解析后的struct，更方便操作。详细请看实例一。<br>
例如获取实时行情： func Quotation() 返回的是解析后的struct, func QuotationJson()返回的是huobi api的原始json。其他接口也类似。<br>
<small>基于火币的API返回的原始Json，对于数字和字符串没有严格的区分，有的接口返回的json中价格是字符串，有的接口中是浮点型。如果客户端是弱类型语言则不需额外处理，但是强类型语言则可能直接导致异常。如果你想保持和火币网一致，可以使用 funcname+Json的函数返回api的原始json，如果想规范点，则使用返回struct的接口，然后把返回的struct format成json，这样输出的json数字和字符串就统一了，不会出现价格可能是字符串也可能是浮点型的问题了。</small>


## 更多
利用放年假的前几天的空暇时间完成，虽然已经对API借口做了测试，但是仍难免有疏漏之处，如果发现bug，请联系QQ：290924805。<br>
也欢迎讨论各种交易策略。
