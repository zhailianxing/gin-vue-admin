// 自动生成模板TestStudentInfo
package Test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// TestStudentInfo 结构体
type TestStudentInfo struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`
}


// TableName TestStudentInfo 表名
func (TestStudentInfo) TableName() string {
  return "test_student_info"
}

