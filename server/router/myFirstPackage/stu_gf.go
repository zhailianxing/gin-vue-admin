package myFirstPackage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type StudentGirlfriendRouter struct {
}

// InitStudentGirlfriendRouter 初始化 StudentGirlfriend 路由信息
func (s *StudentGirlfriendRouter) InitStudentGirlfriendRouter(Router *gin.RouterGroup) {
	StuGFRouter := Router.Group("StuGF").Use(middleware.OperationRecord())
	StuGFRouterWithoutRecord := Router.Group("StuGF")
	var StuGFApi = v1.ApiGroupApp.MyfirstpackageApiGroup.StudentGirlfriendApi
	{
		StuGFRouter.POST("createStudentGirlfriend", StuGFApi.CreateStudentGirlfriend)   // 新建StudentGirlfriend
		StuGFRouter.DELETE("deleteStudentGirlfriend", StuGFApi.DeleteStudentGirlfriend) // 删除StudentGirlfriend
		StuGFRouter.DELETE("deleteStudentGirlfriendByIds", StuGFApi.DeleteStudentGirlfriendByIds) // 批量删除StudentGirlfriend
		StuGFRouter.PUT("updateStudentGirlfriend", StuGFApi.UpdateStudentGirlfriend)    // 更新StudentGirlfriend
	}
	{
		StuGFRouterWithoutRecord.GET("findStudentGirlfriend", StuGFApi.FindStudentGirlfriend)        // 根据ID获取StudentGirlfriend
		StuGFRouterWithoutRecord.GET("getStudentGirlfriendList", StuGFApi.GetStudentGirlfriendList)  // 获取StudentGirlfriend列表
	}
}
