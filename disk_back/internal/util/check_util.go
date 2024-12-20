package util

import "regexp"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/7 上午9:15
 * @FilePath: check_util
 * @Description:
 */

func CheckEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
