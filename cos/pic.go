package cos

import (
	"fmt"
)

// https://cloud.tencent.com/document/product/436/44880

// ThumbPic 返回压缩链接
func ThumbPic(pic string, width int) string {
	return fmt.Sprintf("%v?imageMogr2/thumbnail/%vx", pic, width)
}
