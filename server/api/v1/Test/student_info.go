package Test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/Test"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    TestReq "github.com/flipped-aurora/gin-vue-admin/server/model/Test/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type TestStudentInfoApi struct {
}

var TestStudentShortService = service.ServiceGroupApp.TestServiceGroup.TestStudentInfoService


// CreateTestStudentInfo 创建TestStudentInfo
// @Tags TestStudentInfo
// @Summary 创建TestStudentInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Test.TestStudentInfo true "创建TestStudentInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TestStudentShort/createTestStudentInfo [post]
func (TestStudentShortApi *TestStudentInfoApi) CreateTestStudentInfo(c *gin.Context) {
	var TestStudentShort Test.TestStudentInfo
	err := c.ShouldBindJSON(&TestStudentShort)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := TestStudentShortService.CreateTestStudentInfo(TestStudentShort); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTestStudentInfo 删除TestStudentInfo
// @Tags TestStudentInfo
// @Summary 删除TestStudentInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Test.TestStudentInfo true "删除TestStudentInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TestStudentShort/deleteTestStudentInfo [delete]
func (TestStudentShortApi *TestStudentInfoApi) DeleteTestStudentInfo(c *gin.Context) {
	var TestStudentShort Test.TestStudentInfo
	err := c.ShouldBindJSON(&TestStudentShort)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := TestStudentShortService.DeleteTestStudentInfo(TestStudentShort); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTestStudentInfoByIds 批量删除TestStudentInfo
// @Tags TestStudentInfo
// @Summary 批量删除TestStudentInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TestStudentInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /TestStudentShort/deleteTestStudentInfoByIds [delete]
func (TestStudentShortApi *TestStudentInfoApi) DeleteTestStudentInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := TestStudentShortService.DeleteTestStudentInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTestStudentInfo 更新TestStudentInfo
// @Tags TestStudentInfo
// @Summary 更新TestStudentInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Test.TestStudentInfo true "更新TestStudentInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TestStudentShort/updateTestStudentInfo [put]
func (TestStudentShortApi *TestStudentInfoApi) UpdateTestStudentInfo(c *gin.Context) {
	var TestStudentShort Test.TestStudentInfo
	err := c.ShouldBindJSON(&TestStudentShort)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := TestStudentShortService.UpdateTestStudentInfo(TestStudentShort); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTestStudentInfo 用id查询TestStudentInfo
// @Tags TestStudentInfo
// @Summary 用id查询TestStudentInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Test.TestStudentInfo true "用id查询TestStudentInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TestStudentShort/findTestStudentInfo [get]
func (TestStudentShortApi *TestStudentInfoApi) FindTestStudentInfo(c *gin.Context) {
	var TestStudentShort Test.TestStudentInfo
	err := c.ShouldBindQuery(&TestStudentShort)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reTestStudentShort, err := TestStudentShortService.GetTestStudentInfo(TestStudentShort.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reTestStudentShort": reTestStudentShort}, c)
	}
}

// GetTestStudentInfoList 分页获取TestStudentInfo列表
// @Tags TestStudentInfo
// @Summary 分页获取TestStudentInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query TestReq.TestStudentInfoSearch true "分页获取TestStudentInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TestStudentShort/getTestStudentInfoList [get]
func (TestStudentShortApi *TestStudentInfoApi) GetTestStudentInfoList(c *gin.Context) {
	var pageInfo TestReq.TestStudentInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := TestStudentShortService.GetTestStudentInfoInfoList(pageInfo); err != nil {
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
