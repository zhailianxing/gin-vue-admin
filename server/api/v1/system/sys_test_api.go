package system

import (
	"github.com/gin-gonic/gin"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
)

type MyTestApi struct{}

func (*MyTestApi) TestApi(c *gin.Context) {

	response.OkWithData("you're so beautiful", c)
}
