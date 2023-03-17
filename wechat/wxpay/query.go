package wxpay

import (
	"context"
	"strconv"

	"github.com/funstartech/funstar-shared/ghttp"
)

type tradeState string

const (
	// TradeStateSuccess 支付成功
	TradeStateSuccess tradeState = "SUCCESS"
	// TradeStateRefund 转入退款
	TradeStateRefund tradeState = "REFUND"
	// TradeStateNotPay 未支付
	TradeStateNotPay tradeState = "NOTPAY"
	// TradeStateClosed 已关闭
	TradeStateClosed tradeState = "CLOSED"
	// TradeStateRevoked 已撤销（刷卡支付）
	TradeStateRevoked tradeState = "REVOKED"
	// TradeStateUserPaying 用户支付中
	TradeStateUserPaying tradeState = "USERPAYING"
	// TradeStatePayError 支付失败(其他原因，如银行返回失败)
	TradeStatePayError tradeState = "PAYERROR"
)

// QueryOrderReq 查询订单请求
type QueryOrderReq struct {
	OutTradeNo string `json:"out_trade_no"`
	SubMchID   string `json:"sub_mch_id"`
}

// QueryOrderRsp 查询订单回包
type QueryOrderRsp struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	Respdata struct {
		ReturnCode     string     `json:"return_code"`
		ReturnMsg      string     `json:"return_msg"`
		Appid          string     `json:"appid"`
		MchId          string     `json:"mch_id"`
		SubAppid       string     `json:"sub_appid"`
		SubMchId       string     `json:"sub_mch_id"`
		NonceStr       string     `json:"nonce_str"`
		Sign           string     `json:"sign"`
		ResultCode     string     `json:"result_code"`
		Openid         string     `json:"openid"`
		IsSubscribe    string     `json:"is_subscribe"`
		SubOpenid      string     `json:"sub_openid"`
		TradeType      string     `json:"trade_type"`
		TradeState     tradeState `json:"trade_state"`
		BankType       string     `json:"bank_type"`
		TotalFee       int        `json:"total_fee"`
		FeeType        string     `json:"fee_type"`
		CashFee        int        `json:"cash_fee"`
		CashFeeType    string     `json:"cash_fee_type"`
		CouponIdList   []string   `json:"coupon_id_list"`
		CouponTypeList []string   `json:"coupon_type_list"`
		CouponFeeList  []int32    `json:"coupon_fee_list"`
		TransactionId  string     `json:"transaction_id"`
		OutTradeNo     string     `json:"out_trade_no"`
		Attach         string     `json:"attach"`
		TimeEnd        string     `json:"time_end"`
		TradeStateDesc string     `json:"trade_state_desc"`
	} `json:"respdata"`
}

// QueryOrder 查询订单
func QueryOrder(ctx context.Context, orderID uint64) (*QueryOrderRsp, error) {
	req := &QueryOrderReq{
		OutTradeNo: strconv.FormatUint(orderID, 10),
		SubMchID:   subMchID,
	}
	rsp := &QueryOrderRsp{}
	if err := ghttp.Post(queryOrderPath, req, rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}
