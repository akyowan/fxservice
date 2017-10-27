package adapter

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

var (
	WXAppID           string
	WXMchID           string
	WXKey             string
	WXCallBackUrl     string
	WXUniFiedOrderAPI string
)

func init() {
	WXAppID = "wx76f4d728509ad3e1"
	WXMchID = "1480686862"
	WXKey = "YHPtbJvVxdPXHbkahZRbaqNK3EoTtkiP"
	WXCallBackUrl = "https://api.vincross.com/wxpay/callback"
	WXUniFiedOrderAPI = "https://api.mch.weixin.qq.com/pay/unifiedorder"
}

type UniFiedOrderReq struct {
	AppID          string `xml:"appid,omitempty"`
	MchID          string `xml:"mch_id,omitempty"`
	DeviceInfo     string `xml:"device_info,omitempty"`
	NonceStr       string `xml:"nonce_str,omitempty"`
	Sign           string `xml:"sign,omitempty"`
	SignType       string `xml:"sign_type,omitempty"`
	Body           string `xml:"body,omitempty"`
	Detail         string `xml:"detail,omitempty"`
	Attach         string `xml:"attach,omitempty"`
	OutTradeNo     string `xml:"out_trade_no,omitempty"`
	FeeType        string `xml:"fee_type,omitempty"`
	TotalFee       int    `xml:"total_fee,omitempty"`
	SpbillCreateIP string `xml:"spbill_create_ip,omitempty"`
	TimeStart      string `xml:"time_start,omitempty"`
	TimeExpire     string `xml:"time_expire,omitempty"`
	GoodsTag       string `xml:"goods_tag,omitempty"`
	NotifyUrl      string `xml:"notify_url,omitempty"`
	TradeType      string `xml:"trade_type,omitempty"`
	ProductID      string `xml:"product_id,omitempty"`
	LimitPay       string `xml:"limit_pay,omitempty"`
	OpenID         string `xml:"open_id,omitempty"`
	SceneInfo      string `xml:"scene_info,omitempty"`
}

type UniFiedOrderResp struct {
	ReturnCode string `xml:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg,omitempty"`
	AppID      string `xml:"appid,omitempty"`
	MchID      string `xml:"mch_id,omitempty"`
	DeviceInfo string `xml:"device_info,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty"`
	ErrCode    string `xml:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty"`
	TradeType  string `xml:"trade_type,omitempty"`
	PrepayID   string `xml:"prepay_id,omitempty"`
	CodeURL    string `xml:"code_url,omitempty"`
	MWebURL    string `xml:"mweb_url,omitempty"`
}

type PayResult struct {
	ReturnCode         string `xml:"return_code,omitempty"`
	ReturnMsg          string `xml:"return_msg,omitempty"`
	AppID              string `xml:"appid,omitempty"`
	MchID              string `xml:"mch_id,omitempty"`
	DeviceInfo         string `xml:"device_info,omitempty"`
	NonceStr           string `xml:"nonce_str,omitempty"`
	Sign               string `xml:"sign,omitempty"`
	SignType           string `xml:"sign_type,omitempty"`
	ResultCode         string `xml:"result_code,omitempty"`
	ErrCode            string `xml:"err_code,omitempty"`
	ErrCodeDes         string `xml:"err_code_des,omitempty"`
	OpenID             string `xml:"open_id,omitempty"`
	IsSubscribe        string `xml:"is_subscribe,omitempty"`
	TradeType          string `xml:"trade_type,omitempty"`
	BankType           string `xml"bank_type,omitempty"`
	TotalFee           int    `xml:"total_fee,omitempty"`
	SettlementTotalFee int    `xml:"settlement_total_fee,omitempty"`
	FeeType            string `xml:"fee_type,omitempty"`
	CashFee            int    `xml:"cash_fee,omitempty"`
	CashFeeType        string `xml:"cash_fee_type,omitempty"`
	CouponFee          int    `xml:"coupon_fee,omitempty"`
	CouponCount        int    `xml:"coupon_count,omitempty"`
	CouponType0        string `xml:"coupon_type_0,omitempty"`
	CouponID0          string `xml:"coupon_id_0,omitempty"`
	CouponFee0         int    `xml:"coupon_fee_0,omitempty"`
	TransactionID      string `xml:"transaction_id,omitempty"`
	OutTradeNO         string `xml:"out_trade_no,omitempty"`
	Attach             string `xml:"attach,omitempty"`
	TimeEnd            string `xml:"time_end,omitempty"`
}

func WXSign(info *UniFiedOrderReq, key string) string {
	reqMap := make(map[string]interface{})
	reqMap["appid"] = info.AppID
	reqMap["mch_id"] = info.MchID
	reqMap["device_info"] = info.DeviceInfo
	reqMap["nonce_str"] = info.NonceStr
	reqMap["sign_type"] = info.SignType
	reqMap["body"] = info.Body
	reqMap["detail"] = info.Detail
	reqMap["attach"] = info.Attach
	reqMap["out_trade_no"] = info.OutTradeNo
	reqMap["fee_type"] = info.FeeType
	reqMap["total_fee"] = info.TotalFee
	reqMap["spbill_create_ip"] = info.SpbillCreateIP
	reqMap["time_start"] = info.TimeStart
	reqMap["time_expire"] = info.TimeExpire
	reqMap["goods_tag"] = info.GoodsTag
	reqMap["notify_url"] = info.NotifyUrl
	reqMap["trade_type"] = info.TradeType
	reqMap["product_id"] = info.ProductID
	reqMap["limit_pay"] = info.LimitPay
	reqMap["open_id"] = info.OpenID
	reqMap["scene_info"] = info.SceneInfo
	sortedKeys := make([]string, 0)
	for k, _ := range reqMap {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)
	var signStrings string
	for _, k := range sortedKeys {
		value := fmt.Sprintf("%v", reqMap[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}
	if key != "" {
		signStrings = signStrings + "key=" + key
	}
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))

	return upperSign
}

func UniFiedOrder(r *UniFiedOrderReq) (*UniFiedOrderResp, error) {
	data, err := xml.Marshal(r)
	if err != nil {
		return nil, err
	}
	dataStr := strings.Replace(string(data), "UniFiedOrderReq", "xml", -1)
	body := bytes.NewBuffer([]byte(dataStr))

	resp, err := http.Post(WXUniFiedOrderAPI, "application/xml;charset=utf8", body)
	if err != nil {
		return nil, err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respData UniFiedOrderResp
	if xml.Unmarshal(respBody, &respData); err != nil {
		return nil, err
	}
	return &respData, nil
}
