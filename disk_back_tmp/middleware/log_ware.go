package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"time"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/2 下午3:35
 * @FilePath: log_ware
 * @Description: 在rpc请求前后打印操作日志
 */

func LogMW(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		info := rpcinfo.GetRPCInfo(ctx)

		reqBuf := bytes.NewBuffer([]byte{})
		reqEncoder := json.NewEncoder(reqBuf)
		//防止html转义
		reqEncoder.SetEscapeHTML(false)
		_ = reqEncoder.Encode(req)

		startTime := time.Now()
		klog.CtxInfof(ctx, "_request_in_trace_id_%d method = %v, req = %v", info.Invocation().SeqID(), info.To().Method(), reqBuf.String())
		_ = next(ctx, req, resp)

		respBuf := bytes.NewBuffer([]byte{})
		respEncoder := json.NewEncoder(respBuf)
		respEncoder.SetEscapeHTML(false)
		_ = respEncoder.Encode(resp)
		klog.CtxInfof(ctx, "_request_out_trace_id_%d method = %v, resp = %v, cost = %v", info.Invocation().SeqID(), info.To().Method(), respBuf.String(), time.Since(startTime))
		return
	}
}
