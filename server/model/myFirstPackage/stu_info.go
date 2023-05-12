// 自动生成模板StudentInfo
package myFirstPackage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// StudentInfo 结构体
type StudentInfo struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`
      Age  *int `json:"age" form:"age" gorm:"column:age;comment:;"`
}


// TableName StudentInfo 表名
func (StudentInfo) TableName() string {
  return "student_info"
}

