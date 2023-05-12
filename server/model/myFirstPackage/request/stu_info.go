package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/myFirstPackage"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type StudentInfoSearch struct{
    myFirstPackage.StudentInfo
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
