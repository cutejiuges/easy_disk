package errno

import (
	"errors"
	"fmt"
	"github.com/cutejiuges/disk_back/kitex_gen/base"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/1 上午1:13
 * @FilePath: err_wrapper
 * @Description:
 */

type BizError struct {
	Code  int32
	Msg   string
	Erred error
}

func (e *BizError) Error() string {
	if e.Erred == nil {
		return fmt.Sprintf("code = %d, Msg = %s", e.Code, e.Msg)
	}
	return fmt.Sprintf("code = %d, msg = %s, erred = %v", e.Code, e.Msg, e.Erred)
}

func (e *BizError) Wrap(err error) *BizError {
	e.Erred = err
	return e
}

func NewBaseResp(err error) *base.BaseResp {
	resp := base.NewBaseResp()
	var e *BizError
	if errors.As(err, &e) {
		resp.SetStatusCode(e.Code)
		resp.SetStatusMsg(e.Msg)
	} else {
		resp.SetStatusCode(int32(ErrCodeInternal))
		resp.SetStatusMsg(err.Error())
	}
	return resp
}

func NewBaseRespWithOK() *base.BaseResp {
	resp := base.NewBaseResp()
	resp.SetStatusCode(int32(ErrCodeOK))
	resp.SetStatusMsg(ErrMsgMap[int32(ErrCodeOK)])
	return resp
}

func NewBaseRespWithMsg(msg string) *base.BaseResp {
	resp := base.NewBaseResp()
	resp.SetStatusMsg(msg)
	resp.SetStatusCode(int32(ErrCodeCommon))
	return resp
}
