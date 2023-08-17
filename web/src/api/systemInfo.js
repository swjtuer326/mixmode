import axios from '@/utils/axios'
// @Tags systrm
// @Summary 获取配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /system/getSystemConfig [post]
export const getSystemConfig = () => {
  return axios({
    url: '/system/getSystemConfig',
    method: 'post'
  })
}


// @Tags system
// @Summary 获取服务器运行状态
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /system/getServerInfo [post]
export const getSystemState = () => {
  return axios({
    url: '/system/SystemInfo',
    method: 'post',
  })
}
