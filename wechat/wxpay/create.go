package wxpay

// https://developers.weixin.qq.com/miniprogram/dev/wxcloudrun/src/development/pay/order/unified.html

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/funstartech/funstar-proto/go/wxpay"
	"github.com/funstartech/funstar-shared/cutils"
	"github.com/funstartech/funstar-shared/gheader"
	"github.com/funstartech/funstar-shared/ghttp"
)

const (
	timeFormat = "20060102150405"
	// 商家名称
	mchName = "繁星赏"
	// 子商户号
	subMchID = "1638991683"
	// 请求路径
	createOrderPath = "http://api.weixin.qq.com/_/pay/unifiedorder"
)

type container struct {
	Service string `json:"service"`
	Path    string `json:"path"`
}

// CreateOrderReq 创建订单请求
type CreateOrderReq struct {
	CallbackType   int        `json:"callback_type"`
	EnvID          string     `json:"env_id"`
	FunctionName   string     `json:"function_name"`
	Container      *container `json:"container"`
	SubMchID       string     `json:"sub_mch_id"`
	DeviceInfo     string     `json:"device_info"`
	NonceStr       string     `json:"nonce_str"`
	Body           string     `json:"body"`
	Detail         string     `json:"detail"`
	Attach         string     `json:"attach"`
	OutTradeNo     string     `json:"out_trade_no"`
	FeeType        string     `json:"fee_type"`
	TotalFee       int32      `json:"total_fee"`
	SpbillCreateIP string     `json:"spbill_create_ip"`
	TimeStart      string     `json:"time_start"`
	TimeExpire     string     `json:"time_expire"`
	GoodsTag       string     `json:"goods_tag"`
	TradeType      string     `json:"trade_type"`
	LimitPay       string     `json:"limit_pay"`
	Openid         string     `json:"openid"`
	SubOpenid      string     `json:"sub_openid"`
	Receipt        string     `json:"receipt"`
}

// CreateOrderRsp 创建订单回包
type CreateOrderRsp struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	Respdata struct {
		ReturnCode string           `json:"return_code"`
		ReturnMsg  string           `json:"return_msg"`
		Payment    *wxpay.WxPayment `json:"payment"`
		AppID      string           `json:"appid"`
		MchID      string           `json:"mch_id"`
		SubAppID   string           `json:"sub_appid"`
		SubMchID   string           `json:"sub_mch_id"`
		DeviceInfo string           `json:"device_info"`
		NonceStr   string           `json:"nonce_str"`
		Sign       string           `json:"sign"`
		ResultCode string           `json:"result_code"`
		ErrCode    string           `json:"err_code"`
		ErrCodeDes string           `json:"err_code_des"`
		TradeType  string           `json:"trade_type"`
		PrepayID   string           `json:"prepay_id"`
	} `json:"respdata"`
}

// CreateOrder 创建订单
func CreateOrder(ctx context.Context, orderID, summary string, price int32, callbackPath string) (*CreateOrderRsp, error) {
	openid, err := gheader.GetWxOpenID(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	req := CreateOrderReq{
		EnvID:        os.Getenv("CBR_ENV_ID"),
		CallbackType: 2,
		Container: &container{
			Service: "gateway",
			Path:    callbackPath,
		},
		SubMchID:       subMchID,
		DeviceInfo:     "WEB",
		NonceStr:       cutils.RandString(32),
		Body:           fmt.Sprintf("%s-%s", mchName, summary),
		OutTradeNo:     orderID,
		TotalFee:       price,
		SpbillCreateIP: cutils.GetIP(),
		TimeStart:      now.Format(timeFormat),
		TimeExpire:     now.Add(time.Minute * 5).Format(timeFormat),
		TradeType:      "JSAPI",
		Openid:         openid,
	}
	rsp := &CreateOrderRsp{}
	if err := ghttp.Post(createOrderPath, req, rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}
