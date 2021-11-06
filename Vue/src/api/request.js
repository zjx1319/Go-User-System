const axios = require('axios').default;
const instance = axios.create({
    baseURL: '',
    timeout: 60000,
  });
export default instance