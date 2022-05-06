package system

import (
	"gin/global"
	"gin/model/common/request"
	"gin/model/common/response"
	"gin/model/system"
	systemReq "gin/model/system/request"
	systemRes "gin/model/system/response"
	"gin/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	//"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type BaseApi struct {}

//
// Register
// @Description:
// @receiver b
// @param c
//
func (b *BaseApi) Register(c *gin.Context) {
	var r systemReq.Register
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user := &system.SysUser{Username: r.Username, NickName: r.NickName, Password: r.Password}
	err, userReturn := userService.Register(*user)

	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)
	}
}

//
// GetUserList
// @Description:
// @receiver b
// @param c
//
func (b *BaseApi) GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := userService.GetUserInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (b *BaseApi) FindUserById(c *gin.Context) {

	if id, err :=strconv.Atoi(c.Param("id"));err!=nil{
		if err,ReqUser:=userService.FindUserById(id);err!=nil{
			global.GVA_LOG.Error("获取失败!",zap.Error(err))
			response.FailWithMessage("获取失败",c)
		}else {
			global.GVA_LOG.Info("获取成功!",zap.Int("id", id))
			response.OkWithDetailed(gin.H{"userInfo":ReqUser},"获取成功",c)
		}
	}
	
	response.FailWithMessage("类型转换错误",c)
}