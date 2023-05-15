package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/myFirstPackage"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type StudentGirlfriendSearch struct{
    myFirstPackage.StudentGirlfriend
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
