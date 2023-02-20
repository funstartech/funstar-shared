package cookie

import (
	"github.com/funstartech/funstar-proto/go/common"
	"google.golang.org/protobuf/proto"
)

// Encode 编码
func Encode(seq uint32) []byte {
	cookieBytes, _ := proto.Marshal(&common.Cookie{
		Seq: seq,
	})
	return cookieBytes
}

// Decode 解码
func Decode(cookieBytes []byte) uint32 {
	if len(cookieBytes) == 0 {
		return 0
	}
	cookie := &common.Cookie{}
	_ = proto.Unmarshal(cookieBytes, cookie)
	return cookie.Seq
}
