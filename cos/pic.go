package cos

import (
	"fmt"
	"regexp"
	"strings"
)

// https://cloud.tencent.com/document/product/436/44880

// ThumbPic 返回压缩链接
func ThumbPic(pic string, width int) string {
	if strings.Contains(pic, "imageMogr2") {
		reg := regexp.MustCompile(`http.*?\?`)
		b := reg.FindAllString(pic, 1)
		if len(b) == 0 {
			return pic
		}
		pic = b[0][:len(b[0])-1]
	}
	return fmt.Sprintf("%v?imageMogr2/thumbnail/%vx", pic, width)
}
