package localutil

import (
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/2 下午11:30
 * @FilePath: localutil
 * @Description:
 */

func SonicMarshal(v any) []byte {
	bytes, _ := sonic.Marshal(v)
	return bytes
}

func SonicUnmarshal(data []byte, v any) error {
	return sonic.Unmarshal(data, v)
}

func SonicMarshalString(v any) string {
	bytes, _ := sonic.Marshal(v)
	return string(bytes)
}

func SonicUnmarshalString(data string, v any) error {
	return sonic.UnmarshalString(data, v)
}

func Converter(source, dest any) error {
	bytes, err := sonic.Marshal(source)
	if err != nil {
		klog.Error("json marshal error: ", err)
		return err
	}

	err = sonic.Unmarshal(bytes, dest)
	if err != nil {
		klog.Error("json unmarshal error: ", err)
		return err
	}
	return nil
}
