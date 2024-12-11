package mw

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cutejiuges/disk_api/biz/model/user_server"
	"github.com/cutejiuges/disk_api/infra/localutils"
	"github.com/cutejiuges/disk_api/rpc"
	user_back "github.com/cutejiuges/disk_back/kitex_gen/user_server"
	"github.com/hertz-contrib/jwt"
	"net/http"
	"time"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/11 下午10:00
 * @FilePath: jwt_middleware
 * @Description: 用户登录中间件
 */

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = "identity"
)

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:      "cloud_disk",
		Key:        []byte("cloud_disk_secret"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginReq user_server.UserSignInRequest
			if err := c.BindAndValidate(&loginReq); err != nil {
				return nil, err
			}
			var rpcReq user_back.UserSignInRequest
			if err := localutils.Converter(&loginReq, &rpcReq); err != nil {
				return nil, err
			}
			user, err := rpc.GetUserServerClient().UserSignIn(ctx, &rpcReq)
			if err != nil {
				return nil, err
			}
			return user, nil
		},

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(user_back.UserSignInData); ok {
				return jwt.MapClaims{
					IdentityKey:  v.UserName,
					"account_id": v.GetAccountId(),
					"email":      v.GetEmail(),
					"user_name":  v.GetUserName(),
					"admin":      v.GetAdmin(),
				}
			}
			return jwt.MapClaims{}
		},

		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, utils.H{
				"code":    code,
				"token":   token,
				"expire":  expire.Format(time.RFC3339),
				"message": "success",
			})
		},

		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",

		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())
			return e.Error()
		},

		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"code":    code,
				"message": message,
			})
		},

		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return user_back.UserSignInData{
				AccountId: claims["account_id"].(*int64),
				Email:     claims["email"].(*string),
				UserName:  claims["user_name"].(*string),
				Admin:     claims["admin"].(*bool),
			}
		},
	})
	if err != nil {
		panic(err)
	}
}

func AuthAdmin() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		claims := jwt.ExtractClaims(ctx, c)
		admin := claims["admin"].(*bool)
		if !*admin {
			c.JSON(http.StatusOK, utils.H{
				"code": -1,
				"msg":  "当前登陆用户无操作权限",
			})
		}
	}
}
