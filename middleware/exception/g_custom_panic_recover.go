package exception

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func SetUp() gin.HandlerFunc{
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case BizError:

					bizE := err.(*BizError)

					handleBizError(ctx, bizE)
					t.Ctx.Input.SetData("___status", bizE.Code)
					t.Finish()
					return
				case SystemError:
					systemError := err.(*b_error.SystemError)

					beego.Error("System Error:", systemError.Error())
					t.Abort(strconv.Itoa(systemError.Code))
					return
				}
			}
		}()
		context.Next()
	}
}

func handleBizError(ctx *gin.Context, bizE *BizError) {
	logs.Warn("[BIZE] biz exception: code", bizE.Code, " message ", bizE.Message)
	rd := &ResponseData{
		Ret:        bizE.Code,
		Message:    bizE.Message,
		ServerTime: time.Now().UnixNano() / 1000000,
	}
	hasIndent := beego.BConfig.RunMode != beego.PROD
	jsonpCallback := ctx.Request.Form.Get("callback")
	if jsonpCallback != "" {
		ctx.Output.JSONP(rd, hasIndent)
	} else {
		ctx.Output.JSON(rd, hasIndent, false)
	}

}
