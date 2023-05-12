package myFirstPackage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/myFirstPackage"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    myFirstPackageReq "github.com/flipped-aurora/gin-vue-admin/server/model/myFirstPackage/request"
)

type StudentInfoService struct {
}

// CreateStudentInfo 创建StudentInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (stuInfoService *StudentInfoService) CreateStudentInfo(stuInfo myFirstPackage.StudentInfo) (err error) {
	err = global.GVA_DB.Create(&stuInfo).Error
	return err
}

// DeleteStudentInfo 删除StudentInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (stuInfoService *StudentInfoService)DeleteStudentInfo(stuInfo myFirstPackage.StudentInfo) (err error) {
	err = global.GVA_DB.Delete(&stuInfo).Error
	return err
}

// DeleteStudentInfoByIds 批量删除StudentInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (stuInfoService *StudentInfoService)DeleteStudentInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]myFirstPackage.StudentInfo{},"id in ?",ids.Ids).Error
	return err
}

// UpdateStudentInfo 更新StudentInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (stuInfoService *StudentInfoService)UpdateStudentInfo(stuInfo myFirstPackage.StudentInfo) (err error) {
	err = global.GVA_DB.Save(&stuInfo).Error
	return err
}

// GetStudentInfo 根据id获取StudentInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (stuInfoService *StudentInfoService)GetStudentInfo(id uint) (stuInfo myFirstPackage.StudentInfo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&stuInfo).Error
	return
}

// GetStudentInfoInfoList 分页获取StudentInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (stuInfoService *StudentInfoService)GetStudentInfoInfoList(info myFirstPackageReq.StudentInfoSearch) (list []myFirstPackage.StudentInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&myFirstPackage.StudentInfo{})
    var stuInfos []myFirstPackage.StudentInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&stuInfos).Error
	return  stuInfos, total, err
}
