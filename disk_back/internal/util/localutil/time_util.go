package localutil

import (
	"time"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/2 下午11:32
 * @FilePath: time_util
 * @Description:
 */

func FormatTime(t time.Time, layout string) string {
	timeStr := t.Format(layout)
	return timeStr
}

func ParseTime(s, layout string) (time.Time, error) {
	return time.ParseInLocation(layout, s, time.Local)
}
