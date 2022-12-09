package myPkg

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type MyApiRouter struct {
}

// 从router/example/exa_customer.go 拷贝 初始化函数
func (m *MyApiRouter) InitMyApiRouter(group *gin.RouterGroup) {
	apiRouter := group.Group("myapi") // 设置 组名字
	api := v1.ApiGroupApp.MypkgApiGroup.MyApi
	{
		apiRouter.GET("createApi", api.MyCreateApi)
	}
}
