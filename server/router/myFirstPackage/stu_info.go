package myFirstPackage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type StudentInfoRouter struct {
}

// InitStudentInfoRouter 初始化 StudentInfo 路由信息
func (s *StudentInfoRouter) InitStudentInfoRouter(Router *gin.RouterGroup) {
	stuInfoRouter := Router.Group("stuInfo").Use(middleware.OperationRecord())
	stuInfoRouterWithoutRecord := Router.Group("stuInfo")
	var stuInfoApi = v1.ApiGroupApp.MyfirstpackageApiGroup.StudentInfoApi
	{
		stuInfoRouter.POST("createStudentInfo", stuInfoApi.CreateStudentInfo)   // 新建StudentInfo
		stuInfoRouter.DELETE("deleteStudentInfo", stuInfoApi.DeleteStudentInfo) // 删除StudentInfo
		stuInfoRouter.DELETE("deleteStudentInfoByIds", stuInfoApi.DeleteStudentInfoByIds) // 批量删除StudentInfo
		stuInfoRouter.PUT("updateStudentInfo", stuInfoApi.UpdateStudentInfo)    // 更新StudentInfo
	}
	{
		stuInfoRouterWithoutRecord.GET("findStudentInfo", stuInfoApi.FindStudentInfo)        // 根据ID获取StudentInfo
		stuInfoRouterWithoutRecord.GET("getStudentInfoList", stuInfoApi.GetStudentInfoList)  // 获取StudentInfo列表
	}
}
