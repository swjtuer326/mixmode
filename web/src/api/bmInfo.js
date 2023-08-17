import axios from "@/utils/axios";

export const getTemData = (data) => {
    return axios({
      url: '/bm1684x/TemData',
      method: 'post',
      data
})}

export const getPowerData = (data) => {
    return axios({
      url: '/bm1684x/PowerData',
      method: 'post',
      data
})}