package localutils

import "time"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/2 下午7:53
 * @FilePath: time_util
 * @Description: 时间格式化工具
 */

func FormatTime(t time.Time, layout string) string {
	timeStr := t.Format(layout)
	return timeStr
}
