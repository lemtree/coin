/**
 * 火币 www.huobi.com API 行情接口
 * Auth HuHeKun
 * Date 2017-01-17
 */
package huobi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

var KlineUrl = HOST + "%s" + "/" + "%s" + "_kline_" + "%03d" + "_json.js"
var QuotationUrl = HOST + "%s" + "/" + "ticker_%s_json.js"
var DepthUrl = HOST + "%s/depth_%s_%d.js"
var RealTimeTransactionUrl = HOST + "%s/detail_%s_json.js"

/**
 * 获取K线图
 * coinType 数字货币 BTC\LTC
 * currency 法币类型 CNY\USD
 * period 周期
 * length 数据条数
 */
func (hb *HuobiClient) Kline(coinType CoinT, currencyType CurrencyT, period PeriodT, length int) ([]Kline, error) {
	klines := []Kline{}
	dataMap := [][]interface{}{}
	jsonBlob, err := hb.KlineJson(coinType, currencyType, period, length)
	if err != nil {
		return klines, err
	}
	err = json.Unmarshal(jsonBlob, &dataMap)
	if err != nil {
		return klines, err
	}

	for _, row := range dataMap {
		kline := Kline{}
		kline.DateTime, _ = strconv.Atoi(row[0].(string))
		kline.Open = row[1].(float64)
		kline.High = row[2].(float64)
		kline.Low = row[3].(float64)
		kline.Close = row[4].(float64)
		kline.Volue = row[5].(float64)
		klines = append(klines, kline)
	}

	return klines, nil
}

/**
 * 获取K线图
 * coinType 数字货币 BTC\LTC
 * currency 法币类型 CNY\USD
 * period 周期
 * length 数据条数
 */
func (hb *HuobiClient) KlineJson(coinType CoinT, currencyType CurrencyT, period PeriodT, length int) ([]byte, error) {
	uri := hb.generateKlineApiUrl(coinType, currencyType, period)
	v := url.Values{}
	v.Set("length", strconv.Itoa(length))
	parameter := v.Encode()
	return hb.SendRequest(uri, parameter)
}

// 根据cointype返回coin的小写名称 btc/ltc
func (hb *HuobiClient) getCoinName(coinType CoinT) string {
	if coinName, ok := coinNames[coinType]; ok {
		return coinName
	}
	return ""
}

// 根据currency返回market cny:staticmarket usd:usdmarket
func (hb *HuobiClient) getMarketName(currencyType CurrencyT) string {
	if marketName, ok := marketNames[currencyType]; ok {
		return marketName
	}
	return ""
}

func (hb *HuobiClient) generateKlineApiUrl(coinType CoinT, currencyType CurrencyT, period PeriodT) string {
	coinName := hb.getCoinName(coinType)
	marketName := hb.getMarketName(currencyType)
	return fmt.Sprintf(KlineUrl, marketName, coinName, period)
}

/**
 * BTC-CNY K线图
 */
func (hb *HuobiClient) KlineBtcCny(period PeriodT, length int) ([]Kline, error) {
	return hb.Kline(COIN_BTC, CURRENCY_CNY, period, length)
}

/**
 * BTC-USD K线图
 */
func (hb *HuobiClient) KlineBtcUsd(period PeriodT, length int) ([]Kline, error) {
	return hb.Kline(COIN_BTC, CURRENCY_USD, period, length)
}

/**
 * LTC-USD K线图
 */
func (hb *HuobiClient) KlineLtcCny(period PeriodT, length int) ([]Kline, error) {
	return hb.Kline(COIN_LTC, CURRENCY_CNY, period, length)
}

/**
 * 实时行情
 * coinType 数字货币 BTC\LTC
 * currency 法币类型 CNY\USD
 */
func (hb *HuobiClient) Quotation(coinType CoinT, currencyType CurrencyT) (*RealTimeQuotation, error) {
	realTimeQuotation := &RealTimeQuotation{}
	jsonBlob, err := hb.QuotationJson(coinType, currencyType)
	if err != nil {
		return realTimeQuotation, err
	}
	err = json.Unmarshal(jsonBlob, realTimeQuotation)
	if err != nil {
		return realTimeQuotation, err
	}
	return realTimeQuotation, nil
}

/**
 * 实时行情(返回API的json字符串)
 * coinType 数字货币 BTC\LTC
 * currency 法币类型 CNY\USD
 */
func (hb *HuobiClient) QuotationJson(coinType CoinT, currencyType CurrencyT) ([]byte, error) {
	coinName := hb.getCoinName(coinType)
	marketName := hb.getMarketName(currencyType)
	uri := fmt.Sprintf(QuotationUrl, marketName, coinName)
	return hb.SendRequest(uri, "")
}

/**
 * 交易深度
 * coinType 数字货币 BTC\LTC
 * currency 法币类型 CNY\USD
 * length 返回的数据条数
 */
func (hb *HuobiClient) Depth(coinType CoinT, currencyType CurrencyT, length int) (*Depth, error) {
	depths := &Depth{}
	jsonBlob, err := hb.DepthJson(coinType, currencyType, length)
	if err != nil {
		return depths, err
	}
	err = json.Unmarshal(jsonBlob, depths)
	if err != nil {
		return depths, err
	}
	return depths, nil
}

/**
 * 交易深度(返回API的json字符串)
 * coinType 数字货币 BTC\LTC
 * currency 法币类型 CNY\USD
 * length 返回的数据条数
 */
func (hb *HuobiClient) DepthJson(coinType CoinT, currencyType CurrencyT, length int) ([]byte, error) {
	coinName := hb.getCoinName(coinType)
	marketName := hb.getMarketName(currencyType)
	uri := fmt.Sprintf(DepthUrl, marketName, coinName, length)
	return hb.SendRequest(uri, "")
}

/**
 * 买卖盘实时成交数据
 * coinType 数字货币 BTC\LTC
 * currency 法币类型 CNY\USD
 */
func (hb *HuobiClient) RealTimeTransaction(coinType CoinT, currencyType CurrencyT) (*RealTimeTransactionData, error) {
	transactionData := &RealTimeTransactionData{}
	jsonBlob, err := hb.RealTimeTransactionJson(coinType, currencyType)
	if err != nil {
		return transactionData, err
	}
	err = json.Unmarshal(jsonBlob, transactionData)
	if err != nil {
		return transactionData, err
	}
	return transactionData, nil
}

/**
 * 买卖盘实时成交数据
 * coinType 数字货币 BTC\LTC
 * currency 法币类型 CNY\USD
 */
func (hb *HuobiClient) RealTimeTransactionJson(coinType CoinT, currencyType CurrencyT) ([]byte, error) {
	coinName := hb.getCoinName(coinType)
	marketName := hb.getMarketName(currencyType)
	uri := fmt.Sprintf(RealTimeTransactionUrl, marketName, coinName)
	return hb.SendRequest(uri, "")
}
