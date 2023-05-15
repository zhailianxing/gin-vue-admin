// 自动生成模板StudentGirlfriend
package myFirstPackage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// StudentGirlfriend 结构体
type StudentGirlfriend struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`
      Age  *int `json:"age" form:"age" gorm:"column:age;comment:;"`
      Sex  *int `json:"sex" form:"sex" gorm:"column:sex;comment:;"`
}


// TableName StudentGirlfriend 表名
func (StudentGirlfriend) TableName() string {
  return "student_girlfriend"
}

