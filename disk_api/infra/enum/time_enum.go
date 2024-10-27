package enum

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/2 下午7:55
 * @FilePath: time_enum
 * @Description: 时间相关的枚举
 */

type TimeLayout string

const (
	TimeLayoutCompleteMinus    TimeLayout = "2006-01-02 15:04:05"
	TimeLayoutCompleteDivision TimeLayout = "2006/01/02 15:04:05"
	TimeLayoutDateYYMMDD       TimeLayout = "20060102"
)
