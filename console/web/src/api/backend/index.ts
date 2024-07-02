import md5 from 'js-md5';
import createAxios from '/@/utils/axios'
import { useAdminInfo } from '/@/stores/adminInfo'
import {cloneDeep} from "lodash-es";

export const url = '/auth/'

export function index() {
    return createAxios({
        url: url + 'index',
        method: 'get',
    })
}

export function login(method: 'get' | 'post', params: object = {}) {
    let data = cloneDeep(params)
    data.password = md5(data.password)
    return createAxios({
        url: url + 'login',
        data: data,
        method: method,
    })
}

export function logout() {
    const adminInfo = useAdminInfo()
    return createAxios({
        url: url + 'logout',
        method: 'POST',
        data: {
            refreshToken: adminInfo.getToken('refresh'),
        },
    })
}
