package enum

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/2 下午11:33
 * @FilePath: time_enum
 * @Description:
 */

type TimeLayout string

const (
	TimeLayoutCompleteMinus    TimeLayout = "2006-01-02 15:04:05"
	TimeLayoutCompleteDivision TimeLayout = "2006/01/02 15:04:05"
	TimeLayoutDateYYMMDD       TimeLayout = "20060102"
)
