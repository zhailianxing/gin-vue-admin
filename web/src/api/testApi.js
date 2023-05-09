import service from '@/utils/request'

export const testApi = (data) => {
  return service({
    url: '/test/test_api',
    method: 'post',
    data
  })
}

