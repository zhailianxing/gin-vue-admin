import service from '@/utils/request'

// @Tags StudentGirlfriend
// @Summary 创建StudentGirlfriend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StudentGirlfriend true "创建StudentGirlfriend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /StuGF/createStudentGirlfriend [post]
export const createStudentGirlfriend = (data) => {
  return service({
    url: '/StuGF/createStudentGirlfriend',
    method: 'post',
    data
  })
}

// @Tags StudentGirlfriend
// @Summary 删除StudentGirlfriend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StudentGirlfriend true "删除StudentGirlfriend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /StuGF/deleteStudentGirlfriend [delete]
export const deleteStudentGirlfriend = (data) => {
  return service({
    url: '/StuGF/deleteStudentGirlfriend',
    method: 'delete',
    data
  })
}

// @Tags StudentGirlfriend
// @Summary 删除StudentGirlfriend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除StudentGirlfriend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /StuGF/deleteStudentGirlfriend [delete]
export const deleteStudentGirlfriendByIds = (data) => {
  return service({
    url: '/StuGF/deleteStudentGirlfriendByIds',
    method: 'delete',
    data
  })
}

// @Tags StudentGirlfriend
// @Summary 更新StudentGirlfriend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StudentGirlfriend true "更新StudentGirlfriend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /StuGF/updateStudentGirlfriend [put]
export const updateStudentGirlfriend = (data) => {
  return service({
    url: '/StuGF/updateStudentGirlfriend',
    method: 'put',
    data
  })
}

// @Tags StudentGirlfriend
// @Summary 用id查询StudentGirlfriend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.StudentGirlfriend true "用id查询StudentGirlfriend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /StuGF/findStudentGirlfriend [get]
export const findStudentGirlfriend = (params) => {
  return service({
    url: '/StuGF/findStudentGirlfriend',
    method: 'get',
    params
  })
}

// @Tags StudentGirlfriend
// @Summary 分页获取StudentGirlfriend列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取StudentGirlfriend列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /StuGF/getStudentGirlfriendList [get]
export const getStudentGirlfriendList = (params) => {
  return service({
    url: '/StuGF/getStudentGirlfriendList',
    method: 'get',
    params
  })
}
