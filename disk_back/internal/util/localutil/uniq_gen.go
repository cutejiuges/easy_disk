package localutil

import (
	"crypto/sha1"
	"fmt"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/16 下午10:37
 * @FilePath: uniq_gen
 * @Description:
 */

func GetSha1Key(content []byte) string {
	h := sha1.New()
	h.Write(content)
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func GetFileLockKey(key string) string {
	return fmt.Sprintf("{%s}:%s", "fileSha1Key", key)
}
