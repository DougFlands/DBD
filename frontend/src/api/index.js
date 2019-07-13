import axios from 'axios'

axios.get = (url, params) => {
  return axios.request({
    url: url,
    method: 'get',
    params: params || {},
  })
};

axios.interceptors.response.use(
  response => {
    let data = response.data
    if (data.code === 200) {
      return data.data;
    } else {
      return Promise.reject(data)
    }
  },
  error => {
    return Promise.reject({
      errcode: error.response.status
    })
  }
);

const common = {
  getList(page = 1) {
    return axios.get(`http://localhost:3000`, {
      page,
    });
  },
};

export default {
  common
};