package Test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Test"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    TestReq "github.com/flipped-aurora/gin-vue-admin/server/model/Test/request"
)

type TestStudentInfoService struct {
}

// CreateTestStudentInfo 创建TestStudentInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (TestStudentShortService *TestStudentInfoService) CreateTestStudentInfo(TestStudentShort Test.TestStudentInfo) (err error) {
	err = global.GVA_DB.Create(&TestStudentShort).Error
	return err
}

// DeleteTestStudentInfo 删除TestStudentInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (TestStudentShortService *TestStudentInfoService)DeleteTestStudentInfo(TestStudentShort Test.TestStudentInfo) (err error) {
	err = global.GVA_DB.Delete(&TestStudentShort).Error
	return err
}

// DeleteTestStudentInfoByIds 批量删除TestStudentInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (TestStudentShortService *TestStudentInfoService)DeleteTestStudentInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Test.TestStudentInfo{},"id in ?",ids.Ids).Error
	return err
}

// UpdateTestStudentInfo 更新TestStudentInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (TestStudentShortService *TestStudentInfoService)UpdateTestStudentInfo(TestStudentShort Test.TestStudentInfo) (err error) {
	err = global.GVA_DB.Save(&TestStudentShort).Error
	return err
}

// GetTestStudentInfo 根据id获取TestStudentInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (TestStudentShortService *TestStudentInfoService)GetTestStudentInfo(id uint) (TestStudentShort Test.TestStudentInfo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&TestStudentShort).Error
	return
}

// GetTestStudentInfoInfoList 分页获取TestStudentInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (TestStudentShortService *TestStudentInfoService)GetTestStudentInfoInfoList(info TestReq.TestStudentInfoSearch) (list []Test.TestStudentInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&Test.TestStudentInfo{})
    var TestStudentShorts []Test.TestStudentInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&TestStudentShorts).Error
	return  TestStudentShorts, total, err
}
