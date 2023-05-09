package system

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
)

type TestApiRouter struct{}

func (s *TestApiRouter) InitTestRouter(Router *gin.RouterGroup) {
	testRouter := Router.Group("test")
	testApi := v1.ApiGroupApp.SystemApiGroup.MyTestApi
	{
		// 127.0.0.1:8080/api/test/test_api
		testRouter.POST("test_api", testApi.TestApi)
	}
}
