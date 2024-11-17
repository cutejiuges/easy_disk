package util

import (
	"crypto/sha256"
	"fmt"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/16 下午10:37
 * @FilePath: uniq_gen
 * @Description:
 */

func GetSha256Key(content []byte) string {
	h := sha256.New()
	h.Write(content)
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
