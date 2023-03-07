package wxpay

import (
	"context"
	"strconv"

	"github.com/funstartech/funstar-shared/ghttp"
)

// CloseOrderReq 关闭订单请求
type CloseOrderReq struct {
	OutTradeNo string `json:"out_trade_no"`
	SubMchID   string `json:"sub_mch_id"`
}

// CloseOrderRsp 关闭订单回包
type CloseOrderRsp struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	Respdata struct {
		ReturnCode string `json:"return_code"`
		ReturnMsg  string `json:"return_msg"`
		AppID      string `json:"appid"`
		MchID      string `json:"mch_id"`
		SubAppID   string `json:"sub_appid"`
		SubMchID   string `json:"sub_mch_id"`
		NonceStr   string `json:"nonce_str"`
		Sign       string `json:"sign"`
		ResultCode string `json:"result_code"`
	} `json:"respdata"`
}

// CloseOrder 关闭订单
func CloseOrder(ctx context.Context, orderID uint64) (*CloseOrderRsp, error) {
	req := &CloseOrderReq{
		OutTradeNo: strconv.FormatUint(orderID, 10),
		SubMchID:   subMchID,
	}
	rsp := &CloseOrderRsp{}
	if err := ghttp.Post(closeOrderPath, req, rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}
