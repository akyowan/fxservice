package adapter

import (
	"bytes"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"time"
	"vcrlibraries/loggers"
)

var Orders map[string]Order

func init() {
	Orders = make(map[string]Order)
}

type Order struct {
	OrderID    string      `json:"orderID,omitempty"`
	Body       string      `json:"body,omitempty"`
	Detail     string      `json:"detail,omitempty"`
	Attach     string      `json:"attach,omitempty"`
	TotalPrice int         `json:"totalPrice,omitempty"`
	CodeURL    string      `json:"CodeURL,omitempty"`
	GoodID     string      `json:"goodID,omitempty"`
	PayMethod  string      `json:"pay_method,omitempty"`
	Status     OrderStatus `json:"status"`
	Created    *time.Time  `json:"create_time,omitempty"`
	Updated    *time.Time  `json:"update_time,omitempty"`
}

type OrderStatus int

const (
	_ OrderStatus = iota
	OrderStatusCreated
	OrderStatusPaying
	OrderStatusPaid
)

func SumitOrder(order *Order, clientIP string) (*Order, error) {
	now := time.Now()
	order.Created = &now

	uReq := &UniFiedOrderReq{
		AppID:          WXAppID,
		MchID:          WXMchID,
		Body:           order.Body,
		Detail:         order.Detail,
		Attach:         order.Attach,
		OutTradeNo:     order.OrderID,
		TotalFee:       order.TotalPrice,
		NonceStr:       RandHex(16),
		NotifyUrl:      WXCallBackUrl,
		SpbillCreateIP: clientIP,
		TradeType:      order.PayMethod,
	}
	uReq.Sign = WXSign(uReq, WXKey)
	uResp, err := UniFiedOrder(uReq)
	if err != nil {
		return nil, err
	}
	if uResp.ReturnCode != "SUCCESS" {
		loggers.Warn.Println(uResp.ReturnMsg)
		return nil, errors.New("wxpay error")
	}
	if uResp.ReturnCode != "SUCCESS" {
		loggers.Warn.Printf("wxpay error code:%s des:%s", uResp.ResultCode, uResp.ReturnMsg)
		return nil, errors.New("wxpay error")
	}
	now = time.Now()
	order.Status = OrderStatusPaying
	order.CodeURL = uResp.CodeURL
	order.Updated = &now

	Orders[order.OrderID] = *order
	return order, nil
}

func RandHex(n int) string {
	u1 := uuid.NewV4()
	d, _ := u1.MarshalBinary()
	buf := new(bytes.Buffer)
	for i := 0; i < n/2; i++ {
		buf.WriteString(fmt.Sprintf("%02x", d[i]))
	}

	return buf.String()
}
