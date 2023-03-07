package ghttp

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/funstartech/funstar-shared/cutils"
	"github.com/funstartech/funstar-shared/log"
	jsoniter "github.com/json-iterator/go"
)

// Post 发送post请求
func Post(path string, req interface{}, rsp interface{}) error {
	mJson, err := jsoniter.Marshal(req)
	if err != nil {
		return err
	}
	resp, err := http.Post(path, "application/json", bytes.NewReader(mJson))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err = jsoniter.Unmarshal(body, rsp); err != nil {
		return err
	}
	log.Debugf("http request: %v, req: %s, rsp: %s", path, mJson, cutils.Obj2Json(rsp))
	return nil
}
