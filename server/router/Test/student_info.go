package Test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestStudentInfoRouter struct {
}

// InitTestStudentInfoRouter 初始化 TestStudentInfo 路由信息
func (s *TestStudentInfoRouter) InitTestStudentInfoRouter(Router *gin.RouterGroup) {
	TestStudentShortRouter := Router.Group("TestStudentShort").Use(middleware.OperationRecord())
	TestStudentShortRouterWithoutRecord := Router.Group("TestStudentShort")
	var TestStudentShortApi = v1.ApiGroupApp.TestApiGroup.TestStudentInfoApi
	{
		TestStudentShortRouter.POST("createTestStudentInfo", TestStudentShortApi.CreateTestStudentInfo)   // 新建TestStudentInfo
		TestStudentShortRouter.DELETE("deleteTestStudentInfo", TestStudentShortApi.DeleteTestStudentInfo) // 删除TestStudentInfo
		TestStudentShortRouter.DELETE("deleteTestStudentInfoByIds", TestStudentShortApi.DeleteTestStudentInfoByIds) // 批量删除TestStudentInfo
		TestStudentShortRouter.PUT("updateTestStudentInfo", TestStudentShortApi.UpdateTestStudentInfo)    // 更新TestStudentInfo
	}
	{
		TestStudentShortRouterWithoutRecord.GET("findTestStudentInfo", TestStudentShortApi.FindTestStudentInfo)        // 根据ID获取TestStudentInfo
		TestStudentShortRouterWithoutRecord.GET("getTestStudentInfoList", TestStudentShortApi.GetTestStudentInfoList)  // 获取TestStudentInfo列表
	}
}
