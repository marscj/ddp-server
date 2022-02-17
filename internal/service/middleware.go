package service

import (
	"ddp-server/internal/consts"
	"ddp-server/internal/model"
	"ddp-server/utility/response"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type sMiddleware struct {
	LoginUrl string // 登录路由地址
}

var (
	insMiddleware = sMiddleware{
		LoginUrl: "/login",
	}
)

// 中间件管理服务
func Middleware() *sMiddleware {
	return &insMiddleware
}

func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		err  error
		res  interface{}
		code gcode.Code = gcode.CodeOK
	)
	res, err = r.GetHandlerResponse()
	if err != nil {
		code = gerror.Code(err)
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		response.JsonExit(r, code.Code(), err.Error())
	} else {
		response.JsonExit(r, gcode.CodeOK.Code(), "", res)
	}
}

// 自定义上下文对象
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		Session: r.Session,
		Data:    make(g.Map),
	}
	Context().Init(r, customCtx)
	if userEntity := Session().GetUser(r.Context()); userEntity.Id > 0 {
		adminId := g.Cfg().MustGet(r.Context(), "setting.adminId", consts.DefaultAdminId).Uint()
		customCtx.User = &model.ContextUser{
			Id:       userEntity.Id,
			Passport: userEntity.Passport,
			Nickname: userEntity.Nickname,
			Avatar:   userEntity.Avatar,
			IsAdmin:  userEntity.Id == adminId,
		}
	}
	// 将自定义的上下文对象传递到模板变量中使用
	r.Assigns(g.Map{
		"Context": customCtx,
	})
	// 执行下一步请求逻辑
	r.Middleware.Next()
}
