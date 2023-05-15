package myFirstPackage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/myFirstPackage"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    myFirstPackageReq "github.com/flipped-aurora/gin-vue-admin/server/model/myFirstPackage/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type StudentGirlfriendApi struct {
}

var StuGFService = service.ServiceGroupApp.MyfirstpackageServiceGroup.StudentGirlfriendService


// CreateStudentGirlfriend 创建StudentGirlfriend
// @Tags StudentGirlfriend
// @Summary 创建StudentGirlfriend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body myFirstPackage.StudentGirlfriend true "创建StudentGirlfriend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /StuGF/createStudentGirlfriend [post]
func (StuGFApi *StudentGirlfriendApi) CreateStudentGirlfriend(c *gin.Context) {
	var StuGF myFirstPackage.StudentGirlfriend
	err := c.ShouldBindJSON(&StuGF)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := StuGFService.CreateStudentGirlfriend(StuGF); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteStudentGirlfriend 删除StudentGirlfriend
// @Tags StudentGirlfriend
// @Summary 删除StudentGirlfriend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body myFirstPackage.StudentGirlfriend true "删除StudentGirlfriend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /StuGF/deleteStudentGirlfriend [delete]
func (StuGFApi *StudentGirlfriendApi) DeleteStudentGirlfriend(c *gin.Context) {
	var StuGF myFirstPackage.StudentGirlfriend
	err := c.ShouldBindJSON(&StuGF)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := StuGFService.DeleteStudentGirlfriend(StuGF); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteStudentGirlfriendByIds 批量删除StudentGirlfriend
// @Tags StudentGirlfriend
// @Summary 批量删除StudentGirlfriend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除StudentGirlfriend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /StuGF/deleteStudentGirlfriendByIds [delete]
func (StuGFApi *StudentGirlfriendApi) DeleteStudentGirlfriendByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := StuGFService.DeleteStudentGirlfriendByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateStudentGirlfriend 更新StudentGirlfriend
// @Tags StudentGirlfriend
// @Summary 更新StudentGirlfriend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body myFirstPackage.StudentGirlfriend true "更新StudentGirlfriend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /StuGF/updateStudentGirlfriend [put]
func (StuGFApi *StudentGirlfriendApi) UpdateStudentGirlfriend(c *gin.Context) {
	var StuGF myFirstPackage.StudentGirlfriend
	err := c.ShouldBindJSON(&StuGF)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := StuGFService.UpdateStudentGirlfriend(StuGF); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindStudentGirlfriend 用id查询StudentGirlfriend
// @Tags StudentGirlfriend
// @Summary 用id查询StudentGirlfriend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query myFirstPackage.StudentGirlfriend true "用id查询StudentGirlfriend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /StuGF/findStudentGirlfriend [get]
func (StuGFApi *StudentGirlfriendApi) FindStudentGirlfriend(c *gin.Context) {
	var StuGF myFirstPackage.StudentGirlfriend
	err := c.ShouldBindQuery(&StuGF)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reStuGF, err := StuGFService.GetStudentGirlfriend(StuGF.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reStuGF": reStuGF}, c)
	}
}

// GetStudentGirlfriendList 分页获取StudentGirlfriend列表
// @Tags StudentGirlfriend
// @Summary 分页获取StudentGirlfriend列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query myFirstPackageReq.StudentGirlfriendSearch true "分页获取StudentGirlfriend列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /StuGF/getStudentGirlfriendList [get]
func (StuGFApi *StudentGirlfriendApi) GetStudentGirlfriendList(c *gin.Context) {
	var pageInfo myFirstPackageReq.StudentGirlfriendSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := StuGFService.GetStudentGirlfriendInfoList(pageInfo); err != nil {
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
