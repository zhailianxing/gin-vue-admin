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

type StudentInfoApi struct {
}

var stuInfoService = service.ServiceGroupApp.MyfirstpackageServiceGroup.StudentInfoService


// CreateStudentInfo 创建StudentInfo
// @Tags StudentInfo
// @Summary 创建StudentInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body myFirstPackage.StudentInfo true "创建StudentInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stuInfo/createStudentInfo [post]
func (stuInfoApi *StudentInfoApi) CreateStudentInfo(c *gin.Context) {
	var stuInfo myFirstPackage.StudentInfo
	err := c.ShouldBindJSON(&stuInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := stuInfoService.CreateStudentInfo(stuInfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteStudentInfo 删除StudentInfo
// @Tags StudentInfo
// @Summary 删除StudentInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body myFirstPackage.StudentInfo true "删除StudentInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /stuInfo/deleteStudentInfo [delete]
func (stuInfoApi *StudentInfoApi) DeleteStudentInfo(c *gin.Context) {
	var stuInfo myFirstPackage.StudentInfo
	err := c.ShouldBindJSON(&stuInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := stuInfoService.DeleteStudentInfo(stuInfo); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteStudentInfoByIds 批量删除StudentInfo
// @Tags StudentInfo
// @Summary 批量删除StudentInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除StudentInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /stuInfo/deleteStudentInfoByIds [delete]
func (stuInfoApi *StudentInfoApi) DeleteStudentInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := stuInfoService.DeleteStudentInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateStudentInfo 更新StudentInfo
// @Tags StudentInfo
// @Summary 更新StudentInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body myFirstPackage.StudentInfo true "更新StudentInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /stuInfo/updateStudentInfo [put]
func (stuInfoApi *StudentInfoApi) UpdateStudentInfo(c *gin.Context) {
	var stuInfo myFirstPackage.StudentInfo
	err := c.ShouldBindJSON(&stuInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := stuInfoService.UpdateStudentInfo(stuInfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindStudentInfo 用id查询StudentInfo
// @Tags StudentInfo
// @Summary 用id查询StudentInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query myFirstPackage.StudentInfo true "用id查询StudentInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /stuInfo/findStudentInfo [get]
func (stuInfoApi *StudentInfoApi) FindStudentInfo(c *gin.Context) {
	var stuInfo myFirstPackage.StudentInfo
	err := c.ShouldBindQuery(&stuInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if restuInfo, err := stuInfoService.GetStudentInfo(stuInfo.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"restuInfo": restuInfo}, c)
	}
}

// GetStudentInfoList 分页获取StudentInfo列表
// @Tags StudentInfo
// @Summary 分页获取StudentInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query myFirstPackageReq.StudentInfoSearch true "分页获取StudentInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stuInfo/getStudentInfoList [get]
func (stuInfoApi *StudentInfoApi) GetStudentInfoList(c *gin.Context) {
	var pageInfo myFirstPackageReq.StudentInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := stuInfoService.GetStudentInfoInfoList(pageInfo); err != nil {
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
