package myFirstPackage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/myFirstPackage"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    myFirstPackageReq "github.com/flipped-aurora/gin-vue-admin/server/model/myFirstPackage/request"
)

type StudentGirlfriendService struct {
}

// CreateStudentGirlfriend 创建StudentGirlfriend记录
// Author [piexlmax](https://github.com/piexlmax)
func (StuGFService *StudentGirlfriendService) CreateStudentGirlfriend(StuGF myFirstPackage.StudentGirlfriend) (err error) {
	err = global.GVA_DB.Create(&StuGF).Error
	return err
}

// DeleteStudentGirlfriend 删除StudentGirlfriend记录
// Author [piexlmax](https://github.com/piexlmax)
func (StuGFService *StudentGirlfriendService)DeleteStudentGirlfriend(StuGF myFirstPackage.StudentGirlfriend) (err error) {
	err = global.GVA_DB.Delete(&StuGF).Error
	return err
}

// DeleteStudentGirlfriendByIds 批量删除StudentGirlfriend记录
// Author [piexlmax](https://github.com/piexlmax)
func (StuGFService *StudentGirlfriendService)DeleteStudentGirlfriendByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]myFirstPackage.StudentGirlfriend{},"id in ?",ids.Ids).Error
	return err
}

// UpdateStudentGirlfriend 更新StudentGirlfriend记录
// Author [piexlmax](https://github.com/piexlmax)
func (StuGFService *StudentGirlfriendService)UpdateStudentGirlfriend(StuGF myFirstPackage.StudentGirlfriend) (err error) {
	err = global.GVA_DB.Save(&StuGF).Error
	return err
}

// GetStudentGirlfriend 根据id获取StudentGirlfriend记录
// Author [piexlmax](https://github.com/piexlmax)
func (StuGFService *StudentGirlfriendService)GetStudentGirlfriend(id uint) (StuGF myFirstPackage.StudentGirlfriend, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&StuGF).Error
	return
}

// GetStudentGirlfriendInfoList 分页获取StudentGirlfriend记录
// Author [piexlmax](https://github.com/piexlmax)
func (StuGFService *StudentGirlfriendService)GetStudentGirlfriendInfoList(info myFirstPackageReq.StudentGirlfriendSearch) (list []myFirstPackage.StudentGirlfriend, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&myFirstPackage.StudentGirlfriend{})
    var StuGFs []myFirstPackage.StudentGirlfriend
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&StuGFs).Error
	return  StuGFs, total, err
}
