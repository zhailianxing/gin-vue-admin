import service from '@/utils/request'

// @Tags TestStudentInfo
// @Summary 创建TestStudentInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestStudentInfo true "创建TestStudentInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TestStudentShort/createTestStudentInfo [post]
export const createTestStudentInfo = (data) => {
  return service({
    url: '/TestStudentShort/createTestStudentInfo',
    method: 'post',
    data
  })
}

// @Tags TestStudentInfo
// @Summary 删除TestStudentInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestStudentInfo true "删除TestStudentInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TestStudentShort/deleteTestStudentInfo [delete]
export const deleteTestStudentInfo = (data) => {
  return service({
    url: '/TestStudentShort/deleteTestStudentInfo',
    method: 'delete',
    data
  })
}

// @Tags TestStudentInfo
// @Summary 删除TestStudentInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TestStudentInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /TestStudentShort/deleteTestStudentInfo [delete]
export const deleteTestStudentInfoByIds = (data) => {
  return service({
    url: '/TestStudentShort/deleteTestStudentInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags TestStudentInfo
// @Summary 更新TestStudentInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestStudentInfo true "更新TestStudentInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /TestStudentShort/updateTestStudentInfo [put]
export const updateTestStudentInfo = (data) => {
  return service({
    url: '/TestStudentShort/updateTestStudentInfo',
    method: 'put',
    data
  })
}

// @Tags TestStudentInfo
// @Summary 用id查询TestStudentInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TestStudentInfo true "用id查询TestStudentInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /TestStudentShort/findTestStudentInfo [get]
export const findTestStudentInfo = (params) => {
  return service({
    url: '/TestStudentShort/findTestStudentInfo',
    method: 'get',
    params
  })
}

// @Tags TestStudentInfo
// @Summary 分页获取TestStudentInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TestStudentInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /TestStudentShort/getTestStudentInfoList [get]
export const getTestStudentInfoList = (params) => {
  return service({
    url: '/TestStudentShort/getTestStudentInfoList',
    method: 'get',
    params
  })
}
