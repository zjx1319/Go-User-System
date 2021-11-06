import request from '@/api/request'

export function setToken(token) {
	request.defaults.headers.common['Authorization'] = "Bearer " + token
}
