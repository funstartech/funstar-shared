package wujitool

import (
	"fmt"
	"os"
	"regexp"
)

// TransPicURL wuji图片链接转换为cdn链接
func TransPicURL(pic string) string {
	reg := regexp.MustCompile(`/img.*?&`)
	b := reg.FindAllString(pic, 1)
	if len(b) == 0 {
		return ""
	}
	return fmt.Sprintf("https://%v%v", os.Getenv("CDN_DOMAIN"), b[0][:len(b[0])-1])
}
