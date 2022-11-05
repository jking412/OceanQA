import axios from "axios"

const request = axios.create({
    baseURL: 'http://localhost:9999/v1',
    timeout: 10000,
})

export default request
