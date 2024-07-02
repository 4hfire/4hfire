import { defineStore } from 'pinia'
import { ADMIN_INFO } from '/@/stores/constant/cacheKey'
import type { AdminInfo } from '/@/stores/interface'

export const useAdminInfo = defineStore('adminInfo', {
    state: (): AdminInfo => {
        return {
            id: 1,
            username: 'admin',
            nickname: 'Admin',
            avatar: '/static/images/avatar.png',
            last_login_time: '2024-06-17 11:37:31',
            token: 'bf0719f4-a057-4cd9-92f0-1e41e9b9a69d',
            // 是否是superAdmin，用于判定是否显示终端按钮等，不做任何权限判断
            super: false,
        }
    },
    actions: {
        dataFill(state: AdminInfo) {
            this.$state = { ...this.$state, ...state }
        },
        removeToken() {
            this.token = ''
        },
        setToken(token: string, type: 'auth' | 'refresh') {
            const field = type == 'auth' ? 'token' : 'refresh_token'
            this[field] = token
        },
        getToken(type: 'auth' | 'refresh' = 'auth') {
            return this.token
        },
        setSuper(val: boolean) {
            this.super = val
        },
    },
    persist: {
        key: ADMIN_INFO,
    },
})
