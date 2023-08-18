import axios from "@/utils/axios";

export const testGDMA = () => {
    return axios({
      url: '/sophon/TestGDMA',
      method: 'post',
})}

export const testCDMA = () => {
    return axios({
      url: '/sophon/TestCDMA',
      method: 'post',
})}