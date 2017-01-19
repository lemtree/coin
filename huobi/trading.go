/**
 * 火币 www.huobi.com API 交易接口
 * Auth HuHeKun
 * Date 2017-01-17
 * 目前 trade_password trade_id 还未实现，以及BTC\LTC提现等暂时未用到，如有需求后续完成
 */
package huobi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/lemtree/funcs"
)

var huobiDebug bool

type HuobiClient struct {
	httpClient *http.Client
	AccessKey  string
	SecretKey  string
	Market     string
	//	trade_password
	UserAgent string
}

func (hb *HuobiClient) SetApiKey(AccessKey, SecretKey string) {
	hb.AccessKey = AccessKey
	hb.SecretKey = SecretKey
}

// 默认市场为人民币市场
func NewHuobiClient() *HuobiClient {
	client := &http.Client{}
	return &HuobiClient{httpClient: client, AccessKey: "", SecretKey: "", Market: "cny", UserAgent: "HHK Client"}
}

/**
 * 生成请求的参数
 */
func (hb *HuobiClient) generateParameter(parameter map[string]string) string {
	v := url.Values{}
	for key, val := range parameter {
		v.Set(key, val)
	}
	v.Set("access_key", hb.AccessKey)
	v.Set("created", strconv.Itoa(int(time.Now().Unix())))
	v.Set("secret_key", hb.SecretKey)
	v.Set("sign", tools.MD5([]byte(v.Encode())))
	v.Del("secret_key")
	v.Set("market", hb.Market)
	return v.Encode()
}

/**
 * 获取个人账户信息
 */
func (hb *HuobiClient) GetAccountInfo() (*AccountInfo, error) {
	account := &AccountInfo{}
	jsonBlob, err := hb.GetAccountInfoJson()
	//jsonBlob = []byte(`{"total":"2163.32","net_asset":"2163.32","available_cny_display":"2163.32","available_btc_display":"1.23000","available_ltc_display":"2.2300","frozen_cny_display":"3.23","frozen_btc_display":"4.23","frozen_ltc_display":"5.23","loan_cny_display":"6.23","loan_btc_display":"7.23","loan_ltc_display":"8.23"}`)
	if err != nil {
		return account, err
	}
	err = json.Unmarshal(jsonBlob, &account)
	if err != nil {
		return account, err
	}
	return account, nil

}

/**
 * 获取个人账户信息(返回API的json字符串)
 */
func (hb *HuobiClient) GetAccountInfoJson() ([]byte, error) {
	p := map[string]string{"method": "get_account_info"}
	parameter := hb.generateParameter(p)
	return hb.SendTradingRequest(parameter)
}

/**
 * 获取当前所有正在进行的委托
 * coinType 数字货币 BTC\LTC
 */
func (hb *HuobiClient) GetOrders(coinType coinT) ([]Order, error) {
	orders := []Order{}
	jsonBlob, err := hb.GetOrdersJson(coinType)
	if err != nil {
		return orders, err
	}
	err = json.Unmarshal(jsonBlob, &orders)
	if err != nil {
		return orders, err
	}
	return orders, nil
}

/**
 * 获取当前所有正在进行的委托(返回API的json字符串)
 * coinType 数字货币 BTC\LTC
 */
func (hb *HuobiClient) GetOrdersJson(coinType coinT) ([]byte, error) {
	p := map[string]string{"method": "get_orders", "coin_type": formatCoinType(coinType)}
	parameter := hb.generateParameter(p)
	return hb.SendTradingRequest(parameter)
}

/**
 * 获取委托订单详情
 * coinType 数字货币 BTC\LTC
 * id 订单id
 */
func (hb *HuobiClient) OrderInfo(coinType coinT, id int) (*OrderInfo, error) {
	orderInfo := &OrderInfo{}
	jsonBlob, err := hb.OrderInfoJson(coinType, id)
	//	jsonBlob = []byte(`{"id":3748640502,"type":1,"order_price":"3000.00","order_amount":"0.0100","processed_price":"1.11","processed_amount":"2.22","vot":"3.33","fee":"4.44","total":"5.55","status":8}`)
	if err != nil {
		return orderInfo, err
	}
	err = json.Unmarshal(jsonBlob, orderInfo)
	if err != nil {
		return orderInfo, err
	}
	return orderInfo, nil
}

/**
 * 获取委托订单详情(返回API的json字符串)
 * coinType 数字货币 BTC\LTC
 * id 订单id
 */
func (hb *HuobiClient) OrderInfoJson(coinType coinT, id int) ([]byte, error) {
	p := map[string]string{"method": "order_info", "coin_type": formatCoinType(coinType), "id": strconv.Itoa(id)}
	parameter := hb.generateParameter(p)
	return hb.SendTradingRequest(parameter)
}

/**
 * 限价买入
 * coinType 买入的数字货币 BTC\LTC
 * price 买价，期望购买的价位
 * amount 数字货币数量，要购买数字货币的数量
 */
func (hb *HuobiClient) Buy(coinType coinT, price, amount float64) (*Result, error) {
	result := &Result{}
	jsonBlob, err := hb.BuyJson(coinType, price, amount)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(jsonBlob, result)
	if err != nil {
		return result, err
	}
	return result, nil
}

/**
 * 限价买入(返回API的json字符串)
 * coinType 买入的数字货币 BTC\LTC
 * price 买价，期望购买的价位
 * amount 数字货币数量，要购买数字货币的数量
 */
func (hb *HuobiClient) BuyJson(coinType coinT, price, amount float64) ([]byte, error) {
	p := map[string]string{"method": "buy", "coin_type": formatCoinType(coinType), "price": formatPrice(price), "amount": formatCoinAmout(amount)}
	parameter := hb.generateParameter(p)
	return hb.SendTradingRequest(parameter)
}

/**
 * 限价卖出
 * coinType 卖出的数字货币 BTC\LTC
 * price 卖价，期望卖出的价位
 * amount 数字货币数量，卖出的coinType的数量
 */
func (hb *HuobiClient) Sell(coinType coinT, price, amount float64) (*Result, error) {
	result := &Result{}
	jsonBlob, err := hb.SellJson(coinType, price, amount)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(jsonBlob, result)
	if err != nil {
		return result, err
	}
	return result, nil
}

/**
 * 限价卖出(返回API的json字符串)
 * coinType 卖出的数字货币 BTC\LTC
 * price 卖价，期望卖出的价位
 * amount 数字货币数量，卖出的coinType的数量
 */
func (hb *HuobiClient) SellJson(coinType coinT, price, amount float64) ([]byte, error) {
	p := map[string]string{"method": "sell", "coin_type": formatCoinType(coinType), "price": formatPrice(price), "amount": formatCoinAmout(amount)}
	parameter := hb.generateParameter(p)
	return hb.SendTradingRequest(parameter)
}

/**
 * 市价买入
 * coinType 买入的数字货币 BTC\LTC
 * amount 法币金额。要买入多少钱的数字货币，单位CNY/USD，金额精确到0.01元，最少为1元
 */
func (hb *HuobiClient) BuyMarket(coinType coinT, amount float64) (*Result, error) {
	result := &Result{}
	jsonBlob, err := hb.BuyMarketJson(coinType, amount)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(jsonBlob, result)
	if err != nil {
		return result, err
	}
	return result, nil
}

/**
 * 市价买入(返回API的json字符串)
 * coinType 买入的数字货币 BTC\LTC
 * amount 法币金额。要买入多少钱的数字货币，单位CNY/USD，金额精确到0.01元，最少为1元
 */
func (hb *HuobiClient) BuyMarketJson(coinType coinT, amount float64) ([]byte, error) {
	p := map[string]string{"method": "buy_market", "coin_type": formatCoinType(coinType), "amount": formatPrice(amount)}
	parameter := hb.generateParameter(p)
	return hb.SendTradingRequest(parameter)
}

/**
 * 市价卖出
 * coinType 卖出的数字货币 BTC\LTC
 * amount 要卖出的btc/ltc数量，精确到0.0001个币, 但是最小卖出为0.001币（比如卖出3.1024个币）
 */
func (hb *HuobiClient) SellMarket(coinType coinT, amount float64) (*Result, error) {
	result := &Result{}
	jsonBlob, err := hb.SellMarketJson(coinType, amount)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(jsonBlob, result)
	if err != nil {
		return result, err
	}
	return result, nil
}

/**
 * 市价卖出(返回API的json字符串)
 * coinType 卖出的数字货币 BTC\LTC
 * amount 要卖出的btc/ltc数量，精确到0.0001个币
 */
func (hb *HuobiClient) SellMarketJson(coinType coinT, amount float64) ([]byte, error) {
	p := map[string]string{"method": "sell_market", "coin_type": formatCoinType(coinType), "amount": formatCoinAmout(amount)}
	parameter := hb.generateParameter(p)
	return hb.SendTradingRequest(parameter)
}

/**
 * 取消委托单
 * coinType 数字货币 BTC\LTC
 * id 要取消的委托id
 */
func (hb *HuobiClient) CancelOrder(coinType coinT, id int) (*Result, error) {
	result := &Result{}
	jsonBlob, err := hb.CancelOrderJson(coinType, id)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(jsonBlob, result)
	if err != nil {
		return result, err
	}
	return result, nil
}

/**
 * 取消委托单(返回API的json字符串)
 * coinType 数字货币 BTC\LTC
 * id 要取消的委托id
 */
func (hb *HuobiClient) CancelOrderJson(coinType coinT, id int) ([]byte, error) {
	p := map[string]string{"method": "cancel_order", "coin_type": formatCoinType(coinType), "id": strconv.Itoa(id)}
	parameter := hb.generateParameter(p)
	return hb.SendTradingRequest(parameter)
}

/**
 * 查询个人最新10条成交订单
 * coinType 数字货币 BTC\LTC
 */
func (hb *HuobiClient) GetNewDealOrders(coinType coinT) ([]OrderTraded, error) {
	orders := []OrderTraded{}
	jsonBlob, err := hb.GetNewDealOrdersJson(coinType)
	if err != nil {
		return orders, err
	}
	err = json.Unmarshal(jsonBlob, &orders)
	if err != nil {
		return orders, err
	}
	return orders, nil
}

/**
 * 查询个人最新10条成交订单(返回API的json字符串)
 * coinType 数字货币 BTC\LTC
 */
func (hb *HuobiClient) GetNewDealOrdersJson(coinType coinT) ([]byte, error) {
	p := map[string]string{"method": "get_new_deal_orders", "coin_type": formatCoinType(coinType)}
	parameter := hb.generateParameter(p)
	return hb.SendTradingRequest(parameter)
}

/**
 * 发送交易请求到api接口获取数据，用于发起需要用户交易秘钥的请求
 */
func (hb *HuobiClient) SendTradingRequest(parameter string) ([]byte, error) {
	if len(hb.AccessKey) < 1 || len(hb.SecretKey) < 1 {
		return []byte{}, errors.New("AccessKey和SecretKey不能为空")
	}
	url := API_V3_URI
	return hb.SendRequest(url, parameter)
}

/**
 * 发送请求到api接口获取数据
 * uri 请求的uri
 * parameter 请求的数据
 */
func (hb *HuobiClient) SendRequest(uri, parameter string) ([]byte, error) {
	body := ioutil.NopCloser(strings.NewReader(parameter)) //把form数据编码
	req, _ := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //post方式需要
	req.Header.Add("User-Agent", hb.UserAgent)
	resp, err := hb.httpClient.Do(req) //发送请求
	defer resp.Body.Close()            //一定要关闭resp.Body
	data, err := ioutil.ReadAll(resp.Body)
	return data, err
}

/**
 * 设置http header的User-Agent
 * userAgent User-Agent信息
 * parameter 请求的数据
 */
func (hb *HuobiClient) SetUserAgent(userAgent string) {
	hb.UserAgent = userAgent
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

/**
 * 格式价格为string，最大2位小数。hubi的金额精确到2位小数
 * 所有提交post的参数均为string类型
 */
func formatPrice(price float64) string {
	return strconv.FormatFloat(price, 'f', 2, 64)
}

/**
 * 格式化比特币数量为string，最大4位小数。huobi所有币种精确到4位小数
 * 所有提交post的参数均为string类型
 */
func formatCoinAmout(amount float64) string {
	return strconv.FormatFloat(amount, 'f', 4, 64)
}

/**
 * 格式化数字货币类型为string。
 * 所有提交post的参数均为string类型
 */
func formatCoinType(coinType coinT) string {
	return strconv.Itoa(int(coinType))
}
