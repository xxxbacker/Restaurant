import axios, { InternalAxiosRequestConfig } from 'axios';

const $authHost = axios.create({
	baseURL: 'http://localhost:8080/api',
})
const $host = axios.create({
	baseURL: 'http://localhost:8080/auth'
})

const authInterceptor = (config: InternalAxiosRequestConfig) =>{
	config.headers.authorization = `Bearer ${localStorage.getItem('token')}`
	return config
}

$authHost.interceptors.request.use(authInterceptor)

export {
	$host,
	$authHost
}
