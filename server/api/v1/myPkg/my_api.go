package myPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type MyApi struct {
}

func (m *MyApi) MyCreateApi(ctx *gin.Context) {
	myApiService.CreateApiService() //service 层
	response.Ok(ctx)
}
