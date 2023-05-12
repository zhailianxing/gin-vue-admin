import service from '@/utils/request'

// @Tags StudentInfo
// @Summary 创建StudentInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StudentInfo true "创建StudentInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stuInfo/createStudentInfo [post]
export const createStudentInfo = (data) => {
  return service({
    url: '/stuInfo/createStudentInfo',
    method: 'post',
    data
  })
}

// @Tags StudentInfo
// @Summary 删除StudentInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StudentInfo true "删除StudentInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /stuInfo/deleteStudentInfo [delete]
export const deleteStudentInfo = (data) => {
  return service({
    url: '/stuInfo/deleteStudentInfo',
    method: 'delete',
    data
  })
}

// @Tags StudentInfo
// @Summary 删除StudentInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除StudentInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /stuInfo/deleteStudentInfo [delete]
export const deleteStudentInfoByIds = (data) => {
  return service({
    url: '/stuInfo/deleteStudentInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags StudentInfo
// @Summary 更新StudentInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StudentInfo true "更新StudentInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /stuInfo/updateStudentInfo [put]
export const updateStudentInfo = (data) => {
  return service({
    url: '/stuInfo/updateStudentInfo',
    method: 'put',
    data
  })
}

// @Tags StudentInfo
// @Summary 用id查询StudentInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.StudentInfo true "用id查询StudentInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /stuInfo/findStudentInfo [get]
export const findStudentInfo = (params) => {
  return service({
    url: '/stuInfo/findStudentInfo',
    method: 'get',
    params
  })
}

// @Tags StudentInfo
// @Summary 分页获取StudentInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取StudentInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stuInfo/getStudentInfoList [get]
export const getStudentInfoList = (params) => {
  return service({
    url: '/stuInfo/getStudentInfoList',
    method: 'get',
    params
  })
}
